package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvProp struct {
	GetMongoUrl func() string
	GetUserSecretKey func() string
	GetAdminSecretKey func() string
	GetUserSecretRefreshKey func() string
	GetPort func() string
	GetSMTPserverUrl func() string
}

func VerifyEnvVariable() {
	err := godotenv.Load()
	if (err != nil) {
		log.Fatal("Error loading .env file")
	}
}

var Env = EnvProp{
	GetMongoUrl: func () string {
		return os.Getenv("MONGODB_URL");
	},
	GetUserSecretKey: func () string {
		return os.Getenv("USER_SECRET_KEY");
	},
	GetAdminSecretKey: func () string {
		return os.Getenv("ADMIN_SECRET_KEY");
	},
	GetUserSecretRefreshKey: func () string {
		return os.Getenv("USER_SECRET_REFRESH_KEY");
	},
	GetPort: func () string {
		return os.Getenv("PORT");
	},
	GetSMTPserverUrl: func() string {
		return os.Getenv("SMTP_SERVER");
	},
}