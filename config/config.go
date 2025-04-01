package config

import (
	"os"
	"strconv"
)

type Config struct {
	MongoURI     string
	DatabaseName string

	JWTSecret     string
	JWTExpiration int64

	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
	SMTPFrom     string
	UseEmailAPI  bool
	SendGridKey  string
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func LoadConfig() (*Config, error) {
	jwtExpHours, err := strconv.Atoi((getEnv("JWT_EXPIRATION", "24")))
	if err != nil {
		jwtExpHours = 24
	}
	jwtSecret := getEnv("JWT_SECRET", "default_jwt_secret")

	smtpPort, err := strconv.Atoi(getEnv("SMTP_PORT", "587"))
	if err != nil {
		smtpPort = 587
	}
	smtpHost := getEnv("SMTP_HOST", "smtp.example.com")
	smtpUserName := getEnv("SMTP_USERNAME", "username")
	smtpPassword := getEnv("SMTP_PASSWORD", "password")
	smtpFrom := getEnv("SMTP_FROM", "")

	useEmailAPI := getEnv("USE_EMAIL_API", "false") == "true"
	sendgridApiKey := getEnv("SENDGRID_API_KEY", "")

	mongoURI := getEnv("MONGO_URI", "mongodb://localhost:27017")
	databaseName := getEnv("DATABASE_NAME", "task-mate-pro")

	return &Config{
		MongoURI:      mongoURI,
		DatabaseName:  databaseName,
		JWTSecret:     jwtSecret,
		JWTExpiration: int64(jwtExpHours) * 3600,
		SMTPHost:      smtpHost,
		SMTPPort:      smtpPort,
		SMTPUsername:  smtpUserName,
		SMTPPassword:  smtpPassword,
		SMTPFrom:      smtpFrom,
		UseEmailAPI:   useEmailAPI,
		SendGridKey:   sendgridApiKey,
	}, nil
}
