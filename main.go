package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/oauth2-mongo-learn/config"
	"github.com/uchupx/oauth2-mongo-learn/services/oauth2"
)

func main() {
	conf := config.Loader()

	r := gin.Default()
	srv := oauth2.CreateServe(conf)

	r.GET("/authorize", func(c *gin.Context) {
		err := srv.HandleAuthorizeRequest(c.Writer, c.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}
	})

	router := r.Group("")

	router.Use(func(c *gin.Context) {
		_, err := srv.ValidationBearerToken(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/token", func(c *gin.Context) {
		srv.HandleTokenRequest(c.Writer, c.Request)
	})

	r.Run(":8000")

}
