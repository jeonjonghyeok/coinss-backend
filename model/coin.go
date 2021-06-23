package model

type Coin struct {
	Name string `
	form:"name"
	json:"name" 
	example:"bitcoin"
	binding:"required"`
	Symbol string `
	form:"symbol"
	json:"symbol" 
	example:"btc"
	binding:"required"`
}

type Wallet struct {
	Currency            string `json:"currency"`
	Balance             string `json:"balance"`
	Locked              string `json:"locked"`
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
	UnitCurrency        string `json:"unit_currency"`
}
