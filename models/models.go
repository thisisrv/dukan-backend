// DB SCHEMA
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct{
	ID 				primitive.ObjectID	`json:"_id" bson:"_id"`
	Name 			string				`json:"name" bson:"name"`
	Quantity 		int 				`json:"quantity" bson:"quantity"`
	Cost_Price 		int 				`json:"cost_price" bson:"cost_price"`
	Selling_Price	int 				`json:"selling_price" bson:"selling_price"`
}


type Sales struct{
	ID 				primitive.ObjectID		`json:"_id" bson:"_id"`
	Date 			string					`json:"date" bson:"date"`
	Items		    []map[string]interface{}	`json:"items" bson:"items"`
}

//make datewise collection for sales