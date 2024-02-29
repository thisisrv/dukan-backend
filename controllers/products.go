package controllers

import (
	"context"
	"dukan/models"
	"fmt"
	"log"
	"strconv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ********************************* MONGO DB Functions for Products ********************************

func getAllProducts() []primitive.M {


	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var lists []primitive.M

	for cursor.Next(context.TODO()) {
		var product bson.M
		err := cursor.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}

		lists = append(lists, product)
	}

	defer cursor.Close(context.Background())

	return lists	
}

func insertOneProduct(product models.Product){

	//Generate a new Object id
	product.ID = primitive.NewObjectID()
	
	//PUSH to DB
	result, err := collection.InsertOne(context.Background(), product)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("rows added", result.InsertedID)
}

func updateOneProduct(productId string, field string, value string){
	id, err := primitive.ObjectIDFromHex(productId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	var update bson.M
	if field == "name"{
		update = bson.M{"$set": bson.M{field: value}}
	}else {
		integerValue, _ := strconv.Atoi(value)
		update = bson.M{"$set": bson.M{field: integerValue}}
	}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("rows affected", result.ModifiedCount)
}

func deleteOneProduct(productId string){

	id, err := primitive.ObjectIDFromHex(productId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.DeletedCount)
}