package oauth2

import (
	"gopkg.in/go-oauth2/mongo.v3"

	"github.com/uchupx/oauth2-mongo-learn/config"
)

type mongoConn struct {
	tokenConn  mongo.TokenStore
	clientConn mongo.ClientStore
}

func Connection(conf *config.Config) mongoConn {
	var conn mongoConn

	conn.clientConn = clientConn(conf)
	conn.tokenConn = tokenConn(conf)

	return conn
}

func clientConn(conf *config.Config) mongo.ClientStore {
	mongoClientStore := mongo.NewClientStore(mongo.NewConfig(
		conf.Database.MongoDB.URL,
		conf.Database.MongoDB.Name,
	))

	return *mongoClientStore

}

func tokenConn(conf *config.Config) mongo.TokenStore {
	mongoTokenStore := mongo.NewTokenStore(mongo.NewConfig(
		conf.Database.MongoDB.URL,
		conf.Database.MongoDB.Name,
	))

	return *mongoTokenStore
}
