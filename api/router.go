package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jeonjonghyeok/coinss-backend/controller"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5000
// @BasePath /api/v1
func Start() {

	r := gin.Default()
	c := controller.NewController()

	//duplicate logging
	//r.Use(gin.Logger())

	r.Use(cors.Default())

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		log.Println(recovered)
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", recovered))
	}))
	r.GET("/panic", func(c *gin.Context) {
		// panic with a string -- the custom middleware could save this to a database or report it to the user
		panic("test error")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	v1 := r.Group("api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("signup", c.AddUser)
			user.POST("signin", c.SigninUser)
		}
		coin := v1.Group("/coin")
		{
			coin.GET("list", c.Coins)
			coin.GET("wallet", c.Wallet)
			coin.POST("favorite", c.Favorite)
			coin.GET("favorites", c.Favorites)
			coin.POST("search", c.Search)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run("localhost:5000")
}
