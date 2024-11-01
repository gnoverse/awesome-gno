package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"awesome-gno/generators"

	"gopkg.in/yaml.v3"
)

func main() {
	inputFile := flag.String("input", "data/list.yaml", "Input YAML file")
	outputMD := flag.String("md", "output/markdown/README.md", "Output Markdown file")
	outputHTML := flag.String("html", "output/html/index.html", "Output HTML file")
	templateDir := flag.String("templates", "templates", "Templates directory")
	flag.Parse()

	// Ensure output directories exist
	os.MkdirAll(filepath.Dir(*outputMD), 0755)
	os.MkdirAll(filepath.Dir(*outputHTML), 0755)

	// Read and parse YAML
	yamlData, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var data generators.ListData
	if err := yaml.Unmarshal(yamlData, &data); err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	// Generate Markdown
	mdGen := generators.NewMarkdownGenerator(data)
	markdown := mdGen.Generate()
	if err := os.WriteFile(*outputMD, []byte(markdown), 0644); err != nil {
		log.Fatalf("Error writing Markdown file: %v", err)
	}
	fmt.Printf("Generated Markdown: %s\n", *outputMD)

	// Generate HTML
	htmlGen := generators.NewHTMLGenerator(data, *templateDir)
	if err := htmlGen.Generate(*outputHTML); err != nil {
		log.Fatalf("Error generating HTML: %v", err)
	}
	fmt.Printf("Generated HTML: %s\n", *outputHTML)
}
