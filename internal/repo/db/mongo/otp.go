package mongorepo

import (
	"context"

	mongomodel "github.com/ivanruslimcdohl/sqe-otp/internal/repo/db/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type otpRepo struct {
	*mongo.Collection
}

func newOTPRepo(d *mongo.Database) otpRepo {
	r := otpRepo{}
	r.Collection = d.Collection(r.getModel().ColName())

	return r
}

func (otpRepo) getModel() mongomodel.Model {
	return mongomodel.OTP{}
}

func (r otpRepo) Get(ctx context.Context, otpCode string) (mongomodel.OTP, error) {
	m := mongomodel.OTP{}
	filter := bson.M{"code": otpCode}
	err := r.FindOne(ctx, filter).Decode(&m)
	if err != nil {
		return m, err
	}

	return m, nil
}

func (r otpRepo) Insert(ctx context.Context, m mongomodel.OTP) (id string, err error) {
	res, err := r.InsertOne(ctx, m)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r otpRepo) Validate(ctx context.Context, otpCode string) error {
	update := bson.D{{"$set", bson.D{{"is_validated", true}}}}

	_, err := r.UpdateOne(ctx, bson.D{{"code", otpCode}}, update, nil)
	if err != nil {
		return err
	}

	return nil

}
