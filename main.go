package main

import (
	"dukan/router"
	"fmt"
	"log"
	"net/http"
)

func main()  {
	fmt.Println("MongoDB API")
	fmt.Println("Server is starting up ...")
	r := router.Router()
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", r))
}
