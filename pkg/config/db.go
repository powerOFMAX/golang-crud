package config

import (
	"context"
	"log"
	"time"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"app/pkg/controllers"
	"github.com/joho/godotenv"
)

func getEnvValue(key string) string {
  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}

func Connect() {
	// Database Config

	// Get .Env
	dbUser := getEnvValue("DB_USER")
	dbPw := getEnvValue("DB_PW")
	dbUrl := getEnvValue("DB_URL")
	dbName := getEnvValue("DB_NAME")

	clientOptions := options.Client().ApplyURI("mongodb+srv://" + dbUser+ ":"+ dbPw + "@" + dbUrl)
	client, err := mongo.NewClient(clientOptions)

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//To close the connection at the end
	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	db := client.Database(dbName)
	controllers.MessageCollection(db)
	return
}