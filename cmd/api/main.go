package main

import (
	"fmt"
	"log"
	"net/http"

	"go-practice2/internal/handlers"
	"go-practice2/internal/middleware"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/user", middleware.AuthMiddleware(http.HandlerFunc(handlers.UserHandler)))

	fmt.Println("Server is running 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
