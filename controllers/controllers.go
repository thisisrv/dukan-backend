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

const connectionString =  "mongodb+srv://mathswithrv:gCYWloxyVOY1PsqO@cluster0.umulikh.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
// const connectionString = "mongodb://localhost:27017/"
//IMP
var collection *mongo.Collection
var mongoclient *mongo.Client 
// var salesCollection *mongo.Collection

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
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
	w.Header().Set("Access-Control-Allow-Methods", "GET") // Allow only GET requests

	//set content type to json
	w.Header().Set("Content-Type", "application/json")
	allproducts := getAllProducts()
	json.NewEncoder(w).Encode(allproducts)
}

func CreateOneProduct(w http.ResponseWriter, r *http.Request){

    // Set CORS headers
    w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
	w.Header().Set("Access-Control-Allow-Methods", "POST") // Allow only POST requests

    // Set content type to JSON
    w.Header().Set("Content-Type", "application/json")

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
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "POST")

	params := mux.Vars(r)

	var productData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&productData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateOneProduct(params["id"], productData)

	fmt.Println(productData)
	fmt.Println("Fileds updated")
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneProduct(w http.ResponseWriter, r *http.Request){
	// Set CORS headers
    w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
	w.Header().Set("Access-Control-Allow-Methods", "POST") // Allow only POST requests

	 // Set content type to JSON
	 w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	deleteOneProduct(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}


// ############################### SALE #######################################

func GetAllSales(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	allSales := getAllSales(params["date"])
	json.NewEncoder(w).Encode(allSales)
}

func CreateOneSale(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "POST")

	var sale models.Sales

	err := json.NewDecoder(r.Body).Decode(&sale)

	if err != nil {
		log.Fatal(err)
		return
	}

	insertSale(sale)

	// Encode the product with the updated ObjectId
    encodedProduct, err := json.Marshal(sale)
    if err != nil {
        log.Fatal(err)
        return
    }

	// Write the encoded product to the response
    w.Write(encodedProduct)
}
