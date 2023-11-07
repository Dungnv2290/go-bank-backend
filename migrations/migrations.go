package migrations

import (
	"github.com/dungnguyen/go-bank-backend/database"
	"github.com/dungnguyen/go-bank-backend/helpers"
	"github.com/dungnguyen/go-bank-backend/interfaces"
)

func createAccounts() {
	users := &[2]interfaces.User{
		{Username: "Martin", Email: "martin@matin.com"},
		{Username: "Michael", Email: "michael@michael.com"},
	}
	for _, user := range users {
		generatedPassword := helpers.HashAndSalt([]byte(user.Username))
		user := &interfaces.User{Username: user.Username, Email: user.Email, Password: generatedPassword}
		database.DB.Create(user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(user.Username + "'s" + " account"), Balance: uint(10000), UserID: user.ID}
		database.DB.Create(account)
	}
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	Transaction := &interfaces.Transaction{}
	database.DB.AutoMigrate(User, Account, Transaction)

	createAccounts()
}
