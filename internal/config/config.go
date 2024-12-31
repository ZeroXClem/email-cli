package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type SMTPProfile struct {
	Name         string `json:"name"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	IsDefault    bool   `json:"is_default"`
}

type Config struct {
type TLSConfig struct {
	Enabled            bool   `json:"enabled"`
	SkipVerify         bool   `json:"skip_verify"`
	ServerName         string `json:"server_name,omitempty"`
	CertificatePath    string `json:"certificate_path,omitempty"`
	PrivateKeyPath     string `json:"private_key_path,omitempty"`
}

type EncryptedData struct {
	Data string `json:"data"`
	IV   string `json:"iv"`
}
	SMTPProfiles []SMTPProfile `json:"smtp_profiles"`
}

const (
	defaultConfigDir  = ".email-cli"
	defaultConfigFile = "config.json"
)

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(homeDir, defaultConfigDir)
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return "", err
	}

	return filepath.Join(configDir, defaultConfigFile), nil
}

func LoadConfig() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{SMTPProfiles: []SMTPProfile{}}, nil
		}
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func SaveConfig(config *Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0600)
}

func GetDefaultProfile(config *Config) *SMTPProfile {
	for _, profile := range config.SMTPProfiles {
		if profile.IsDefault {
			return &profile
		}
	}
	return nil
}

func AddProfile(config *Config, profile SMTPProfile) error {
	config.SMTPProfiles = append(config.SMTPProfiles, profile)
	return SaveConfig(config)
}

func RemoveProfile(config *Config, name string) error {
	for i, profile := range config.SMTPProfiles {
		if profile.Name == name {
			config.SMTPProfiles = append(config.SMTPProfiles[:i], config.SMTPProfiles[i+1:]...)
			return SaveConfig(config)
		}
	}
	return nil
}
