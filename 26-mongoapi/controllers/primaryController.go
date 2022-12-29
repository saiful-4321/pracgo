package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const connectionString = "mongodb+srv://root:bd.saif@4321#@cluster0.wiunfmj.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// Most Important
var collection *mongo.Collection

// Connect with mongo db
func init() {
	//  client option
	clientOption := options.client().ApplyURI(connectionString)

}
