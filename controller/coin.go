package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeonjonghyeok/coinss-backend/rds"
	upbit "github.com/jeonjonghyeok/coinss-backend/token"
)

// Coin-List godoc
// @Summary Coin-List
// @Description get coinlist
// @Tags coin
// @Accept  json
// @Produce  json
// @Param t query string false "token"
// @Success 200 {object} rds.Resp_Quote
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /coin/list [get]
func (c *Controller) CoinList(ctx *gin.Context) {
	t := ctx.Request.URL.Query().Get("token")
	upbit.Parse(t)

	respQuote, err := rds.GetCoinlist()
	if err != nil {
		panic(err)
	}
	m := respQuote.Data

	ctx.JSON(http.StatusOK, m)

}
