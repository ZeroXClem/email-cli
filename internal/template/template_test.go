package template

import (
	"testing"
	"os"
	"path/filepath"
)

func TestLoadTemplate(t *testing.T) {
	// Create temporary template file
	tmpDir := t.TempDir()
	tmplContent := "Subject: {{.Subject}}\nHello {{.Body.name}}!"
	tmplPath := filepath.Join(tmpDir, "test.tmpl")

	err := os.WriteFile(tmplPath, []byte(tmplContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test template: %v", err)
	}

	tmpl, err := LoadTemplate("test")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	data := TemplateData{
		Subject: "Test Email",
		Body: map[string]interface{}{
			"name": "John",
		},
	}

	result, err := ParseTemplate(tmpl, data)
	if err != nil {
		t.Errorf("Failed to parse template: %v", err)
	}

	expected := "Subject: Test Email\nHello John!"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}