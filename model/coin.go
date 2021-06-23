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
