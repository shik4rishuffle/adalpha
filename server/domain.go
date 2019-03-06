package main

import "time"

// A Portfolio is a collection of accounts.
type Portfolio struct {
	Name     string     `json:"name"`
	Accounts []*Account `json:"accounts"`
}

// Account holds assets. Trades are placed against accounts.
type Account struct {
	Name     string             `json:"name"`
	Holdings map[string]Decimal `json:"holdings"`
	Trades   []*Trade           `json:"trades"`
}

// Decimal is Unscaledx10^[Exponent]
type Decimal struct {
	Unscaled int `json:"unscaled"`
	Exponent int `json:"exponent"`
}

// NewDecimal creates a Decimal.
func NewDecimal(unscaled, exponent int) Decimal {
	return Decimal{Unscaled: unscaled, Exponent: exponent}
}

// A Fund is a type of asset, which can be bought and sold.
type Fund struct {
	ISIN         string        `json:"isin"`
	Name         string        `json:"name"`
	CurrentPrice int           `json:"current-price"`
	History      []FundHistory `json:"history"`
}

// FundHistory is the fund price at a specific point in time.
type FundHistory struct {
	Date  time.Time `json:"date"`
	Price int       `json:"price"`
}

// TradeStatus varies throughout the lifetime of a trade.
type TradeStatus int

const (
	// TradePending is the initial state of trades
	TradePending TradeStatus = iota + 1

	// TradeCancelled means the trade will not take place.
	TradeCancelled

	// TradeContracted means the trade has a legally binding agreement.
	TradeContracted

	// TradeSettled means the assets now belong to the buyer and the money has been sent to the seller.
	TradeSettled
)

// A Trade is placed on an account. The amount is in pence GBP, 1 == £0.01.
type Trade struct {
	FundName string      `json:"fundName"`
	Amount   int         `json:"amount"`
	Status   TradeStatus `json:"status"`
}
