package main

import (
	"Calc/pkg/calc"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// Lasciate ogne speranza, voi ch’entrate
	http.HandleFunc("/api/v1/calculate", func(w http.ResponseWriter, r *http.Request) {
		var requestBody struct {
			Expression string `json:"expression"`
		}
		type errorBody struct {
			Error string `json:"error"`
		}

		// We don't want it to panic, do we?
		defer func() {
			err := recover()
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(&errorBody{
					Error: "Internal server Error",
				})
			}
		}()

		// We don't want it to process the wrong methods, do we?
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(&errorBody{
				Error: "Method Not Allowed",
			})
			return
		}

		// We don't want it to process requests without body, do we?
		if r.Body == nil {
			panic("empty request body")
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			panic(err)
			return
		}

		// Here we make the Calc function to process the expression
		res, err := calc.Calc(requestBody.Expression)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			if err := json.NewEncoder(w).Encode(errorBody{Error: "Expression is not valid"}); err != nil {
				panic(err)
			}
			return
		}

		// This whole thing should send the response if all go or panic if not
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(struct {
			Result string `json:"result"`
		}{
			Result: fmt.Sprintf("%.2f", res),
		}); err != nil {
			panic("error encoding response")
		}
	})

	// Let us go then, you and i...
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
