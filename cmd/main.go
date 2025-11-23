package main

import (
	"log"

	"github.com/joho/godotenv"

	"avito_internship_PR/internal/config"
	"avito_internship_PR/internal/db"
	"avito_internship_PR/internal/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
		log.Println("no .env file found")
	}

	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	conn, err := db.NewPostgresConnection(&cfg.DB)
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(conn.GormDB)
	if err != nil {
		panic(err)
	}
	r := router.SetupRouter(conn.GormDB, cfg)
	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
