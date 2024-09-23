package main

import (
	"html/template"
	"log"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

type List struct {
	Metadata struct {
		Title       string   `yaml:"title"`
		Description string   `yaml:"description"`
		Banner      string   `yaml:"banner"`
		Badge       string   `yaml:"badge"`
		Purposes    []string `yaml:"purposes"`
		Notice      string   `yaml:"notice"`
	} `yaml:"metadata"`
	Contents []Section `yaml:"contents"`
}

type Section struct {
	Section     string `yaml:"section"`
	Description string `yaml:"description,omitempty"`
	Items       []Item `yaml:"items"`
}

type Item struct {
	Title       string `yaml:"title"`
	URL         string `yaml:"url"`
	Description string `yaml:"description"`
}

// Custom functions
func add(a, b int) int {
	return a + b
}

func urlize(s string) string {
	// Transform the string into a URL-friendly anchor
	// For example, "Official Links" becomes "official-links"
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	s = regexp.MustCompile(`[^\w\s-]`).ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, " ", "-")
	return s
}

func main() {
	// Read YAML file
	data, err := os.ReadFile("data/list.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var list List
	err = yaml.Unmarshal(data, &list)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Create a FuncMap with your custom functions
	funcMap := template.FuncMap{
		"add":    add,
		"urlize": urlize,
	}

	// Generate README.md
	readmeTemplate, err := template.New("README.md.tmpl").Funcs(funcMap).ParseFiles("templates/README.md.tmpl")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	readmeFile, err := os.Create("README.md")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer readmeFile.Close()

	err = readmeTemplate.Execute(readmeFile, list)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

// Similarly, generate your HTML pages if needed
