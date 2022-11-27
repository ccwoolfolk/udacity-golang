package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getCustomer(w http.ResponseWriter, r *http.Request) {

}

func getCustomers(w http.ResponseWriter, r *http.Request) {

}

func addCustomer(w http.ResponseWriter, r *http.Request) {

}

func updateCustomer(w http.ResponseWriter, r *http.Request) {

}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {

}

func main() {
	port := 3000
	router := mux.NewRouter()

	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id:[0-9]+}", getCustomer).Methods("GET")
	router.HandleFunc("/customers/{id:[0-9]+}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id:[0-9]+}", deleteCustomer).Methods("DELETE")

	fmt.Println(fmt.Sprintf("Serving on port %d", port))
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
