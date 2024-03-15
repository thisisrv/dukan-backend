package router

import (
	"dukan/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func healthCheck(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"Health Check Fine")
}

func Router() *mux.Router{
	router := mux.NewRouter()

	//handle routes to Products db
	router.HandleFunc("/products",controllers.GetAllProducts).Methods("GET")  //ok
	router.HandleFunc("/addProduct",controllers.CreateOneProduct).Methods("POST")		//ok
	router.HandleFunc("/deleteProduct/{id}",controllers.DeleteOneProduct).Methods("POST")	//ok
	router.HandleFunc("/updateProduct/{id}",controllers.UpdateOneProduct).Methods("POST")	//ok
	router.HandleFunc("/health-check", healthCheck).Methods("GET")

	//handle routes to Sales
	router.HandleFunc("/sales/{date}",controllers.GetAllSales).Methods("GET") 
	router.HandleFunc("/addSale",controllers.CreateOneSale).Methods("POST")	

	return router
}

