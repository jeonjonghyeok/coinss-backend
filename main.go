package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeonjonghyeok/coinss-backend/controller"
	_ "github.com/jeonjonghyeok/coinss-backend/docs"
	"github.com/jeonjonghyeok/coinss-backend/psql"
	"github.com/jeonjonghyeok/coinss-backend/rds"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	DB_USER     = "jjh"
	DB_PASSWORD = "jjh"
	DB_NAME     = "jjh"
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
func main() {
	if err := psql.Connect(fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)); err != nil {
		log.Fatal(err)
	}
	if err := rds.Connect(); err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	c := controller.NewController()
	r.Use(gin.Logger())
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))
	r.GET("/panic", func(c *gin.Context) {
		// panic with a string -- the custom middleware could save this to a database or report it to the user
		panic("foo")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ohai")
	})

	v1 := r.Group("api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("signup", c.AddUser)
			user.POST("signin", c.SigninUser)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.Use(handlePanic)
	r.Run(":5000")
}
