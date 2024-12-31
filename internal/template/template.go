package template

import (
	"os"
	"path/filepath"
	"text/template"
)

type TemplateData struct {
	To      string
	From    string
	Subject string
	Body    map[string]interface{}
}

func LoadTemplate(name string) (*template.Template, error) {
	templateDir := "templates"
	templateFile := filepath.Join(templateDir, name+".tmpl")

	content, err := os.ReadFile(templateFile)
	if err != nil {
		return nil, err
	}

	return template.New(name).Parse(string(content))
}

func ParseTemplate(tmpl *template.Template, data TemplateData) (string, error) {
	var result strings.Builder
	if err := tmpl.Execute(&result, data); err != nil {
		return "", err
	}
	return result.String(), nil
}