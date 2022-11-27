package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Tests happy path of submitting a well-formed GET /customers request
func TestGetCustomersHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/customers", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getCustomers)
	handler.ServeHTTP(rr, req)

	// Checks for 200 status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("getCustomers returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Checks for JSON response
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("Content-Type does not match: got %v want %v",
			ctype, "application/json")
	}

	// Seed data is array of length 3
	var customers []map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&customers)
	if err != nil {
		t.Errorf("Not an array of customers: %v", rr.Body.String())
	} else if len(customers) != 3 {
		t.Errorf("Seed data not of length 3")
	}

	// Serializes to camel-cased
	if _, ok := customers[0]["id"]; !ok {
		t.Errorf("No 'id' key")
	}
}

// Tests happy path of submitting a well-formed POST /customers request
func TestAddCustomerHandler(t *testing.T) {
	requestBody := strings.NewReader(`
		{
			"name": "Example Name",
			"role": "Example Role",
			"email": "Example Email",
			"phone": 5550199,
			"contacted": true
		}
	`)

	req, err := http.NewRequest("POST", "/customers", requestBody)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addCustomer)
	handler.ServeHTTP(rr, req)

	// Checks for 201 status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("addCustomer returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Checks for JSON response
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("Content-Type does not match: got %v want %v",
			ctype, "application/json")
	}
}

// Tests unhappy path of deleting a user that doesn't exist
func TestDeleteCustomerHandler(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/customers/e7847fee-3a0e-455e-b151-519bdb9851c7", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteCustomer)
	handler.ServeHTTP(rr, req)

	// Checks for 404 status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("deleteCustomer returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

// Tests unhappy path of getting a user that doesn't exist
func TestGetCustomerHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/customers/e7847fee-3a0e-455e-b151-519bdb9851c7", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getCustomer)
	handler.ServeHTTP(rr, req)

	// Checks for 404 status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("getCustomer returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
