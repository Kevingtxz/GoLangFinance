package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	transaction "github.com/Kevingtxz/GoLangFinance/model/transaction"
	"github.com/Kevingtxz/GoLangFinance/util"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:    "Salário",
			Amount:   1200.0,
			Type:     0,
			CreateAt: util.StringToDate("2022-01-01T01:01:01"),
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
