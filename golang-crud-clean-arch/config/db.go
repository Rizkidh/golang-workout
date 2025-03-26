package config

import (
	"context"
	"fmt"
	"os"
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

	mongoURI := os.Getenv("MONGO_URI")
	mongoDBName := os.Getenv("MONGO_DB_NAME")

	if mongoURI == "" || mongoDBName == "" {
		panic("MongoDB configuration is missing in environment variables")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(mongoURI).
		SetServerAPIOptions(serverAPI).
		SetMaxPoolSize(20).                 // Maksimum 20 koneksi dalam pool
		SetMinPoolSize(5).                  // Minimal 5 koneksi dalam pool
		SetMaxConnecting(10).               // Limit concurrent new connections
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
	mongoDBName := os.Getenv("MONGO_DB_NAME")
	return client.Database(mongoDBName).Collection(coll)
}
