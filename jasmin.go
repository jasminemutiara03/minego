package jasmine

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertRencanastudi(db string, rencanastudi RencanaStudi) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("rencanastudi").InsertOne(context.TODO(), rencanastudi)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertNilai(db string, nilai Nilai) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("nilai").InsertOne(context.TODO(), nilai)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertMatakuliah(db string, matakuliah MataKuliah) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("matakuliah").InsertOne(context.TODO(), matakuliah)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertTranskrip(db string, transkrip Transkrip) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("transkrip").InsertOne(context.TODO(), transkrip)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertUsers(db string, users Users) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("users").InsertOne(context.TODO(), users)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
