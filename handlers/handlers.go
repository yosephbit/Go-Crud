package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    person "main/pkg"
)

// NewRouter creates a new router with defined routes
func NewRouter() *mux.Router {
    router := mux.NewRouter()

    // Routes
    router.HandleFunc("/person", getAllPersons).Methods("GET")
    router.HandleFunc("/person/{id}", getPerson).Methods("GET")
    router.HandleFunc("/person", createPerson).Methods("POST")
    router.HandleFunc("/person/{id}", updatePerson).Methods("PUT")
    router.HandleFunc("/person/{id}", deletePerson).Methods("DELETE")

    return router
}

// getAllPersons returns all persons
func getAllPersons(w http.ResponseWriter, r *http.Request) {
    persons := person.GetAllPersons()

    // Encode response as JSON
    json.NewEncoder(w).Encode(persons)
}

// getPerson returns the person with the specified ID
func getPerson(w http.ResponseWriter, r *http.Request) {
    // Get the person ID from the request URL
    params := mux.Vars(r)
    personID := params["id"]

    person, err := person.GetPerson(personID)
    if err != nil {
        http.NotFound(w, r)
        return
    }

    // Encode response as JSON
    json.NewEncoder(w).Encode(person)
}

// createPerson creates a new person
func createPerson(w http.ResponseWriter, r *http.Request) {
    // Decode request body into Person struct
    var p person.Person
    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Create the person
    newPerson, err := person.CreatePerson(p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Encode response as JSON
    json.NewEncoder(w).Encode(newPerson)
}

// updatePerson updates an existing person
func updatePerson(w http.ResponseWriter, r *http.Request) {
    // Get the person ID from the request URL
    params := mux.Vars(r)
    personID := params["id"]

    // Decode request body into Person struct
    var updatedPerson person.Person
    err := json.NewDecoder(r.Body).Decode(&updatedPerson)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Update the person in the database
    err = person.UpdatePerson(personID, updatedPerson)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Encode response as JSON
    json.NewEncoder(w).Encode(updatedPerson)
}

// deletePerson deletes an existing person
func deletePerson(w http.ResponseWriter, r *http.Request) {
    // Get the person ID from the request URL
    params := mux.Vars(r)
    personID := params["id"]

    // Delete the person
    err := person.DeletePerson(personID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Send success response
    w.WriteHeader(http.StatusNoContent)
}
