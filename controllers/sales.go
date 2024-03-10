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

func getAllSales(date string) []interface{} {
    cursor, err := mongoclient.Database("Sales").Collection(date).Find(context.Background(), bson.D{{}})
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(context.Background())

    var sales []interface{}

    for cursor.Next(context.Background()) {
        var sale map[string]interface{}
        if err := cursor.Decode(&sale); err != nil {
            log.Fatal(err)
        }

        // Check if the sale has an "items" field
        if items, ok := sale["items"].(interface{}); ok {
            for _, item := range items.(primitive.A) {
                // Perform type assertion to access the underlying map
                if itemMap, ok := item.(interface{}); ok {
                    // Append the item map to the sales slice
                    sales = append(sales, itemMap)
                }
				
            }
			// sales = append(sales, items)
        }
    }

    if err := cursor.Err(); err != nil {
        log.Fatal(err)
    }
    return sales
}


func insertSale(sale models.Sales){

	//Generate a new Object id
	sale.ID = primitive.NewObjectID()
	
	result, err := mongoclient.Database("Sales").Collection(sale.Date).InsertOne(context.Background(), sale)

	if err != nil {
		log.Fatal(err)
	}

	// Iterate over items and call updateInventory function
    for _, item := range sale.Items {
        key := item["key"].(string)
        quantity := item["quantity"].(float64)

        // Call updateInventory function
        updateInventory(key, quantity)
    }

	fmt.Println("rows added", result.InsertedID)
}

func updateInventory(key string, quantity float64){
	id, err := primitive.ObjectIDFromHex(key)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	// Define the update to decrement the quantity
    update := bson.M{"$inc": bson.M{"quantity": -quantity}}

	// Perform the update operation
    _, err = collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        log.Fatal(err)
    }

}