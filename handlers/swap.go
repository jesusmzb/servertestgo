package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"servertestgo/database"
	"servertestgo/models"
	"servertestgo/services"
	"time"

	"github.com/labstack/echo/v4"
)

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
	//al crear la orden le ponemos los diez segundos de vigencia que solicita el sistema
	myorder := models.Order{Id: id, Op: "order", Args: &myorderarg, ExpTime: string((time.Now().Add(time.Second*10).UnixNano() / int64(time.Millisecond)))}

	responseData := services.PlaceOrder(myorder)
	fmt.Println(responseData)

	requestOKXString, _ := json.Marshal(myorder)
	responseOKXString, _ := json.Marshal(responseData)

	datastorage := models.OrderEntity{Request_okx: string(requestOKXString), Vigency: time.Now(), Response_okx: string(responseOKXString)}
	database.OrderStore(datastorage)
	u, _ := json.Marshal(datastorage)

	return c.String(http.StatusOK, string(u))

}

func SwapAll(c echo.Context) error {
	fmt.Println("Se solicito el swapAll")
	datos := database.OrderAll()
	alldata, _ := json.Marshal(datos)
	return c.String(http.StatusOK, string(alldata))
}
