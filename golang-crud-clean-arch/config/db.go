package config

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func MongoConnect() *mongo.Client {
	if mongoClient != nil {
		return mongoClient // Gunakan koneksi yang sudah ada
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI("mongodb+srv://lepi1:rizkidh123@cluster0.kfcww.mongodb.net/?appName=Cluster0").
		SetServerAPIOptions(serverAPI).
		SetMaxPoolSize(20).                 // Maksimum 20 koneksi dalam pool
		SetMinPoolSize(5).                  // Minimal 5 koneksi dalam pool
		SetConnectTimeout(10 * time.Second) // Timeout saat koneksi dibuat

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	mongoClient = client // Simpan koneksi ke variabel global agar bisa digunakan kembali
	return client
}

func MongoCollection(coll string) *mongo.Collection {
	client := MongoConnect() // Gunakan koneksi yang sudah ada
	return client.Database("gotrial").Collection(coll)
}
