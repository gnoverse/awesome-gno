package generators

import (
	"html/template"
	"os"
	"path/filepath"
	"time"
)

type HTMLGenerator struct {
	Data        ListData
	TemplateDir string
}

func NewHTMLGenerator(data ListData, templateDir string) *HTMLGenerator {
	return &HTMLGenerator{
		Data:        data,
		TemplateDir: templateDir,
	}
}

type templateData struct {
	Metadata    Metadata
	Contents    []Section
	LastUpdated string
}

func (g *HTMLGenerator) Generate(outputPath string) error {
	// Parse template
	tmpl, err := template.New("awesome-gno.html").Funcs(template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"len": func(s interface{}) int {
			switch v := s.(type) {
			case []Section:
				return len(v)
			case []Item:
				return len(v)
			default:
				return 0
			}
		},
		"or": func(a, b interface{}) interface{} {
			if a != nil {
				return a
			}
			return b
		},
	}).ParseFiles(filepath.Join(g.TemplateDir, "awesome-gno.html"))
	if err != nil {
		return err
	}

	// Create output file
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Prepare template data
	data := templateData{
		Metadata:    g.Data.Metadata,
		Contents:    g.Data.Contents,
		LastUpdated: time.Now().Format("January 2, 2006"),
	}

	// Execute template
	return tmpl.Execute(f, data)
}
