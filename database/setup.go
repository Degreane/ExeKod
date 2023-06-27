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
	var err error
	mggURI := fmt.Sprintf("%s://%s:%s/%s", _dbSettings["Driver"], _dbSettings["Host"], _dbSettings["Port"], _dbSettings["DB"])
	log.Printf("Database->initDB I : %+v\n\n", mggURI)
	_connection, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mggURI))
	if err != nil {
		log.Fatalf("Unable to connect to database :%+v\n", err)
	}
	err = _connection.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Unable to Ping database :%+v\n", err)
	}
	log.Printf("Database->initDB I : %+v\n", _connection)
	_connection.Disconnect(context.TODO())
}
