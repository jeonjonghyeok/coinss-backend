package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
// @Success 200 {object} rds.Resp_Quote
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /coin/list [get]
func (c *Controller) CoinList(ctx *gin.Context) {
	respQuote, err := rds.GetCoinlist()
	if err != nil {
		panic(err)
	}
	coinlist := respQuote.Data

	ctx.JSON(http.StatusOK, coinlist)

}

type testHeader struct {
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

	h := testHeader{}
	if err := ctx.ShouldBindHeader(&h); err != nil {
		ctx.JSON(200, err)
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
		log.Print(err)
		os.Exit(1)
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)

	var wallet []model.Wallet
	json.Unmarshal(respBody, &wallet)
	log.Println(wallet)

	ctx.JSON(http.StatusOK, wallet)

}
