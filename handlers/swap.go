package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"servertestgo/models"
	"servertestgo/services"

	"github.com/labstack/echo/v4"
)

type swapResponse struct {
	request  string
	response string
}

func Swap(c echo.Context) error {

	fmt.Println("Se solicito el swap")
	var jsonmap map[string]interface{} = map[string]interface{}{}
	if err := c.Bind(&jsonmap); err != nil {
		return err
	}

	var id string
	var ok bool
	if x, found := jsonmap["id"]; found {
		if id, ok = x.(string); !ok {
			fmt.Println("id:", id)
		}
	} else {
		return c.String(http.StatusBadRequest, `{"message": "Se requiere el Id"}`)
	}

	var side string
	if x, found := jsonmap["side"]; found {
		if side, ok = x.(string); !ok {
			fmt.Println("side:", side)
		}
	} else {
		return c.String(http.StatusBadRequest, `{"message": "Se requiere el side"}`)
	}
	var sz string
	if x, found := jsonmap["sz"]; found {
		if sz, ok = x.(string); !ok {
			fmt.Println("sz:", sz)
		}
	} else {
		return c.String(http.StatusBadRequest, `{"message": "Se requiere el sz"}`)
	}
	sideaceptables := []string{"buy", "sell"}
	if contains(sideaceptables, side) == false {
		return c.String(http.StatusBadRequest, `{"message": "Se requiere el side entre buy o sell"}`)
	}
	//datos validados creamos la orden
	myorderarg := models.OrderArg{Side: side, InstId: "ETH-USD", TdMode: "isolated", OrdType: "market", Sz: sz}
	myorder := models.Order{Id: id, Op: "order", Args: &myorderarg}

	responseData := services.PlaceOrder(myorder)

	requeststring, _ := json.Marshal(myorder)

	responsestring, _ := json.Marshal(responseData)

	u, _ := json.Marshal(swapResponse{request: string(requeststring), response: string(responsestring)})

	return c.String(http.StatusOK, string(u))

}

func SwapAll(c echo.Context) error {
	fmt.Println("Se solicito el swapAll")

	// if len(id) < 5 || len(side) < 5 || len(sz) < 5 {
	// 	response.ErrorResponse = "se requiere id , side , sz "
	// 	u, _ := json.Marshal(response)
	// 	return c.String(http.StatusPartialContent, string(u))
	// }

	// fmt.Println(nameEstimate)
	// myOrderBook := services.PlaceOrder(nameEstimate)
	// response.ValorBuy = myOrderBook.Data[0].Asks[0][0]
	// response.ValorSell = myOrderBook.Data[0].Bids[0][0]
	// response.Vigency = time.Now().Add(time.Second * 10)
	// u, _ := json.Marshal(response)
	// datastorage := models.OrderBookEntity{
	// 	Currency:  nameEstimate,
	// 	ValorBuy:  response.ValorBuy,
	// 	ValorSell: response.ValorSell,
	// 	Vigency:   response.Vigency,
	// }
	// database.OrderBookStore(datastorage)
	return c.String(http.StatusOK, " string(u)")
}
