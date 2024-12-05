package helper

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Struktur untuk menyimpan informasi koneksi database
type DBInfo struct {
	DBString string
	DBName   string
}

// Fungsi untuk menghubungkan ke MongoDB
func MongoConnect(info DBInfo) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(info.DBString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Memeriksa koneksi
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("MongoDB connection established.")
	return client, nil
}
