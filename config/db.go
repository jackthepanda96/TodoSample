package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnSQL(c AppConfig) *sql.DB {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		c.DBUsername,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName)

	db, err := sql.Open("mysql", conn)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return db
}
