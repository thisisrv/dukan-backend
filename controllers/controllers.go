// Helper func

package controllers

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)

const connectionString = "mongodb+srv://mathswithrv:gCYWloxyVOY1PsqO@cluster0.umulikh.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "Shop DB"
const collectionName = "Product"

//IMP
var collection *mongo.collection

//connect with mongodb
func init(){

	//clientOption
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
}