package oauth2

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/uchupx/oauth2-mongo-learn/config"

	// "gopkg.in/oauth2.v3"

	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
)

func CreateServe(conf *config.Config) *server.Server {
	mongo := Connection(conf)

	manager := manage.NewDefaultManager()
	manager.MapTokenStorage(&mongo.tokenConn)
	manager.MapClientStorage(&mongo.clientConn)

	// generates.NewJWTAccessGenerate()

	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("00000000"), jwt.SigningMethodHS512))

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

	return srv
}
