package model

type Coin struct {
	Market        string `form:"market" json:"market" binding:"required"`
	KoreanName    string `form:"korean_name" json:"korean_name" binding:"required"`
	EnglishName   string `form:"english_name" json:"english_name" binding:"required"`
	MarketWarning string `form:"market_warning" json:"market_warning" binding:"required"`
}

type Wallet struct {
	Currency            string `json:"currency"`
	Balance             string `json:"balance"`
	Locked              string `json:"locked"`
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
	UnitCurrency        string `json:"unit_currency"`
}

type Resp_Quote struct {
	Status struct {
		Timestamp string `json:"timestamp"`
	} `json:"status"`
	Data []struct {
		Name   string `json:"name" form:"name" binding:"required"`
		Symbol string `json:"symbol" form:"symbol" binding:"required"`
		Quote  struct {
			Usd struct {
				Price float32 `json:"price"`
			} `json:"USD"`
			BTC struct {
				Price float32 `json:"price"`
			} `json:"BTC"`
		} `json:"quote"`
	} `json:"data"`
}
