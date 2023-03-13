package main

import (
	"embed"
	"encoding/csv"
	"html/template"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

const othersCategory = "Others"

type Entry struct {
	Title       string
	Link        string
	Description string
	Category    string
	IsStaffPick bool
}

type Category struct {
	Title   string
	Slug    string
	Note    string
	Order   int
	Entries []*Entry
}

type StaffPick struct {
	Link     string
	Category bool
	Page     bool
}

var (
	//go:embed templates
	fs embed.FS
)

func main() {
	categoriesFile, err := os.Open("categories.csv")
	if err != nil {
		panic(err)
	}
	defer categoriesFile.Close()

	categoriesReader := csv.NewReader(categoriesFile)
	categoriesRecords, err := categoriesReader.ReadAll()
	if err != nil {
		panic(err)
	}
	categories := make(map[string]*Category)
	for i, row := range categoriesRecords {
		if i == 0 {
			continue
		}
		category := Category{
			Title: row[0],
			Slug:  slugify(row[0]),
			Note:  row[1],
			Order: i,
		}

		categories[category.Title] = &category
	}

	staffpicksFile, err := os.Open("staffpicks.csv")
	if err != nil {
		panic(err)
	}
	defer staffpicksFile.Close()

	staffpicksReader := csv.NewReader(staffpicksFile)
	staffpicksRecords, err := staffpicksReader.ReadAll()
	if err != nil {
		panic(err)
	}

	staffpicks := make(map[string]*StaffPick)
	for i, row := range staffpicksRecords {
		if i == 0 {
			continue
		}
		staffpicks[row[0]] = &StaffPick{
			Link: strings.TrimSpace(row[0]),
		}
	}

	entriesFile, err := os.Open("entries.csv")
	if err != nil {
		panic(err)
	}
	defer entriesFile.Close()

	entriesReader := csv.NewReader(entriesFile)
	entriesRecords, err := entriesReader.ReadAll()
	if err != nil {
		panic(err)
	}
	var entriesWithNoCat []*Entry
	for i, row := range entriesRecords {
		if i == 0 {
			continue
		}
		entry := Entry{
			Title:       strings.TrimSpace(row[0]),
			Link:        strings.TrimSpace(row[1]),
			Description: strings.TrimSpace(row[2]),
			Category:    strings.TrimSpace(row[3]),
		}
		if _, ok := staffpicks[entry.Link]; ok {
			entry.IsStaffPick = true
		}
		if _, ok := categories[entry.Category]; !ok {
			entriesWithNoCat = append(entriesWithNoCat, &entry)
			continue
		}

		categories[entry.Category].Entries = append(categories[entry.Category].Entries, &entry)

	}

	if len(entriesWithNoCat) > 0 {
		categories[othersCategory] = &Category{
			Title:   othersCategory,
			Entries: entriesWithNoCat,
			Order:   len(categories),
			Slug:    slugify(othersCategory),
		}
	}

	var cat []*Category
	for _, c := range categories {
		cat = append(cat, c)
	}

	sort.Slice(cat, func(i, j int) bool {
		return cat[i].Order < cat[j].Order
	})

	mardownFile, err := os.Create("README.md")
	if err != nil {
		panic(err)
	}
	defer mardownFile.Close()

	if err := createReadme(mardownFile, cat); err != nil {
		panic(err)
	}

	websiteFile, err := os.Create("website/index.html")
	if err != nil {
		panic(err)
	}
	defer websiteFile.Close()

	if err := createWebsite(websiteFile, cat); err != nil {
		panic(err)
	}

	log.Println("update README.md and Website")
}

func slugify(title string) string {
	slug := strings.ReplaceAll(strings.ToLower(title), " ", "-")
	reg := regexp.MustCompile("[^a-z0-9-_]")
	return reg.ReplaceAllString(slug, "")
}

func createReadme(w io.Writer, categories []*Category) error {
	tmplt, err := template.ParseFS(fs, "templates/readme_template.md")
	if err != nil {
		return err
	}
	return tmplt.Execute(w, categories)
}

func createWebsite(w io.Writer, categories []*Category) error {
	tmplt, err := template.ParseFS(fs, "templates/website_template.html")
	if err != nil {
		return err
	}
	return tmplt.Execute(w, categories)
}
