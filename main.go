package main

import (
	"net/http"
	transaction_http
)

func main() {
	http.HandleFunc("/transactions", transaction_http.GetTransactions)
	http.HandleFunc("/transactions/create", transaction_http.CreateTransactions)
	http.ListenAndServe(":8080", nil)
}
