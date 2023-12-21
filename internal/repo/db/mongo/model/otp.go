package mongomodel

import (
	"time"

	"github.com/ivanruslimcdohl/sqe-otp/internal/kit/timekit"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OTP struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Code        string             `bson:"code" json:"code"`
	UserID      string             `bson:"user_id" json:"user_id"`
	IsValidated bool               `bson:"is_validated" json:"is_validated"`
	ExpiresAt   time.Time          `bson:"expires_at" json:"expires_at"`
}

func (OTP) ColName() string {
	return "otp"
}

func (OTP) Indexes() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys:    bson.M{"code": 1},
			Options: options.Index().SetUnique(true),
		},
	}
}

func (m OTP) CreatedAt() time.Time {
	return timekit.ToWIB(m.ID.Timestamp())
}
