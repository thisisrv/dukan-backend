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
	router.HandleFunc("/deleteProduct/{id}",controllers.DeleteOneProduct).Methods("DELETE")	//ok
	router.HandleFunc("/updateProduct/{id}&{field}&{value}",controllers.UpdateOneProduct).Methods("PUT")	//ok

	return router
}

