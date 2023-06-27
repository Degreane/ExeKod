// package database should connect to the database
// and provide a handler to work with mongo db
// the database is provided with the yaml file
// and thus it should be read from the database there

package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDB() {
	// define Error
	var err error
	// Set the mongoURI header
	mggURI := fmt.Sprintf("%s://%s:%s/%s", _dbSettings["Driver"], _dbSettings["Host"], _dbSettings["Port"], _dbSettings["DB"])
	log.Printf("Database->initDB I : %+v\n\n", mggURI)
	// Set the connection to the database
	_connection, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mggURI))
	if err != nil {
		log.Fatalf("Database->initDB II :%+v\n\n", err)
	}
	// ping database to check if its alive
	err = _connection.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Database->initDB III :%+v\n\n", err)
	}
	// Define Collections
	_usersCollection = _connection.Database(_dbSettings["DB"].(string)).Collection("users")
	log.Printf("Database->initDB IV : %+v\n\n", _connection)
	checkAndInseertAdminUser()
	_connection.Disconnect(context.TODO())
}
