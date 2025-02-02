package main

import (
	"os"

	"github.com/joho/godotenv"

	"server/Database"
	"server/plaid"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	db Database.DatabaseInstance = Database.DatabaseInstance{}
)
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	Plaid.Init()

	godotenv.Load()
	mongodb_user := os.Getenv("DATABASE_USERNAME")
	mongodb_password := os.Getenv("DATABASE_PASSWORD")
	mongodb_uri := os.Getenv("DATABASE_URI")

	db = Database.Connect(mongodb_user, mongodb_password, mongodb_uri)
}

type User struct {
    Email string `bson:"_id"`
    Password string `bson:"password"`
    FirstName string `bson:"first_name"`
    LastName string `bson:"last_name"`
}

func newUserHandler(ctx *gin.Context) {
	user := User{}
	err := ctx.BindJSON(&user)
	check(err)

	already_exists, err := db.CheckUser(user.Email)
	check(err)

	if already_exists != nil {
		ctx.AbortWithStatus(400)
	}

	err = db.AddUser(user.Email, user.Password, user.FirstName, user.LastName)
	check(err)

	ctx.Status(200)
}

func authenticationHandler(ctx *gin.Context) {
}

func main() {

	args := os.Args[1:]
	if len(args) > 0 && args[0] == "--init" {
		db.ForceDropRestart()
	}

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("../frontend/dist", true)))
	r.POST("api/createLinkToken", Plaid.CreateLinkToken)
	r.POST("api/getAccessToken", Plaid.GetAccessToken)
	r.POST("api/newUser", newUserHandler)
	r.POST("api/accounts", accounts)
	r.POST("api/transactions", transactions)
	r.Run()
}
