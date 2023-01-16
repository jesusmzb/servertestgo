package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Op      string    `json:"op" validate:"required"`
	Id      string    `json:"id" validate:"required"`
	Args    *OrderArg `json:"args" validate:"required"`
	ExpTime string    `json:"expTime" validate:"required"`
}

type OrderArg struct {
	InstId       string `json:"instId" validate:"required"`
	TdMode       string `json:"tdMode" validate:"required"`
	Ccy          string `json:"ccy"`
	ClOrdId      string `json:"clOrdId"`
	Tag          string `json:"tag"`
	Side         string `json:"side"`
	PosSide      string `json:"posSide"`
	OrdType      string `json:"ordType" validate:"required"`
	Sz           string `json:"sz"`
	Px           string `json:"px"`
	ReduceOnly   bool   `json:"reduceOnly"`
	TgtCcy       string `json:"tgtCcy"`
	BanAmend     bool   `json:"banAmend"`
	QuickMgnType string `json:"quickMgnType"`
}

type OrderResponse struct {
	Id            string
	Response      string
	Vigency       time.Time
	ValueCurrency string
	errorResponse string
}

type OrderEntity struct {
	gorm.Model
	Request_okx  string
	Vigency      time.Time
	Response_okx string
}
