package config

type Config struct {
	SMTPHost     string `json:"smtp_host"`
	SMTPPort     int    `json:"smtp_port"`
	SMTPUsername string `json:"smtp_username"`
	SMTPPassword string `json:"smtp_password"`
	DefaultFrom  string `json:"default_from"`
}

func LoadConfig(filepath string) (*Config, error) {
	// TODO: Implement config loading from JSON file
	return &Config{}, nil
}
