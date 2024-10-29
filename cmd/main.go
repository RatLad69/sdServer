package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Sd struct {
	Vals []int `json:"sdVals"`
	Size int   `json:"size"`
}

func splitSd(sd Sd) [][]int {
	numRows := sd.Size
	row := make([]int, numRows, numRows)
	valSquare := make([][]int, numRows, numRows)
	for i := 0; i < numRows; i++ {
		row = sd.Vals[numRows*i : numRows*i+numRows]
		valSquare[i] = row
	}

	return valSquare
}

func isLegalPlacement(sd [][]int, col int, row int, val int) bool {
	for i := 0; i < len(sd); i++ {
		if sd[row][i] == val {
			return false
		}
	}

	for i := 0; i < len(sd); i++ {
		if sd[i][col] == val {
			return false
		}
	}

	subSquareRow0 := row - row%3
	subSquareCol0 := col - col%3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if sd[subSquareRow0+i][subSquareCol0+j] == val {
				return false
			}
		}
	}

	return true
}

func sdSolve(sq *[][]int, size int, row int, col int) bool {
	if row == size-1 && col == size {
		return true
	}

	if col == size {
		row++
		col = 0
	}

	if (*sq)[row][col] > 0 {
		return sdSolve(sq, size, row, col+1)
	}

	for val := 1; val <= size; val++ {
		if isLegalPlacement(*sq, col, row, val) {
			(*sq)[row][col] = val
			if sdSolve(sq, size, row, col+1) {
				return true
			}
		}
		(*sq)[row][col] = 0
	}
	return false
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

			sdSquare := splitSd(sd)
			fmt.Println(sdSquare)
			fmt.Println("-------")
			if sdSolve(&sdSquare, sd.Size, 0, 0) {
				fmt.Println(sdSquare)
			} else {
				fmt.Println("Failed to solve")
			}
		}

		msg := map[string]string{"message": "Hello from the solver"}
		json.NewEncoder(w).Encode(msg)
	})

	http.ListenAndServe("localhost:8989", nil)
}
