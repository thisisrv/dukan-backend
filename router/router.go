package router

import (
	"dukan/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	//handle routes to Products db
	router.HandleFunc("/products",controllers.GetAllProducts).Methods("GET")  //ok
	router.HandleFunc("/addProduct",controllers.CreateOneProduct).Methods("POST")		//ok
	router.HandleFunc("/deleteProduct/{id}",controllers.DeleteOneProduct).Methods("POST")	//ok
	router.HandleFunc("/updateProduct/{id}",controllers.UpdateOneProduct).Methods("POST")	//ok

	//handle routes to Sales
	router.HandleFunc("/sales/{date}",controllers.GetAllSales).Methods("GET") 
	router.HandleFunc("/addSale",controllers.CreateOneSale).Methods("POST")	

	return router
}

