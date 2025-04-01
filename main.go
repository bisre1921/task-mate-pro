package main

import (
	"fmt"
	"log"

	"github.com/bisre1921/task-master-pro/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	client, err := config.ConnectMongoDB(cfg)
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB!")

	defer func() {
		if err = client.Disconnect(nil); err != nil {
			log.Fatalf("Error disconnecting MongoDB: %v", err)
		}
		fmt.Println("MongoDB connection closed.")
	}()
}
