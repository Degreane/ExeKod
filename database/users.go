// Here we are going to initialize Users collection in the database
package database

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type accountType uint

const (
	guest accountType = iota << 1
	public
	client
	admin
	webadmin
	sysadmin
)

type user struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Credentials struct {
		UserName string `bson:"username" json:"username"`
		Password string `bson:"password" json:"password"`
	} `bson:"credentials" json:"credentials"`
	Name struct {
		First string `bson:"first" json:"first"`
		Last  string `bson:"last" json:"last"`
	} `bson:"name" json:"name"`
	Comments    []comment   `bson:"comments" json:"comments"`
	AccountType accountType `bson:"type,omitempty" json:"type,omitempty"`
	CreatedAt   time.Time   `bson:"createdat,omitempty" json:"createdat,omitempty"`
	UpdatedAt   time.Time   `bson:"updatedat,omitempty" json:"updatedat,omitempty"`
}

type comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	CreatedAt time.Time          `bson:"createdat,omitempty" json:"createdat,omitempty"`
	UpdatedAt time.Time          `bson:"updatedat,omitempty" json:"updatedat,omitempty"`
	Content   string             `bson:"content" json:"content"`
}

func getUser(query interface{}) (string, error) {
	var _theUser user
	res := _usersCollection.FindOne(context.TODO(), query)
	err := res.Decode(&_theUser)
	if err != nil {
		log.Printf("Database->users->getUser I : % +v\n\n", err)
		return "", err
	}
	log.Printf("Database->users->getUser II : % +v\n\n", _theUser)
	theJson, err := json.Marshal(_theUser)
	if err != nil {
		log.Printf("Database->users->getUser III : % +v\n\n", err)
		return "", err
	}
	log.Printf("Database->users->getUser IV : % +v\n\n", theJson)

	return string(theJson), nil
}

func countUsers(query interface{}) (int64, error) {
	count, err := _usersCollection.CountDocuments(context.TODO(), query)
	if err != nil {
		log.Printf("Database->Users->countUsers I : % +v\n\n", err)
		return 0, err
	}
	return count, nil
}
func insertUser(query interface{}) (*mongo.InsertOneResult, error) {
	return _usersCollection.InsertOne(context.TODO(), query)
}

func checkAndInseertAdminUser() {
	var err error
	adminAccount := struct {
		AccountType accountType `bson:"type,omitempty" json:"type,omitempty"`
	}{AccountType: sysadmin}
	var adminAccountBytes []byte
	adminAccountBytes, err = json.Marshal(adminAccount)
	if err != nil {
		log.Printf("Database->users->checkAndInsertAdminUser <I> : % +v\n\n", err)
	} else {
		var adminQuery interface{}
		err = bson.UnmarshalExtJSON(adminAccountBytes, true, &adminQuery)
		if err != nil {
			log.Printf("Database->users->checkAndInsertAdminUser <II> : % +v\n\n", err)
		} else {
			count, err := countUsers(adminQuery)
			if err != nil {
				log.Printf("Database->users->checkAndInsertAdminUser <III> : % +v\n\n", err)
			} else {
				if count == 0 {
					// if count == 0 means we should insert an admin into the collection and thus we should populate insertion as follows:
					_theUser := user{
						ID: primitive.NewObjectID(),
						Credentials: struct {
							UserName string "bson:\"username\" json:\"username\""
							Password string "bson:\"password\" json:\"password\""
						}{UserName: "fbanna", Password: "Shta2Telik"},
						Name: struct {
							First string "bson:\"first\" json:\"first\""
							Last  string "bson:\"last\" json:\"last\""
						}{First: "Faysal", Last: "AL-Banna"},
						AccountType: sysadmin,
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
						Comments:    []comment{{ID: primitive.NewObjectID(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Content: "Account Created By The System"}},
					}
					_theUserByte, err := json.Marshal(_theUser)
					if err != nil {
						log.Printf("Database->users->checkAndInsertAdminUser <IV> : % +v\n\n", err)
					} else {
						var _theUserInterface interface{}
						err = bson.UnmarshalExtJSON(_theUserByte, true, &_theUserInterface)
						if err != nil {
							log.Printf("Database->users->checkAndInsertAdminUser <V> : % +v\n\n", err)
						} else {
							res, err := insertUser(_theUserInterface)
							if err != nil {
								log.Printf("Database->users->checkAndInsertAdminUser <VI> : % +v\n\n", err)
							} else {
								log.Printf("Database->users->checkAndInsertAdminUser <VII> : % +v\n\n", res.InsertedID)
							}
						}
					}
				} else {
					theUser, err := getUser(adminQuery)
					if err != nil {
						log.Printf("Database->users->checkAndInsertAdminUser <VIII> : % +v\n\n", err)
					} else {
						log.Printf("Database->users->checkAndInsertAdminUser <IX> : % +v\n\n", theUser)
					}
				}
			}
		}

	}

}
