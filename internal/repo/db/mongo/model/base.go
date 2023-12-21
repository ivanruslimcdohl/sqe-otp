package mongomodel

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Model interface {
	ColName() string
	Indexes() []mongo.IndexModel
}

func GetIDFilter(idStr string, kOpt ...string) (bson.D, error) {
	k := "_id"
	if len(kOpt) > 0 {
		k = kOpt[0]
	}
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, err
	}
	return bson.D{{k, objID}}, nil
}

func GetUpdateData(updSt interface{}) (bson.D, error) {
	by, err := bson.Marshal(updSt)
	if err != nil {
		return nil, err
	}

	var updateData bson.D
	err = bson.Unmarshal(by, &updateData)
	if err != nil {
		return nil, err
	}

	return bson.D{{Key: "$set", Value: updateData}}, nil
}
