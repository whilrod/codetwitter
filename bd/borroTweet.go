package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BorroTweet(ID string, UserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("codetwitter")
	col := db.Collection("twitt")

	objID, _ := primitive.ObjectIDFromHex(ID)
	condicion := bson.M{
		"_id":    objID,
		"userid": UserId,
	}
	_, err := col.DeleteOne(ctx, condicion)
	return err

}
