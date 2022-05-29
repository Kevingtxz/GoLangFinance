package transaction_http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"transaction"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var layout = "2006-01-02T15:04:05"
	dateCreated, _ := time.Parse(layout, "2022-05-28T10:00:00")
	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:    "Salário",
			Amount:   1200.0,
			Type:     0,
			CreateAt: dateCreated,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

func CreateTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &res)

	fmt.Println(res)
}
