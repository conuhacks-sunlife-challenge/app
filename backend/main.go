package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/Database"
	"time"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	plaid "github.com/plaid/plaid-go/v31/plaid"
)

var (
	PLAID_CLIENT_ID                      = ""
	PLAID_SECRET                         = ""
	PLAID_ENV                            = ""
	PLAID_PRODUCTS                       = ""
	PLAID_COUNTRY_CODES                  = ""
	PLAID_REDIRECT_URI                   = ""
	APP_PORT                             = ""
	client              *plaid.APIClient = nil
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading the .env file %w", err)
	}

	PLAID_CLIENT_ID = os.Getenv("PLAID_CLIENT_ID")
	PLAID_SECRET = os.Getenv("PLAID_SECRET")

	if PLAID_CLIENT_ID == "" || PLAID_SECRET == "" {
		log.Fatal("Error parsing the API keys from .env file")
	}

	PLAID_ENV = os.Getenv("PLAID_ENV")
	PLAID_PRODUCTS = os.Getenv("PLAID_PRODUCTS")
	PLAID_COUNTRY_CODES = os.Getenv("PLAID_COUNTRY_CODES")
	PLAID_REDIRECT_URI = os.Getenv("PLAID_REDIRECT_URI")
	APP_PORT = os.Getenv("APP_PORT")

	// creating Plaid configuration
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", PLAID_CLIENT_ID)
	configuration.AddDefaultHeader("PLAID-SECRET", PLAID_SECRET)
	configuration.UseEnvironment(environments[PLAID_ENV])
	client = plaid.NewAPIClient(configuration)
}

func pingDb() {
	mongodb_user := os.Getenv("DATABASE_USERNAME")
	mongodb_password := os.Getenv("DATABASE_PASSWORD")
	mongodb_uri := os.Getenv("DATABASE_URI")

	db := Database.Connect(mongodb_user, mongodb_password, mongodb_uri)
	db.Ping()
}

var environments = map[string]plaid.Environment{
	"sandbox": plaid.Sandbox,
}

func main() {
	err := godotenv.Overload()
	if err != nil {
		panic("Could not load dotenv")
	}

	mongodb_user := os.Getenv("DATABASE_USERNAME")
	mongodb_password := os.Getenv("DATABASE_PASSWORD")
	mongodb_uri := os.Getenv("DATABASE_URI")

	db := Database.Connect(mongodb_user, mongodb_password, mongodb_uri)

	args := os.Args[1:]
	if len(args) > 0 && args[0] == "--init" {
		db.ForceDropRestart()
	}

	db.Ping()

	success, err := db.AddUser("a@a", "123", "John", "Doe")
	if err != nil {
		panic(err)
	}
	if !success {
		panic("Didn't add user!")
	}

	r := gin.Default()
	r.POST("api/createLinkToken", createLinkToken)
	r.POST("api/getAccessToken", getAccessToken)
	r.Run()
}

func createLinkToken(c *gin.Context) {
	ctx := context.Background()

	// Get the client_user_id by searching for the current user
	//user := "hey"
	clientUserId := time.Now().String()

	// Create a link_token for the given user
	request := plaid.NewLinkTokenCreateRequest("Plaid Test App", "en", []plaid.CountryCode{plaid.COUNTRYCODE_US}, *plaid.NewLinkTokenCreateRequestUser(clientUserId))
	request.SetProducts([]plaid.Products{plaid.PRODUCTS_AUTH})

	resp, _, err := client.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"link_token": resp.GetLinkToken(),
	})
}

func getAccessToken(c *gin.Context) {
	ctx := context.Background()

	var requestBody struct {
		PublicToken string `json:"public_token"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	publicToken := requestBody.PublicToken

	// exchange the public_token for an access_token
	exchangePublicTokenReq := plaid.NewItemPublicTokenExchangeRequest(publicToken)
	exchangePublicTokenResp, _, err := client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*exchangePublicTokenReq,
	).Execute()

	if err != nil {
		fmt.Println(err)
	}

	// to be saved in the DB
	accessToken := exchangePublicTokenResp.GetAccessToken()
	itemID := exchangePublicTokenResp.GetItemId()

	fmt.Println("access token: " + accessToken)
	fmt.Println("item ID: " + itemID)

	c.JSON(http.StatusOK, gin.H{"public_token_exchange": "complete"})
}
