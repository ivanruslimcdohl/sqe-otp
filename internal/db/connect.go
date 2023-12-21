package db

import (
	"context"
	"fmt"
	"log"

	"github.com/ivanruslimcdohl/sqe-otp/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbClient struct {
	*mongo.Client
}

func New(cfg config.DBCfg) *dbClient {
	log.Println("Connecting to MongoDB...")

	// TODO: get db pwd from secure secret manager
	connURI := fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.DBUser, "1234567890", cfg.DBHost, cfg.DBPort)
	clientOptions := options.Client().ApplyURI(connURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Panicf("failed to connect to mongo, err: %v\n", err)
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Panicf("failed to ping to mongo, err: %v\n", err)
	}

	log.Println("Connected to MongoDB!")

	return &dbClient{client}
}

func (dbC *dbClient) Close() {
	if err := dbC.Disconnect(context.Background()); err != nil {
		log.Panicf("failed to disconnect from mongo, err: %v\n", err)
	}
	log.Println("Connection to MongoDB closed.")
}
