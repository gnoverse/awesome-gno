package generators

import (
	"fmt"
	"strings"
)

type MarkdownGenerator struct {
	Data ListData
}

func NewMarkdownGenerator(data ListData) *MarkdownGenerator {
	return &MarkdownGenerator{Data: data}
}

func (g *MarkdownGenerator) Generate() string {
	var sb strings.Builder

	// Write header
	sb.WriteString(fmt.Sprintf("# %s\n\n", g.Data.Metadata.Title))
	if g.Data.Metadata.Badge != "" {
		sb.WriteString(fmt.Sprintf("[![Awesome](%s)](%s)\n\n",
			g.Data.Metadata.Badge,
			"https://awesome.re"))
	}

	// Write intro
	if g.Data.Metadata.IntroQuestion != "" {
		sb.WriteString(fmt.Sprintf("> %s\n\n", g.Data.Metadata.IntroQuestion))
	}
	sb.WriteString(g.Data.Metadata.IntroText + "\n\n")

	// Write purposes
	for _, purpose := range g.Data.Metadata.Purposes {
		sb.WriteString(fmt.Sprintf("- %s\n", purpose))
	}
	sb.WriteString("\n")

	if g.Data.Metadata.Notice != "" {
		sb.WriteString(fmt.Sprintf("> %s\n\n", g.Data.Metadata.Notice))
	}

	// Write table of contents
	sb.WriteString("## Contents\n\n")
	for _, section := range g.Data.Contents {
		sb.WriteString(fmt.Sprintf("- [%s](#%s)\n",
			section.Section,
			strings.ToLower(strings.ReplaceAll(section.Section, " ", "-"))))
	}
	sb.WriteString("\n")

	// Write sections
	for _, section := range g.Data.Contents {
		sb.WriteString(fmt.Sprintf("## %s\n\n", section.Section))
		if section.Description != "" {
			sb.WriteString(section.Description + "\n\n")
		}
		for _, item := range section.Items {
			sb.WriteString(fmt.Sprintf("- [%s](%s) - %s\n",
				item.Title,
				item.URL,
				item.Description))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
