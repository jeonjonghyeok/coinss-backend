package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jeonjonghyeok/coinss-backend/model"
	"github.com/jeonjonghyeok/coinss-backend/psql"
	"github.com/jeonjonghyeok/coinss-backend/rds"
	upbit "github.com/jeonjonghyeok/coinss-backend/token"
)

// Coin-List godoc
// @Summary Coin-List
// @Description get coinlist
// @Tags coin
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Coin
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /coin/list [get]
func (c *Controller) Coins(ctx *gin.Context) {
	coin, err := rds.GetCoins()
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, coin)
}

type header struct {
	Token string `header:"Token"`
}

// Coin-Wallet godoc
// @Summary Coin-Wallet
// @Description get coinwallet
// @Tags coin
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} model.Wallet
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /coin/wallet [get]
func (c *Controller) Wallet(ctx *gin.Context) {
	const URL = "https://api.upbit.com/v1/accounts"
	var secret_key string
	var access_key string

	h := header{}
	if err := ctx.ShouldBindHeader(&h); err != nil {
		panic(err)
	}
	id, err := upbit.Parse(h.Token)
	if err != nil {
		panic(err)
	}

	log.Println("id = ", id)
	access_key, secret_key, err = psql.FindUserKey(id)
	log.Println("access_key = ", access_key)
	log.Println("secret_key = ", secret_key)

	token, err := upbit.NewUpbit(id, access_key, secret_key)
	log.Println("upbit token = ", token)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	log.Println(resp.Status)
	if resp.Status == "401 Unauthorized" {
		panic("등록되지 않은 IP주소 입니다.")
	}
	respBody, _ := ioutil.ReadAll(resp.Body)

	var wallet []model.Wallet
	if err := json.Unmarshal(respBody, &wallet); err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, wallet)

}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Coin-Quote godoc
// @Summary websocket
// @Description get coinquote
// @Tags coin
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Resp_Quote
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /coin/quote [patch]
func (c *Controller) Quote(ctx *gin.Context) {
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		panic(err)
	}
	for {
		var quote model.Resp_Quote
		quote, err = rds.GetCoinlist()
		if err != nil {
			panic(err)
		}
		ws.WriteJSON(quote)
		time.Sleep(10 * time.Second)
	}
}
