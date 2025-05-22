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
			Host:     "sandbox.smtp.mailtrap.io",
			Port:     2525,
			Username: "b630fa712fb0d2",
			Password: "2dcaf1a05a1a64",
		},
	}
}
