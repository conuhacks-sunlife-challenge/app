package plaid

import (
	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/plaid"
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
	environments                         = map[string]plaid.Environment{
		"sandbox": plaid.Sandbox,
	}
	itemID      = ""
	accessToken = ""
)

func Run() {
	r := gin.Default()
	r.POST("api/createLinkToken", CreateLinkToken)
	r.POST("api/getAccessToken", GetAccessToken)
	r.POST("api/accounts", accounts)
	r.POST("api/transactions", transactions)
	accessToken = "access-sandbox-fe46e356-a90b-4c2d-9b91-180db236f081"
	//accessToken = "access-sandbox-4853d14b-7e5f-4991-b566-2d464f6a01ae"
	r.Run()
}
