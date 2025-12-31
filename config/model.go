package config

import (
	"time"

	"github.com/sugandasu/ruru/nibirudb"
)

type Config struct {
	App      app
	Rest     rest
	Jwt      jwt
	DB       nibirudb.Config
	Google   google
	Frontend frontend
}

type app struct {
	Name    string
	Version string
	Env     string
	Domain  string
}

type rest struct {
	Host           string
	Port           int
	AllowedOrigins []string
	AllowedHeaders []string
}

type jwt struct {
	SecretKey            string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

type google struct {
	SMTPHost        string
	SMTPPort        int
	NoReplyEmail    string
	NoReplyPassword string
}

type frontend struct {
	LoginEmailUrl    string
	ResetPasswordUrl string
}
