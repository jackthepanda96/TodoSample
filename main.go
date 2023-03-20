package main

import (
	"database/sql"
	"log"
	"todo/config"
)

func InitSQL() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mydb")

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return db
}

func main() {

	// conn := InitSQL()
	cfg := config.InitConfig()
	conn := config.ConnSQL(*cfg)

	if conn == nil {
		log.Fatalln("Tidak bisa membuat koneksi")
	}
}
