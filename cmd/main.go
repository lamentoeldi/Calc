package main

import (
	"Calc/internal/handlers"
	"Calc/internal/middleware"
	"fmt"
	"net/http"
)

func main() {
	// Let us go then, you and i...
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", handlers.CalcHandler)
	if err := http.ListenAndServe(":8080", middleware.LogMiddleware(mux)); err != nil {
		fmt.Println(err)
	}
}
