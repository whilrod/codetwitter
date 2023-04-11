package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://whilrod:Sarita.21@cluster0.pj2uxe6.mongodb.net/?retryWrites=true&w=majority")

func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa")
	return client
}

func ChequeoConection() int {
	//err := mongo.Connect(context.TODO(), clientOptions)
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}
