package main

import (
	"fmt"
	"log"
	"os"
	"time"

	handler "github.com/deeep8250/SpendSense/internal/handlers"
	"github.com/deeep8250/SpendSense/internal/repositories"
	"github.com/deeep8250/SpendSense/internal/services"
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
			//this time is important without that loop will quickly finished before db start
			time.Sleep(2 * time.Second)
		}
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("database connection failed")

	}

	fmt.Println("database connect successfully")

	//dependency injection
	AuthRepository := repositories.NewAuthRepository(db)
	AuthService := services.NewAuthService(AuthRepository)
	AuthHandler := handler.NewAuthHanler(AuthService)
	r := gin.Default()
	Routes(r, AuthHandler)
	r.Run(":8080")

}
