package main

import (
	"fmt"
	"log"
	"os"
	"servertestgo/database"
	"servertestgo/server"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Servicio de OKX por Jesús Zarate")

	//cargamos el enviroment
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %\n", err)
	}
	//traemos los datos del enviroment
	PORT := os.Getenv("PORT")
	fmt.Println("Puerto del servidor:", PORT)
	OKX_API_URL := os.Getenv("OKX_API_URL")
	fmt.Println("API OKX:", OKX_API_URL)
	DATABASE_NAME := os.Getenv("DATABASE_NAME")
	database.InitDatabase(DATABASE_NAME)
	server.InitServer(PORT)
}

// func MarketPrice() {
// 	fmt.Println("Calling API...")
// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", "https://www.okx.com/api/v5/public/mark-price?instType=SWAP&instId", nil)
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	req.Header.Add("Accept", "application/json")
// 	req.Header.Add("Content-Type", "application/json")
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	defer resp.Body.Close()
// 	bodyBytes, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	responseObject := OrderBookEntity{}
// 	//fmt.Println(string(bodyBytes))
// 	erro := json.Unmarshal(bodyBytes, &responseObject)
// 	if erro != nil {
// 		fmt.Print(erro.Error())
// 	}
// 	fmt.Println("API Response as struct \n", responseObject)
// }

// func postOrder(method, path, body string) {
// 	fmt.Println("Calling API...")
// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, os.Getenv("OKX_API_URL")+path+body, bytes.NewBuffer([]byte("{'id': '1512','op': 'order','args': [{'side': 'buy','instId': 'BTC-USDT','tdMode': 'isolated','ordType': 'market','sz': '100'}]}")))
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	req.Header.Add("Accept", "application/json")
// 	req.Header.Add("Content-Type", "application/json")

// 	req.Header.Add("OK-ACCESS-KEY", os.Getenv("apikey"))

// 	timestamp, sign := sign(method, path, "{'id': '1512','op': 'order','args': [{'side': 'buy','instId': 'BTC-USDT','tdMode': 'isolated','ordType': 'market','sz': '100'}]}")

// 	req.Header.Add("OK-ACCESS-TIMESTAMP", timestamp)
// 	req.Header.Add("OK-ACCESS-PASSPHRASE", os.Getenv("PASSPHRASE"))
// 	req.Header.Add("OK-ACCESS-SIGN", sign)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	defer resp.Body.Close()
// 	bodyBytes, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	var responseObject OrderBook
// 	fmt.Println(string(bodyBytes))
// 	json.Unmarshal(bodyBytes, &responseObject)
// 	fmt.Println("API Response as struct \n", responseObject)
// }

// func sign(method, path, body string) (string, string) {
// 	format := "2006-01-02T15:04:05.999Z07:00"
// 	t := time.Now().UTC().Format(format)
// 	ts := fmt.Sprint(t)
// 	fmt.Println("Timestamp:" + ts)
// 	s := ts + method + path + body
// 	fmt.Println("sign:" + s)
// 	p := []byte(s)
// 	h := hmac.New(sha256.New, []byte(os.Getenv("secretkey")))
// 	h.Write(p)
// 	return ts, base64.StdEncoding.EncodeToString(h.Sum(nil))
// }

// type Order struct {
// 	InstID      string    `json:"instId"`
// 	Ccy         string    `json:"ccy"`
// 	OrdID       string    `json:"ordId"`
// 	ClOrdID     string    `json:"clOrdId"`
// 	TradeID     string    `json:"tradeId"`
// 	Tag         string    `json:"tag"`
// 	Category    string    `json:"category"`
// 	FeeCcy      string    `json:"feeCcy"`
// 	RebateCcy   string    `json:"rebateCcy"`
// 	Px          float64   `json:"px"`
// 	Sz          int64     `json:"sz"`
// 	Pnl         float64   `json:"pnl"`
// 	AccFillSz   int64     `json:"accFillSz"`
// 	FillPx      float64   `json:"fillPx"`
// 	FillSz      int64     `json:"fillSz"`
// 	FillTime    float64   `json:"fillTime"`
// 	AvgPx       float64   `json:"avgPx"`
// 	Lever       float64   `json:"lever"`
// 	TpTriggerPx float64   `json:"tpTriggerPx"`
// 	TpOrdPx     float64   `json:"tpOrdPx"`
// 	SlTriggerPx float64   `json:"slTriggerPx"`
// 	SlOrdPx     float64   `json:"slOrdPx"`
// 	Fee         float64   `json:"fee"`
// 	Rebate      float64   `json:"rebate"`
// 	State       string    `json:"state"`
// 	TdMode      string    `json:"tdMode"`
// 	PosSide     string    `json:"posSide"`
// 	Side        string    `json:"side"`
// 	OrdType     string    `json:"ordType"`
// 	InstType    string    `json:"instType"`
// 	TgtCcy      string    `json:"tgtCcy"`
// 	UTime       time.Time `json:"uTime"`
// 	CTime       time.Time `json:"cTime"`
// }
