package main

import (
	"fmt"
	"log"
	"os"
	"servertestgo/database"
	"servertestgo/server"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Servicio de OKX por Jesús Zarate")

	//cargamos el enviroment
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %\n", err)
	}
	//traemos los datos del enviroment
	PORT := os.Getenv("PORT")
	fmt.Println("Puerto del servidor:", PORT)
	OKX_API_URL := os.Getenv("OKX_API_URL")
	fmt.Println("API OKX:", OKX_API_URL)
	DATABASE_NAME := os.Getenv("DATABASE_NAME")
	database.InitDatabase(DATABASE_NAME)
	server.InitServer(PORT)
}
