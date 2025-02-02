package main

import (
	"os"
	"server/Database"
	"github.com/joho/godotenv"
)

func pingDb() {
    mongodb_user := os.Getenv("DATABASE_USERNAME")
    mongodb_password := os.Getenv("DATABASE_PASSWORD")


    db := Database.Connect(mongodb_user, mongodb_password)
    db.Ping()
}

func main() {
    err := godotenv.Overload()
    if err != nil {
        panic("Could not load dotenv")
    }

    mongodb_user := os.Getenv("DATABASE_USERNAME")
    mongodb_password := os.Getenv("DATABASE_PASSWORD")

    db := Database.Connect(mongodb_user, mongodb_password)

    args := os.Args[1:]
    if len(args) > 0 && args[0] == "--init"{
	db.ForceDropRestart()
    }

    db.Ping()


}
