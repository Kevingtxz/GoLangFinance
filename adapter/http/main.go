package http

import (
	"net/http"

	actuator "github.com/Kevingtxz/GoLangFinance/adapter/http/actuator"
	transaction "github.com/Kevingtxz/GoLangFinance/adapter/http/transaction"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransactions)

	http.HandleFunc("/health", actuator.Health)

	http.ListenAndServe(":8080", nil)
}
