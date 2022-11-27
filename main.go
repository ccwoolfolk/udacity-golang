package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var customers = []Customer{
	{
		ID:        1,
		Name:      "Alfred",
		Role:      "Butler",
		Email:     "alfie@hotmail.com",
		Phone:     5555555555,
		Contacted: false,
	},
	{
		ID:        2,
		Name:      "Ronald",
		Role:      "President",
		Email:     "ronald@hotmail.com",
		Phone:     5555555556,
		Contacted: true,
	},
	{
		ID:        3,
		Name:      "Mat",
		Role:      "Gambler",
		Email:     "mat@hotmail.com",
		Phone:     5555555557,
		Contacted: false,
	},
}

func getCustomerById(id int) (Customer, error) {
	for _, c := range customers {
		if c.ID == id {
			return c, nil
		}
	}

	return Customer{}, errors.New("No customer found")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idRaw := mux.Vars(r)["id"]
	id, parsingErr := strconv.Atoi(idRaw)

	customer, notFoundErr := getCustomerById(id)

	if parsingErr != nil || notFoundErr != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(customer)
		w.WriteHeader(http.StatusOK)

	}
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(customers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
