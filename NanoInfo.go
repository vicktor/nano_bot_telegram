package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var client http.Client

const apiKey string = "578c64cb-ddb4-4431-8514-ca265f4cfbe6"
const botApiKey string = "1480978128:AAGXt4FIIVQQyUh4oacRlPQqaDr2XEroeNA"
const chatId string = "-549731810"

type CryptoCurrency struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Symbol         string    `json:"symbol"`
	Slug           string    `json:"slug"`
	NumMarketPairs int       `json:"num_market_pairs"`
	DateAdded      time.Time `json:"date_added"`
	Tags           []struct {
		Slug     string `json:"slug"`
		Name     string `json:"name"`
		Category string `json:"category"`
	} `json:"tags"`
	MaxSupply         int         `json:"max_supply"`
	CirculatingSupply float64     `json:"circulating_supply"`
	TotalSupply       float64     `json:"total_supply"`
	IsActive          int         `json:"is_active"`
	Platform          interface{} `json:"platform"`
	CmcRank           int         `json:"cmc_rank"`
	IsFiat            int         `json:"is_fiat"`
	LastUpdated       time.Time   `json:"last_updated"`
	Quote             struct {
		EUR struct {
			Price            float64   `json:"price"`
			Volume24H        float64   `json:"volume_24h"`
			PercentChange1H  float64   `json:"percent_change_1h"`
			PercentChange24H float64   `json:"percent_change_24h"`
			PercentChange7D  float64   `json:"percent_change_7d"`
			PercentChange30D float64   `json:"percent_change_30d"`
			PercentChange60D float64   `json:"percent_change_60d"`
			PercentChange90D float64   `json:"percent_change_90d"`
			MarketCap        float64   `json:"market_cap"`
			LastUpdated      time.Time `json:"last_updated"`
		} `json:"EUR"`
	} `json:"quote"`
}
type NanoResponse struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
	} `json:"status"`
	Data struct {
		Nano    CryptoCurrency `json:"1567"`
		Cardano CryptoCurrency `json:"2010"`
	} `json:"data"`
}

func sendToTelegram(currency string, price float64, amount int32) {

	botURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botApiKey)
	total := price * float64(amount)
	_, err := http.PostForm(botURL, url.Values{"chat_id": {chatId}, "text": {fmt.Sprintf("%s: Current quote: %.5f, current amount for %d coins: %.5f", currency, price, amount, total)}})

	if err != nil {
		fmt.Printf("Error posting to Telegram %s", err.Error())
	}

}

func requestNano() error {

	url := fmt.Sprintf("https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest?id=1567,2010&convert=EUR&CMC_PRO_API_KEY=%s", apiKey)
	request, err := http.NewRequest(
		http.MethodGet,
		url,
		nil)

	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}
	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	var nanoResponse NanoResponse

	err = json.NewDecoder(response.Body).Decode(&nanoResponse)
	if err != nil {
		fmt.Printf("Error decoding body. %s", err.Error())
		return err
	}

	sendToTelegram("Nano", nanoResponse.Data.Nano.Quote.EUR.Price, 19)
	sendToTelegram("Cardano", nanoResponse.Data.Cardano.Quote.EUR.Price, 20)
	return nil
}

func main() {
	requestNano()
}
