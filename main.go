package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Contact struct (Model)
type Contact struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

// Init contacts var as slice Contact struct
var contacts []Contact

// Get all contacts
func getContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

// Get a single contact
func getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Looping through contacts and find one with the id from the params
	for _, item := range contacts {
		if item.Id == params["Id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Contact{})
}

// Create a contact
func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(contact)
}

// Update contact
func updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range contacts {
		if item.Id == params["Id"] {
			contacts = append(contacts[:idx], contacts[idx+1:]...)
			var contact Contact
			_ = json.NewDecoder(r.Body).Decode(&contact)
			contact.Id = params["Id"]
			contacts = append(contacts, contact)
			json.NewEncoder(w).Encode(contact)
			return
		}
	}
}

// Delete contact
func deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range contacts {
		if item.Id == params["Id"] {
			contacts = append(contacts[:idx], contacts[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(contacts)
}

// Main Function
func main() {

	// Init router
	router := mux.NewRouter()

	// Hardcoded data - @todo: add database
	contacts = append(
		contacts, Contact{
			Id: "1", Name: "Chandler Bing", Phone: "305-917-1301", Email: "chandlerbing@office.com",
		},
	)
	contacts = append(
		contacts, Contact{
			Id: "2", Name: "Ross Geller", Phone: "210-684-8953", Email: "rossgeller@office.com",
		},
	)
	contacts = append(
		contacts, Contact{
			Id: "3", Name: "Rachel Green", Phone: "765-338-0312", Email: "rachelgreen@office.com",
		},
	)

	// Route handles & Endpoints
	router.HandleFunc("/contacts", getContacts).Methods("GET")
	router.HandleFunc("/contacts/{Id}", getContact).Methods("GET")
	router.HandleFunc("/contacts", createContact).Methods("POST")
	router.HandleFunc("/contacts/{Id}", updateContact).Methods("PUT")
	router.HandleFunc("/contacts/{Id}", deleteContact).Methods("DELETE")

	// Start server
	fmt.Println("listening at port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
