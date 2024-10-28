package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Sd struct {
	Vals []int `json:"sdVals"`
}

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

		if r.Method == "POST" {
			var sd Sd
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&sd)
			if err != nil {
				http.Error(w, "Could not decode JSON", http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			fmt.Println(sd.Vals)
		}

		msg := map[string]string{"message": "Hello from the solver"}
		json.NewEncoder(w).Encode(msg)
	})

	http.ListenAndServe("localhost:8989", nil)
}
