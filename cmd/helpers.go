package cmd

import (
	"embed"
	"fmt"
	"os"
	"text/template"
)

//go:embed templates/*
var templateFS embed.FS

func capitalize(input string) string {
	if len(input) == 0 {
		return input
	}
	return string(input[0]-32) + input[1:]
}

// Utility function to render a template and create a file
func generateFileFromTemplate(filePath, templatePath string, data map[string]string) {
	tmpl, err := template.ParseFS(templateFS, templatePath)
	if err != nil {
		fmt.Printf("Error parsing template: %s\n", err)
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Printf("Error executing template: %s\n", err)
	}
}
