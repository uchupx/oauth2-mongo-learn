package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-oauth2/mongo.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	manager := manage.NewDefaultManager()

	mongoStore := mongo.NewTokenStore(mongo.NewConfig(
		"mongodb://127.0.0.1:27017",
		"testdb",
	))
	// token memory store

	manager.MapTokenStorage(mongoStore)

	// manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("00000000"), jwt.SigningMethodHS512))

	// Parse and verify jwt access token
	// token, err := jwt.ParseWithClaims(access, &generates.JWTAccessClaims{}, func(t *jwt.Token) (interface{}, error) {
	// 	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("parse error")
	// 	}
	// 	return []byte("00000000"), nil
	// })
	// if err != nil {
	// 	// panic(err)
	// }

	// claims, ok := token.Claims.(*generates.JWTAccessClaims)
	// if !ok || !token.Valid {
	// 	// panic("invalid token")
	// }

	r := gin.Default()

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	r.GET("/authorize", func(c *gin.Context) {
		err := srv.HandleAuthorizeRequest(c.Writer, c.Request)
		if err != nil {
			// http.Error(w, err.Error(), http.StatusBadRequest)
			c.JSON(http.StatusBadRequest, err.Error())
		}

		// c.JSON(200, gin.H{
		// 	"message": "pong",
		// })
	})

	router := r.Group("")

	router.Use(func(c *gin.Context) {
		_, err := srv.ValidationBearerToken(c.Request)
		if err != nil {
			// c.JSON(http.StatusBadRequest, err.Error())
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
