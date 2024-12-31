package config

import (
	"testing"
	"os"
	"path/filepath"
)

func TestConfigOperations(t *testing.T) {
	// Setup temporary config directory
	originalHome := os.Getenv("HOME")
	tmpHome := t.TempDir()
	os.Setenv("HOME", tmpHome)
	defer os.Setenv("HOME", originalHome)

	// Test loading empty config
	conf, err := LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Test adding profile
	profile := SMTPProfile{
		Name: "test",
		Host: "smtp.test.com",
		Port: 587,
		Username: "test@test.com",
		Password: "password",
		IsDefault: true,
	}

	err = AddProfile(conf, profile)
	if err != nil {
		t.Errorf("Failed to add profile: %v", err)
	}

	// Test getting default profile
	defaultProfile := GetDefaultProfile(conf)
	if defaultProfile == nil {
		t.Error("Expected default profile, got nil")
	}

	// Test removing profile
	err = RemoveProfile(conf, "test")
	if err != nil {
		t.Errorf("Failed to remove profile: %v", err)
	}

	// Verify profile was removed
	conf, _ = LoadConfig()
	if len(conf.SMTPProfiles) != 0 {
		t.Error("Expected empty profiles after removal")
	}
}