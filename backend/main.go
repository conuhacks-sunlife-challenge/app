package main

import (
	"os"
	"server/Database"
	"server/Plaid"
)

func pingDb() {
	mongodb_user := os.Getenv("DATABASE_USERNAME")
	mongodb_password := os.Getenv("DATABASE_PASSWORD")
	mongodb_uri := os.Getenv("DATABASE_URI")

	db := Database.Connect(mongodb_user, mongodb_password, mongodb_uri)
	db.Ping()
}

func main() {
	mongodb_user := os.Getenv("DATABASE_USERNAME")
	mongodb_password := os.Getenv("DATABASE_PASSWORD")
	mongodb_uri := os.Getenv("DATABASE_URI")

	db := Database.Connect(mongodb_user, mongodb_password, mongodb_uri)

	args := os.Args[1:]
	if len(args) > 0 && args[0] == "--init" {
		db.ForceDropRestart()
	}

	db.Ping()
	Plaid.Init()
	Plaid.Run()
}
