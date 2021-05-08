package main

import (
	"fmt"
	"github.com/yuriytaranov/go-todo/internal/pkg/server/middleware"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	routerr "github.com/yuriytaranov/go-todo/internal/pkg/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env not found using system variables")
	}
	middleware.Connect()
	r := routerr.Router()

	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
