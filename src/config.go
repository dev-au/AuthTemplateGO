package src

import (
	"AuthTemplate/src/models"
	"encoding/base64"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var Config envData

type envData struct {
	DBHost          string `env:"DB_HOST" envDefault:"localhost"`
	DBPort          string `env:"DB_PORT" envDefault:"5432"`
	DBUser          string `env:"DB_USER"`
	DBPassword      string `env:"DB_PASSWORD"`
	DBName          string `env:"DB_NAME"`
	JWTSecret       string `env:"JWT_SECRET"`
	JWTExpireDays   int    `env:"JWT_EXPIRE_DAYS" envDefault:"60"`
	Port            string `env:"PORT" envDefault:"8000"`
	GmailAccount    string `env:"GMAIL_ACCOUNT"`
	GmailPassword   string `env:"GMAIL_PASSWORD"`
	RedisUrl        string `env:"REDIS_URL"`
	ApplicationUrl  string `env:"APPLICATION_URL"`
	EncryptionKey   []byte
	ApplicationMode string `env:"APPLICATION_MODE"`
}

func (envData) SetupEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	expireDays, err := strconv.Atoi(os.Getenv("JWT_EXPIRE_DAYS"))
	if err != nil {
		log.Fatalf("Invalid JWT_EXPIRE_DAYS: %v", err)
	}

	encryptionKey := os.Getenv("ENCRYPTION_KEY")

	keyByte, err := base64.StdEncoding.DecodeString(encryptionKey)
	if err != nil {
		log.Fatalf("Invalid encryption key: %v", err)
	}

	Config = envData{
		DBHost:          os.Getenv("DB_HOST"),
		DBPort:          os.Getenv("DB_PORT"),
		DBUser:          os.Getenv("DB_USER"),
		DBPassword:      os.Getenv("DB_PASSWORD"),
		DBName:          os.Getenv("DB_NAME"),
		JWTSecret:       os.Getenv("JWT_SECRET"),
		JWTExpireDays:   expireDays,
		Port:            os.Getenv("PORT"),
		GmailAccount:    os.Getenv("GMAIL_ACCOUNT"),
		GmailPassword:   os.Getenv("GMAIL_PASSWORD"),
		RedisUrl:        os.Getenv("REDIS_URL"),
		ApplicationUrl:  os.Getenv("APPLICATION_URL"),
		EncryptionKey:   keyByte,
		ApplicationMode: os.Getenv("APPLICATION_MODE"),
	}
}

var Models = []interface{}{
	&models.User{},
	&models.Role{},
}
