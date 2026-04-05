package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string  `json:"message"`
	Result  float32 `json:"result"`
	Status  string  `json:"status"`
}

type Payload struct {
	Name   string  `json:"name"`
	Income float32 `json:"income"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var payload Payload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	response := Response{
		Message: payload.Name + " Sua renda líquida : ",
		Result:  payload.Income * float32(0.9),
		Status:  "success",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
