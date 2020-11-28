package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Contact struct (Model)
type Contact struct {
	id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

// Init contacts var as slice Contact struct
var contacts []Contact

// Main Function
func main() {

	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	contacts = append(
		contacts, contact{
			id: 1, Name: "Chandler Bing", Phone: "305-917-1301", Email: "chandlerbing@office.com",
		},
	)
	contacts = append(
		contacts, contact{
			id: 2, Name: "Ross Geller", Phone: "210-684-8953", Email: "rossgeller@office.com",
		},
	)
	contacts = append(
		contacts, contact{
			id: 3, Name: "Rachel Green", Phone: "765-338-0312", Email: "rachelgreen@office.com",
		},
	)

	// Route handles & Endpoints
	r.HandleFunc("/contacts", getContact).Methods("GET")
	r.HandleFunc("/contacts/{id}", getContact).Methods("GET")
	r.HandleFunc("/contacts", createContact).Methods("POST")
	r.HandleFunc("/contacts/{id}", updateContact).Methods("PUT")
	r.HandleFunc("/contacts/{id}", delteContact).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":3000", r))
}
