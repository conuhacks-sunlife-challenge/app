package plaid

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/plaid"
)

func Accounts(c *gin.Context) {

	if accessToken == "" {
		fmt.Println("Something went wrong!")
	}

	ctx := context.Background()

	accountsGetResp, _, err := client.PlaidApi.AccountsGet(ctx).AccountsGetRequest(
		*plaid.NewAccountsGetRequest(accessToken),
	).Execute()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accounts": accountsGetResp.GetAccounts(),
	})
}

func Transactions(c *gin.Context) {

	ctx := context.Background()

	const iso8601TimeFormat = "2006-01-02"
	startDate := time.Now().Add(-365 * 24 * time.Hour).Format(iso8601TimeFormat)
	endDate := time.Now().Format(iso8601TimeFormat)

	transactionsResp, _, err := client.PlaidApi.TransactionsGet(ctx).TransactionsGetRequest(
		*plaid.NewTransactionsGetRequest(accessToken, startDate, endDate),
	).Execute()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Transactions": transactionsResp.Transactions,
		"Accounts":     transactionsResp.Accounts,
	})
}
