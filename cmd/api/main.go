package main

import (
	"log"

	"github.com/dedehudianto12/bbs-backend/config"
	"github.com/dedehudianto12/bbs-backend/internal/shared/database"
)

func main(){
	cfg, err := config.Load()
	
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("Database connected")

	defer db.Close()

}