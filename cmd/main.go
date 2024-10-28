package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/api/solver", func(w http.ResponseWriter, r *http.Request) {
		// CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Contorl-ALlow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		msg := map[string]string{"message": "Hello from the solver"}
		json.NewEncoder(w).Encode(msg)
	})

	http.ListenAndServe("localhost:8989", nil)
}
