package controllers

import (
	"context"
	"dukan/models"
	"fmt"
	"log"
	"reflect"

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

func updateOneProduct(productId string, product map[string]interface{}){
	id, err := primitive.ObjectIDFromHex(productId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	// Get the existing product from the database
	var existingProduct models.Product
	err = collection.FindOne(context.Background(), filter).Decode(&existingProduct)

	if err != nil {
		log.Fatal(err)
	}

	update := bson.M{"$set": bson.M{}}

	// // Iterate over the fields of the product map
	// for key, value := range product {
	// 	// Check if the key exists in the existing product
	// 	if _, ok := existingProduct[key]; ok {
	// 		update["$set"].(bson.M)[key] = value
	// 	}
	// }

	// Iterate over the fields of the newProduct object
	v := reflect.ValueOf(existingProduct)
	
	for i := 1; i < v.NumField(); i++ {
		field := v.Type().Field(i).Tag.Get("bson")
		// newValue := v.Field(i).Interface()

		// Check if field is present in product
		value, ok := product[field]
		if ok {
			// Add the field and its value to the update
			update["$set"].(bson.M)[field] = value
		}
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