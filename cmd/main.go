package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	// db connection
	DSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var db *sqlx.DB
	var err error

	for range 5 {
		db, err = sqlx.Connect("postgres", DSN)
		if err != nil {
			log.Println("trying to connect with database....")
		}
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("database connection failed")

	}

	fmt.Println("database connect successfully")

	//dependency injection
	r := gin.Default()
	Routes(r)
	r.Run(":8080")

}
