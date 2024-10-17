package mongodb

import (
	"base_service/config"
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect(dbConfig config.MongoDBConfig) *MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mngdbSource := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
	)
	Client, err := mongo.Connect(ctx, options.Client().ApplyURI(mngdbSource))
	if err != nil {
		return nil
	}
	db := Client.Database(dbConfig.Database)

	return &MongoDB{
		Client:   Client,
		Database: db,
	}
}

func ConnectDB() {
	conf := config.GetConfig()

	MngDb = connect(conf.MDB)
	if MngDb == nil {
		log.Println("Connection is Nill")
		log.Fatal(nil)
	}
	slog.Info("Connected to  Mongo database")
}

func CloseDB() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := MngDb.Client.Disconnect(ctx); err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info("Disconnected from Mongo Database")
}
