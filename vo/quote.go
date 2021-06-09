package vo

type Resp_Quote struct {
	Status struct {
		Timestamp string `json:"timestamp"`
	} `json:"status"`
	Data []struct {
		Name   string   `json:"name"`
		Symbol string   `json:"symbol"`
		Tags   []string `json:"tags"`
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
