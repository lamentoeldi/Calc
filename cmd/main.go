package main

import (
	"Calc/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	// Let us go then, you and i...
	http.HandleFunc("/api/v1/calculate", handlers.CalcHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
