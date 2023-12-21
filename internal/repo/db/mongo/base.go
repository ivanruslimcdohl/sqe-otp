package mongorepo

import (
	"context"
	"log"

	repodb "github.com/ivanruslimcdohl/sqe-otp/internal/repo/db"
	mongomodel "github.com/ivanruslimcdohl/sqe-otp/internal/repo/db/mongo/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(d *mongo.Database) repodb.DB {
	dbR := repodb.DB{}
	dbR.OTP = newOTPRepo(d)

	return dbR
}

func RegisterIndexes(db *mongo.Database) {
	type collRepo interface {
		getModel() mongomodel.Model
	}

	collRepos := []collRepo{
		otpRepo{},
	}

	for _, r := range collRepos {
		m := r.getModel()
		for _, v := range m.Indexes() {
			_, err := db.Collection(m.ColName()).Indexes().CreateOne(context.Background(), v)
			if err != nil {
				log.Panicf("failed to register index mongo, coll: %s, err: %v", m.ColName(), err)
			}
		}
	}
}
