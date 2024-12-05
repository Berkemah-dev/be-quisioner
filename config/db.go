package config

import (
	"log"

	"github.com/Berkemah-dev/be-quisioner/helper"
)

var MongoString string = GetEnv("MONGOSTRING")

var mongoinfo = helper.DBInfo{
	DBString: MongoString,
	DBName:   "berkemahquisioner",
}

var Mongoconn, ErrorMongoconn = helper.MongoConnect(mongoinfo)

func init() {
	if ErrorMongoconn != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", ErrorMongoconn)
	} else {
		log.Println("Connected to MongoDB successfully!")
	}
}
