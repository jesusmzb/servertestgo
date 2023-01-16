package models

import "gorm.io/gorm"

// para manejar las book order, creamos las estructuras
type (
	OrderBook struct {
		Code string `json:"code"`
		Msg  string `json:"msg"`
		Data []data `json:"data"`
	}

	data struct {
		Asks [][]string `json:"asks"`
		Bids [][]string `json:"bids"`
		TS   string     `json:"ts"`
	}

	OrderBookEntity struct {
		gorm.Model
		DepthPrice      float64
		Size            float64
		LiquidatedOrder int
		OrderNumbers    int
	}
)
