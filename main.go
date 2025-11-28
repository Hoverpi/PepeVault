package main

import (
	"log"
	// local imports
	// "PepeVault/automation"
	"PepeVault/config"
	"PepeVault/db"
	"PepeVault/routes"
)

func main() {
	cfg := config.LoadConfig()
	dbConnection := db.InitDB(cfg)
	router := routes.SetupRouter(dbConnection)

	log.Println("Server: http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
