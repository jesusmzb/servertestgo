package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"servertestgo/database"
	"servertestgo/models"
	"servertestgo/services"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type responseEstimate struct {
	ValorBuy  string
	ValorSell string
	Vigency   time.Time
	error     string
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func Estimate(c echo.Context) error {
	fmt.Println("Se solicito la estimación")
	nameEstimate := c.Param("name")
	fmt.Println("nameEstimate:", nameEstimate)
	var response responseEstimate
	currecyaceptables := []string{"1INCH-USD", "1INCH-USDT", "AAVE-USDT", "ADA-USD", "ADA-USDT", "AGLD-USDT", "ALGO-USD", "ALGO-USDT", "ALPHA-USDT", "ANT-USDT", "APE-USDT", "API3-USDT", "APT-USDT", "ATOM-USD", "ATOM-USDT", "AVAX-USD", "AVAX-USDT", "AXS-USDT", "BADGER-USDT", "BAL-USDT", "BAND-USDT", "BAT-USDT", "BCH-USD", "BCH-USDT", "BICO-USDT", "BNB-USDT", "BNT-USDT", "BSV-USD", "BSV-USDT", "BTC-USDC", "BTC-USD", "BTC-USDT", "CELO-USDT", "CEL-USDT", "CFX-USDT", "CHZ-USDT", "COMP-USDT", "CRO-USDT", "CRV-USD", "CRV-USDT", "CSPR-USDT", "CVC-USDT", "DASH-USD", "DASH-USDT", "DOGE-USD", "DOGE-USDT", "DOME-USDT", "DORA-USDT", "DOT-USD", "DOT-USDT", "DYDX-USDT", "EGLD-USDT", "ENJ-USDT", "ENS-USDT", "EOS-USD", "EOS-USDT", "ETC-USD", "ETC-USDT", "ETH-USDC", "ETH-USD", "ETH-USDT", "ETHW-USDT", "FIL-USD", "FIL-USDT", "FITFI-USDT", "FLM-USDT", "FTM-USDT", "GALA-USDT", "GMT-USDT", "GODS-USDT", "GRT-USD", "GRT-USDT", "ICP-USDT", "IMX-USDT", "IOST-USDT", "IOTA-USDT", "JST-USDT", "KISHU-USDT", "KLAY-USDT", "KNC-USDT", "KSM-USD", "KSM-USDT", "LINK-USD", "LINK-USDT", "LOOKS-USDT", "LPT-USDT", "LRC-USDT", "LTC-USD", "LTC-USDT", "LUNA-USDT", "LUNC-USDT", "MANA-USD", "MANA-USDT", "MASK-USDT", "MATIC-USDT", "MINA-USDT", "MKR-USDT", "NEAR-USDT", "NEO-USD", "NEO-USDT", "NFT-USDT", "OMG-USDT", "ONT-USDT", "OP-USDT", "PEOPLE-USDT", "PERP-USDT", "QTUM-USDT", "REN-USDT", "RSR-USDT", "RVN-USDT", "SAND-USD", "SAND-USDT", "SHIB-USDT", "SLP-USDT", "SNX-USDT", "SOL-USD", "SOL-USDT", "STARL-USDT", "STORJ-USDT", "SUSHI-USD", "SUSHI-USDT", "SWEAT-USDT", "THETA-USD", "THETA-USDT", "TON-USDT", "TRB-USDT", "TRX-USD", "TRX-USDT", "UMA-USDT", "UNI-USD", "UNI-USDT", "USTC-USDT", "WAVES-USDT", "XCH-USDT", "XLM-USD", "XLM-USDT", "XMR-USD", "XMR-USDT", "XRP-USD", "XRP-USDT", "XTZ-USDT", "YFII-USDT", "YFI-USD", "YFI-USDT", "YGG-USDT", "ZEC-USD", "ZEC-USDT", "ZEN-USDT", "ZIL-USDT", "ZRX-USDT"}
	if len(nameEstimate) < 5 {
		response.error = "el par dado no es valido , debe estar entre " + strings.Join(currecyaceptables, " ")
		u, _ := json.Marshal(response)
		return c.String(http.StatusPartialContent, string(u))
	}
	if contains(currecyaceptables, nameEstimate) == false {
		response.error = "el par dado no es valido , debe estar entre " + strings.Join(currecyaceptables, " ")
		u, _ := json.Marshal(response)
		return c.String(http.StatusPartialContent, string(u))
	}
	fmt.Println(nameEstimate)
	myOrderBook := services.SolicitarBookOrder(nameEstimate)
	response.ValorBuy = myOrderBook.Data[0].Asks[0][0]
	response.ValorSell = myOrderBook.Data[0].Bids[0][0]
	response.Vigency = time.Now().Add(time.Second * 10)
	u, _ := json.Marshal(response)
	datastorage := models.OrderBookEntity{
		Currency:  nameEstimate,
		ValorBuy:  response.ValorBuy,
		ValorSell: response.ValorSell,
		Vigency:   response.Vigency,
	}
	database.OrderBookStore(datastorage)
	return c.String(http.StatusOK, string(u))
}
