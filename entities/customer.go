package entities

import "gopkg.in/mgo.v2/bson"

type Customer struct {
	ID     		bson.ObjectId `json:"id" bson:"_id"`
	Name		string `bson:"name" json:"name"`
	Age			string `bson:"age" json:"age"`
}

type Customers []Customer
