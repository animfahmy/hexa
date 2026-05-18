package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DBString   string
}

func LoadConfig() *Config {
	// Jalankan godotenv
	if err := godotenv.Load(); err != nil {
		log.Println("Peringatan: File .env tidak ditemukan, menggunakan env system")
	}

	return &Config{
		AppPort: getEnv("APP_PORT", ":8080"), // ":8080" adalah nilai default jika env kosong
		DBString:   getEnv("DB_DSN", ""),
	}
}

// Fungsi pembantu untuk memberikan nilai default jika env tidak diatur
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}