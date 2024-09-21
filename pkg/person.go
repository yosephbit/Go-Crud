package person

import (
    "errors"
    "sync"

    "github.com/google/uuid"
)

// Person struct represents a person with their details
type Person struct {
    ID      string   `json:"id"`
    Name    string   `json:"name"`
    Age     int      `json:"age"`
    Hobbies []string `json:"hobbies"`
}

// In-memory database to store persons
var (
    persons    = make(map[string]Person)
    personsMtx sync.RWMutex
)

// GetAllPersons returns all persons
func GetAllPersons() []Person {
    personsMtx.RLock()
    defer personsMtx.RUnlock()

    result := make([]Person, 0, len(persons))
    for _, p := range persons {
        result = append(result, p)
    }

    return result
}

// GetPerson returns the person with the specified ID
func GetPerson(id string) (Person, error) {
    personsMtx.RLock()
    defer personsMtx.RUnlock()

    p, ok := persons[id]
    if !ok {
        return Person{}, errors.New("person not found")
    }

    return p, nil
}

// CreatePerson creates a new person
func CreatePerson(p Person) (Person, error) {
    personsMtx.Lock()
    defer personsMtx.Unlock()

    // Generate a new UUID for the person
    p.ID = uuid.New().String()

    // Save the person to the database
    persons[p.ID] = p

    return p, nil
}

// UpdatePerson updates an existing person
func UpdatePerson(id string, updatedPerson Person) error {
    personsMtx.Lock()
    defer personsMtx.Unlock()

    _, ok := persons[id]
    if !ok {
        return errors.New("person not found")
    }

    // Update the person in the database
    persons[id] = updatedPerson

    return nil
}

// DeletePerson deletes an existing person
func DeletePerson(id string) error {
    personsMtx.Lock()
    defer personsMtx.Unlock()

    _, ok := persons[id]
    if !ok {
        return errors.New("person not found")
    }

    // Delete the person from the database
    delete(persons, id)

    return nil
}
