package config

// Config holds all configuration for the application
type Config struct {
	SMTP SMTPConfig
}

// SMTPConfig holds SMTP-specific configuration
type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

// GetConfig returns the application configuration
func GetConfig() *Config {
	return &Config{
		SMTP: SMTPConfig{
			Host:     "",
			Port:     0,
			Username: "",
			Password: "",
		},
	}
}
