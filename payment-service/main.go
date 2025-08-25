package main

import (
	"encoding/json"
	"net/http"
)

type Payment struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
	UserID int     `json:"userId"`
}

var payments = []Payment{{ID: 1, Amount: 99.99, UserID: 1}}

func getPayments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}

func main() {
	http.HandleFunc("/payments", getPayments)
	http.ListenAndServe(":5002", nil)
}
