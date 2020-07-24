package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gitlab.com/tego-partner/kardiachain/kai-auth/pkg/auth"
	"gitlab.com/tego-partner/kardiachain/kai-auth/third_party/mongodb"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	mongodbURI := os.Getenv("DB_URI")
	key := os.Getenv("AES_SECRET_KEY")
	port := os.Getenv("PORT")

	fmt.Println(key)

	mongodb.Connect(mongodbURI)
	if err := auth.NewAuthServer(key).Run(port); err != nil {
		log.Fatal("Error starting Authentication server")
	}
}
