// Helper func

package controllers

import (
	"context"
	"dukan/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017/" //"mongodb+srv://mathswithrv:gCYWloxyVOY1PsqO@cluster0.umulikh.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "Products"

//IMP
var collection *mongo.Collection
var mongoclient *mongo.Client 

//connect with mongodb
func init(){

	//clientOption
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	mongoclient = client
	collection = mongoclient.Database("Product").Collection("inventory")
	fmt.Println("Connected to DB")

}

func GetAllProducts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	allproducts := getAllProducts()
	json.NewEncoder(w).Encode(allproducts)
}

func CreateOneProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "POST")

	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		log.Fatal(err)
		return
	}

	insertOneProduct(product)

	// Encode the product with the updated ObjectId
    encodedProduct, err := json.Marshal(product)
    if err != nil {
        log.Fatal(err)
        return
    }

	// Write the encoded product to the response
    w.Write(encodedProduct)
}

func UpdateOneProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "POST")

	params := mux.Vars(r)

	updateOneProduct(params["id"], params["field"], params["value"])

	// fmt.Println(params)
	fmt.Println("Fileds updated")
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "POST")

	params := mux.Vars(r)
	deleteOneProduct(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}



