package transactions

import (
	"github.com/dungnguyen/go-bank-backend/database"
	"github.com/dungnguyen/go-bank-backend/helpers"
	"github.com/dungnguyen/go-bank-backend/interfaces"
)

func CreateTransaction(From uint, To uint, Amount int) {
	transaction := &interfaces.Transaction{From: From, To: To, Amount: Amount}
	database.DB.Save(transaction)
}

func GetTransactionByAccount(id uint) []interfaces.ResponseTransaction {
	transactions := []interfaces.ResponseTransaction{}
	database.DB.Table("transactions").Select("id, transactions.from, transactions.to, amount").Where(interfaces.Transaction{From: id}).Or(interfaces.Transaction{To: id}).Scan(&transactions)
	return transactions
}

func GetMyTransactions(id string, jwt string) map[string]interface{} {
	// Validate JWT
	isValid := helpers.ValidateToken(id, jwt)
	if isValid {
		// Find and return transactions
		accounts := []interfaces.ResponseAccount{}
		database.DB.Table("accounts").Select("id, name, balance").Where("user_id = ?", id).Scan(&accounts)

		transactions := []interfaces.ResponseTransaction{}
		for i := 0; i < len(accounts); i++ {
			accTransactions := GetTransactionByAccount(accounts[i].ID)
			transactions = append(transactions, accTransactions...)
		}

		var response = map[string]interface{}{"message": "all is fine"}
		response["data"] = transactions
		return response
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}
