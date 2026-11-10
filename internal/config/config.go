package config 

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	Port 		string
	DatabaseURL string
	JWTSecret 	string
}

func Load() *Config {
	//godotenv = go+dot+env -> go library ditulis untuk bahasa go ||  dot file yang dibaca yaitu .env  ||  env = environment (variabel lingkungan seperti PORT dan DATABASE_URL)
	_ = godotenv.Load()

	return &Config{
		Port:			getEnv("PORT", "8080"),
		DatabaseURL:	getEnv("DATABASE_URL", ""),
		JWTSecret: 		getEnv("JWT_SECRET", ""),
	}
}

// getEnv mengambil nilai dari variabel lingkungan atau mengembalikan nilai default jika variabel tidak ditemukan
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}