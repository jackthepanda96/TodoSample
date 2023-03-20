package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
}

func InitConfig() *AppConfig {
	res := readConfig()

	if res == nil {
		log.Println("Gagal melakukan setup")
		return nil
	}

	return res
}

func readConfig() *AppConfig {
	err := godotenv.Load(".env")
	res := AppConfig{}
	if err != nil {
		log.Println("Tidak bisa baca konfigurasi")
		return nil
	}
	res.DBUsername = os.Getenv("DBUsername")
	res.DBPassword = os.Getenv("DBPassword")
	res.DBHost = os.Getenv("DBHost")
	cnv, err := strconv.Atoi(os.Getenv("DBPort"))
	if err != nil {
		log.Println("Nilai port tidak valid")
		return nil
	}
	res.DBPort = cnv
	res.DBName = os.Getenv("DBName")

	return &res
}
