package routes

import (
	"BNI-TAP/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RouteMulai() {
	router := mux.NewRouter()
	router.HandleFunc("/inquiry", controllers.Inquiry).Methods("POST")
	fmt.Println("Connected to port 8003")
	log.Print(http.ListenAndServe(":8003", router))
}
