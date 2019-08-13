package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/johnksv/sparebank1-personligklient/client/operations"
	"golang.org/x/oauth2"
	"gopkg.in/AlecAivazis/survey.v1"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	apiclient "github.com/johnksv/sparebank1-personligklient/client"
)

var DescriptionToMemo = map[string]string {
	"KIWI 471 BISLET THERESESGT 3 OSLO": "KIWI BISLETT",
	"KIWI 471 BISLET":                   "KIWI BISLETT",
	"VIPPS *RUTER AS":                   "VIPPS RUTER ",
	"REMA TORGGATA REMA 1000 TO OSLO":   "REMA Torggata",
	"JOKER ADAMSTUEN ULLEVÅLSVN 9 OSLO": "Joker Adamstuen",
	"KIWI 365 ST.OLA ST.OLAVSPLAS OSLO": "KIWI ST.OLAVSPLASS",
	"EXTRA H. HÅRFAGRES OSLO":           "Coop Extra H. Hårfagres",
	"JOKER KIRKEVN KIRKEVN. 90 OSLO.":   "Joker Kirkeveien",
}

func main() {
	var accountNumberFlag int
	flag.IntVar(&accountNumberFlag, "accountnumber", 0, "Account number to fetch transactions from.")
	flag.Parse()

	//_ = os.Setenv("DEBUG", "1")
	conf := &oauth2.Config{
		ClientID:     os.Getenv("ClientID"),
		ClientSecret: os.Getenv("ClientSecret"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.sparebank1.no/oauth/authorize",
			TokenURL: "https://api.sparebank1.no/oauth/token",
		},
	}

	url := conf.AuthCodeURL("1029384756",
		oauth2.SetAuthURLParam("finInst", os.Getenv("FinInst")),
		oauth2.SetAuthURLParam("redirect_uri", os.Getenv("RedirectURI")),
	)

	var accessToken = os.Getenv("AccessToken")
	if accessToken == "" {
		fmt.Printf("Visit the URL for the auth dialog: %v", url)
		fmt.Println()
		fmt.Println("Enter accessToken to be used.")
		if _, err := fmt.Scan(&accessToken); err != nil {
			log.Fatal(err)
		}
	}

	rt := httptransport.New(
		"api.sparebank1.no",
		"/open/personal/banking/accounts",
		[]string{"https"})
	rt.Consumers["application/vnd.sparebank1.v1+json"] = runtime.JSONConsumer()
	rt.Producers["application/vnd.sparebank1.v1+json"] = runtime.JSONProducer()


	client := apiclient.New(rt, nil)


	bearerTokenAuth := httptransport.BearerToken(accessToken)

	fmt.Println("Fetching accounts.")
	resp, err := client.Operations.GetAccounts(nil, bearerTokenAuth)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("ACCOUNTS:")
	var index = -1
	for i, account := range resp.Payload.Accounts {
		fmt.Println(i, ":",account.Name, "("+ *account.AccountNumber.Formatted +")" ,"-",*account.Owner.FirstName)
		if *account.AccountNumber.Value == fmt.Sprint(accountNumberFlag)  {
			index = i
		}
	}
	if index == -1{
		index = askWhichAccount()
	}
	account := resp.Payload.Accounts[index]
	fmt.Println("Fetching transactions from", account.Name, "("+ *account.AccountNumber.Formatted +")" )

	transactionParams := &operations.GetAccountTransactionsParams{
		AccountID: *account.ID,
	}
	duration, _ := time.ParseDuration("10s")
	transactionParams.SetTimeout(duration)

	transactions, err := client.Operations.GetAccountTransactions(transactionParams, bearerTokenAuth)

	if err != nil || transactions.Payload.Transactions == nil {
		log.Fatal(err)
	}

	transactionsToCSVFile(transactions, *account.Owner.FirstName)

}

func askWhichAccount() int {
	indexString := ""
	prompt := &survey.Input{Message: "Which account to proceed with (index)?"}
	err := survey.AskOne(prompt, &indexString, nil)
	if err != nil {
		log.Fatal(err)
	}
	index, err := strconv.Atoi(indexString)
	if err != nil {
		log.Fatal(err)
	}
	return index
}

func transactionsToCSVFile(transactions *operations.GetAccountTransactionsOK, accountOwner string) {
	dir, _ := os.Getwd()
	file := dir + "/sparebank1-transactions.csv"
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, _ = w.WriteString("Date,Payee,Category,Memo,Outflow,Inflow\n")
	for _, transaction := range transactions.Payload.Transactions {

		t, err := time.Parse("2006-01-02", transaction.AccountingDate.String())
		if err != nil {
			log.Fatal(err)
		}
		date := t.Format("02/01/2006")
		payee := accountOwner
		category := "" //TODO
		memo := mapToMemo(transaction.Description)

		var inflow string
		var outflow string
		if *transaction.Amount.Amount > 0 {
			inflow = fmt.Sprintf("%f", *transaction.Amount.Amount)
			outflow = ""
		} else {
			inflow = ""
			outflow = fmt.Sprintf("%f", math.Abs(*transaction.Amount.Amount))
		}
		_, _ = w.WriteString(date + ",")
		_, _ = w.WriteString(payee + ",")
		_, _ = w.WriteString(category + ",")
		_, _ = w.WriteString(memo + ",")
		_, _ = w.WriteString(outflow + ",")
		_, _ = w.WriteString(inflow + "\n")
	}
	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transactions fetched and stored to file:", file)
}

func mapToMemo(description string) string{
	value, ok := DescriptionToMemo[description]
	if ok {
		return value
	}

	return description
}
