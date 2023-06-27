// package database should connect to the database
// and provide a handler to work with mongo db
// the database is provided with the yaml file
// and thus it should be read from the database there

package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	_connection      *mongo.Client
	_dbSettings      map[string]interface{}
	_usersCollection *mongo.Collection
)

func Initialize(dbSettings map[string]interface{}) {
	_dbSettings = dbSettings
	//log.Printf("Database->Initialize I : % +v\n", dbSettings)
	initDB()
}
