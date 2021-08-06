package model

type Favorite struct {
	Name string `form:"name" json:"name" example:"Bitcoin" binding:"required"`
}

type Coin struct {
	Market      string  `json:"market"`
	KoreanName  string  `json:"korean_name"`
	EnglishName string  `json:"english_name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	HighPrice   float32 `json:"high_price"`
	LowPrice    float32 `json:"low_price"`
	ChangeRate  float32 `json:"change_rate"`
	Change      string  `json:"change"`
}

type Wallet struct {
	Currency            string `json:"currency"`
	Balance             string `json:"balance"`
	Locked              string `json:"locked"`
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
	UnitCurrency        string `json:"unit_currency"`
}

//coin marketcap
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
