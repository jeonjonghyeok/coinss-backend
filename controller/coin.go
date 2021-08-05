package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeonjonghyeok/coinss-backend/model"
	"github.com/jeonjonghyeok/coinss-backend/psql"
	"github.com/jeonjonghyeok/coinss-backend/rds"
	upbit "github.com/jeonjonghyeok/coinss-backend/token"
	"github.com/jeonjonghyeok/coinss-backend/utils"
)

// Coin-List godoc
// @Summary Coin-List
// @Description get coinlist
// @Tags coin
// @Accept  json
// @Produce  json
// @Param favorite body favorite true "Faavorite"
// @Success 200 {object} model.Coin
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/coin/list [get]
func (c *Controller) Coins(ctx *gin.Context) {
	var f favorite
	utils.HandleErr(ctx.BindJSON(&f))

	coins, err := rds.GetCoins(f.Name)
	utils.HandleErr(err)

	ctx.JSON(http.StatusOK, coins)
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
// @Router /api/v1/coin/wallet [get]
func (c *Controller) Wallet(ctx *gin.Context) {
	const URL = "https://api.upbit.com/v1/accounts"

	h := header{}
	utils.HandleErr(ctx.ShouldBindHeader(&h))
	id, err := upbit.Parse(h.Token)
	utils.HandleErr(err)

	user, err := psql.FindUserById(id)
	utils.HandleErr(err)
	token, err := upbit.NewUpbit(id, user.Accesskey, user.Secretkey)
	utils.HandleErr(err)
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	utils.HandleErr(err)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	utils.HandleErr(err)
	log.Println(resp.Status)
	if resp.Status == "401 Unauthorized" {
		panic("등록되지 않은 IP주소 입니다.")
	}
	respBody, _ := ioutil.ReadAll(resp.Body)

	var wallet []model.Wallet
	utils.HandleErr(json.Unmarshal(respBody, &wallet))
	ctx.JSON(http.StatusOK, wallet)

}

type favorite struct {
	Name string `form:"name" json:"name" example:"Bitcoin" binding:"required"`
}

// Favorite godoc
// @Summary Register Favority Coin
// @Description 관심코인 등록
// @ID post-coin-favorite
// @Tags coin
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param favorite body favorite true "Faavorite"
// @Success 200 {object} model.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/coin/favorite [post]
func (c *Controller) Favorite(ctx *gin.Context) {
	h := header{}

	utils.HandleErr(ctx.ShouldBindHeader(&h))

	id, err := upbit.Parse(h.Token)
	utils.HandleErr(err)

	var f favorite
	utils.HandleErr(ctx.BindJSON(&f))
	names := psql.GetFavorites(id)
	if names == "" {
		names = f.Name
	} else {
		names += "," + f.Name
	}
	utils.HandleErr(psql.Favorite(id, names))
	ctx.String(http.StatusOK, "Success")
}

// Favorite godoc
// @Summary Register Favority Coin
// @Description 관심코인 조회
// @ID post-coin-favorites
// @Tags coin
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} model.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/coin/favorites [get]
func (c *Controller) Favorites(ctx *gin.Context) {
	h := header{}
	utils.HandleErr(ctx.ShouldBindHeader(&h))
	id, err := upbit.Parse(h.Token)
	utils.HandleErr(err)

	names := psql.GetFavorites(id)
	if names == "" {
		ctx.JSON(http.StatusOK, nil)
		return
	}
	coins, err := rds.GetCoins(names)
	utils.HandleErr(err)

	ctx.JSON(http.StatusOK, coins)
}

type search struct {
	Search string `json:"search"`
}

// Search godoc
// @Summary Register Favority Coin
// @Description 검색어 저장
// @ID post-coin-search
// @Tags coin
// @Accept  json
// @Produce  json
// @Param search body search true "Search"
// @Success 200 {object} search
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/coin/search [post]
func (c *Controller) Search(ctx *gin.Context) {
	var s search
	utils.HandleErr(ctx.BindJSON(&s))

	psql.SetSearch(s.Search)

	ctx.String(http.StatusOK, "Success")
}
