package handlers

import (
	"Calc/pkg/calc"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	// Lasciate ogne speranza, voi ch’entrate

	// We won't let the memory leak occur
	defer r.Body.Close()

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
			log.Printf("[ERROR]: %v", err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(w).Encode(errorBody{
				Error: "Internal server Error",
			}); err != nil {
				panic(err)
			}
		}
	}()

	// We don't want it to process the wrong methods, do we?
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		if err := json.NewEncoder(w).Encode(errorBody{
			Error: "Method Not Allowed",
		}); err != nil {
			panic(err)
		}
		return
	}

	// We don't want it to process requests without body, do we?
	if r.Body == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errorBody{
			Error: "Bad request",
		}); err != nil {
			panic(err)
		}
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errorBody{
			Error: "Bad request",
		}); err != nil {
			panic(err)
		}
		return
	}

	// Here we make the Calc function to process the expression
	res, err := calc.Calc(requestBody.Expression)
	if err != nil {
		log.Printf("[ERROR]: %v", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(errorBody{
			Error: "Expression is not valid",
		}); err != nil {
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
		panic(err)
	}
}
