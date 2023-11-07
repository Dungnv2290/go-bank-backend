package main

import (
	"github.com/dungnguyen/go-bank-backend/api"
	"github.com/dungnguyen/go-bank-backend/database"
)

func main() {
	// Do migration
	// migrations.Migrate()

	// Init database
	database.InitDatabase()
	api.StartApi()
}
