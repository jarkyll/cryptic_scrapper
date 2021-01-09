package main

import (
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-boilerplate/routers"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}

	err = mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if err != nil {
		log.Fatal("Error connecting to db.")
	}



	server := routers.NewRouter()



	server.Run("localhost:8080")
}



