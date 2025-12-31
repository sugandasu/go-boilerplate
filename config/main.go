package config

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/sugandasu/ruru/nibirudb"
)

func Load() *Config {
	godotenv.Load(".env")
	app := app{
		Name:    os.Getenv("APP_NAME"),
		Version: os.Getenv("APP_VERSION"),
		Env:     os.Getenv("APP_ENV"),
		Domain:  os.Getenv("APP_DOMAIN"),
	}

	restPort, err := strconv.Atoi(os.Getenv("REST_PORT"))
	if err != nil {
		restPort = 8000
	}
	host := os.Getenv("REST_HOST")
	if host == "" {
		host = "localhost"
	}
	allowedOrigins := os.Getenv("REST_ALLOWED_ORIGINS")
	rest := rest{
		Host:           host,
		Port:           restPort,
		AllowedOrigins: strings.Split(allowedOrigins, ","),
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		dbPort = 5432
	}
	dbMaxIdleConnection, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	if err != nil {
		dbMaxIdleConnection = 10
	}
	dbMaxOpenConnection, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	if err != nil {
		dbMaxOpenConnection = 20
	}
	dbDebugMode, err := strconv.ParseBool(os.Getenv("DB_DEBUG_MODE"))
	if err != nil {
		dbDebugMode = false
	}

	db := nibirudb.Config{
		Driver:                os.Getenv("DB_DRIVER"),
		Host:                  os.Getenv("DB_HOST"),
		User:                  os.Getenv("DB_USER"),
		Password:              os.Getenv("DB_PASSWORD"),
		Name:                  os.Getenv("DB_NAME"),
		Port:                  dbPort,
		MaxConnectionLifeTime: os.Getenv("DB_MAX_CONNECTION_LIFE_TIME"),
		MaxConnectionIdleTime: os.Getenv("DB_MAX_CONNECTION_IDLE_TIME"),
		MaxIdleConnections:    dbMaxIdleConnection,
		MaxOpenConnections:    dbMaxOpenConnection,
		DebugMode:             dbDebugMode,
		Timeout:               os.Getenv("DB_TIMEOUT"),
		WriteTimeout:          os.Getenv("DB_WRITE_TIMEOUT"),
		ReadTimeout:           os.Getenv("DB_READ_TIMEOUT"),
		SSLMode:               os.Getenv("DB_SSL_MODE"),
		Timezone:              os.Getenv("DB_TIMEZONE"),
	}

	accessTokenDuration, err := time.ParseDuration(os.Getenv("JWT_ACCESS_TOKEN_DURATION"))
	if err != nil {
		accessTokenDuration = time.Hour * 24
	}
	refreshTokenDuration, err := time.ParseDuration(os.Getenv("JWT_REFRESH_TOKEN_DURATION"))
	if err != nil {
		refreshTokenDuration = time.Hour * 24 * 7
	}
	jwt := jwt{
		SecretKey:            os.Getenv("JWT_SECRET_KEY"),
		AccessTokenDuration:  accessTokenDuration,
		RefreshTokenDuration: refreshTokenDuration,
	}

	googleSmtpPort, err := strconv.Atoi(os.Getenv("GOOGLE_SMTP_PORT"))
	if err != nil {
		googleSmtpPort = 587
	}
	google := google{
		SMTPHost:        os.Getenv("GOOGLE_SMTP_HOST"),
		SMTPPort:        googleSmtpPort,
		NoReplyEmail:    os.Getenv("GOOGLE_NOREPLY_EMAIL"),
		NoReplyPassword: os.Getenv("GOOGLE_NOREPLY_PASSWORD"),
	}

	frontend := frontend{
		LoginEmailUrl:    os.Getenv("FRONTEND_LOGIN_EMAIL_URL"),
		ResetPasswordUrl: os.Getenv("FRONTEND_RESET_PASSWORD_URL"),
	}

	return &Config{
		App:      app,
		Rest:     rest,
		Jwt:      jwt,
		DB:       db,
		Google:   google,
		Frontend: frontend,
	}
}
