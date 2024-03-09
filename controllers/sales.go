package controllers

import (
	"context"
	"dukan/models"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ********************************* MONGO DB Functions for Sales ********************************

func getAllSales() []primitive.M {

	cursor, err := salesCollection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var sales []primitive.M

	for cursor.Next(context.TODO()) {
		var sale bson.M
		err := cursor.Decode(&sale)
		if err != nil {
			log.Fatal(err)
		}

		sales = append(sales, sale)
	}

	defer cursor.Close(context.Background())

	return sales	
}

func insertSale(sale models.Sales){

	//Generate a new Object id
	sale.ID = primitive.NewObjectID()
	
	result, err := mongoclient.Database("Sales").Collection(sale.Date).InsertOne(context.Background(), sale)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("rows added", result.InsertedID)
}

