/*
Written by Justus Languell 2021
Project to practice GO
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	"github.com/shopspring/decimal"
)

func init() {
	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
	os.Setenv(common.EnvApiKeyID, "PKEFMGH4Q9KG1RQ2AIKS")
	os.Setenv(common.EnvApiSecretKey, "fpUrKsMHkpHmqMJ6y7PjAN4yHnYQc8ytBhaTtst5")
	fmt.Printf("Running w/ credentials [%v %v]\n", common.Credentials().ID, common.Credentials().Secret)
}

func execBUY(symbol string, size float64) {
	alpaca.PlaceOrder(alpaca.PlaceOrderRequest{
		AssetKey:    &symbol,
		Qty:         decimal.NewFromFloat(size),
		Side:        alpaca.Buy,
		Type:        alpaca.Mrket,
		TimeInForce: alpaca.Day,
	})
}

// Host a static page
func route(path, page string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	http.HandleFunc(page, func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, string(data))
	})
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	//_, err := io.Copy(os.Stdout, r.Body)
	fmt.Println("WebHook Accepted")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	var data map[string]interface{}
	json.Unmarshal([]byte(body), &data)
	//[close:135.11 key:e783d60168b9b44d7414f4c72f7b0d5f price:135.14 sent:2021-02-08T14:34:20Z side:buy size:200 ticker:AAPL]

}

func main() {
	//route("index.html", "/")
	//http.HandleFunc("/webhook", handleWebhook)
	// Submit a market order to buy 1 share of Apple at market price
	symbol := "AAPL"
	alpaca.PlaceOrder(alpaca.PlaceOrderRequest{
		AssetKey:    &symbol,
		Qty:         decimal.NewFromFloat(1),
		Side:        alpaca.Buy,
		Type:        alpaca.Market,
		TimeInForce: alpaca.Day,
	})
	//execBUY("AAPL", 1)
	//log.Fatal(http.ListenAndServe(":8080", nil))

}a
