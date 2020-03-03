package main // Entry point, main func, packaged.

// Modules
import (
	"encoding/json" // Data Type (json) Conversion/Parsing.
	"log"           // Error Logging.
	"math/rand"
	"net/http" // Basic REST API infrastructure.
	"strconv"  // Data Type (string) Conversion/Parsing.

	"github.com/gorilla/mux" // HTTP router and URI/Path matching, for web servers. REST API infrastructure.
)

// Contact struct type to act as template for each record, added to our sudo-database. Contains fields to capture data & perform CRUD operations against.
type Contact struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Email   string `json:"Email"`
	Contact int    `json:"Contact"`
}

// Contacts represents empty instance of our sudo-database, to store contact records.
var Contacts []Contact

/*
w.Write([]byte(fmt.Sprintf(`{"ID": %d, "Name": "%s", "Email": "%s", "Contact": %d }`, int, str, str, int))) // return data.
*/

// Handler function/method, that staisfies Handler interface. Called on "/home" GET Request, (handles this request path).
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") // Tell client json is sent back in response, so client understands it.
	http.ServeFile(w, r, "index.html")
}

// Handler function/method, that staisfies Handler interface. Called on "/new" POST Request, (handles this request path). JSON Data to be sent in request body.
func postNewContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Tell client json is sent back in response, so client understands it.

	r.ParseForm() // Parses the request body in a key value map.
	x := r.Form   // Store the parsed form.

	//decoder := json.NewDecoder(r.Body) // New decoder, that reads request, to decode JSON in the request body.
	contactNumber, err := strconv.Atoi(x["contact"][0]) // Parse query string as Int.
	if err != nil {                                     //Error: ID query string passed is not a number.
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "ID is not a number!"}`))
		return
	}

	var newContact Contact

	newContact.ID = rand.Intn(1000) // Assign random ID integer between 1 and 1000.
	newContact.Name = x["name"][0]
	newContact.Email = x["email"][0]
	newContact.Contact = contactNumber
	Contacts = append(Contacts, newContact) // Append new contact to current db of contacts.
	json.NewEncoder(w).Encode(Contacts)     // Display full contacts db.
	w.WriteHeader(http.StatusCreated)       // Respond with status 201, to indicate successful creation request.
}

// Handler function/method, that staisfies Handler interface. Called on "/find?id=<id>" GET Request, (handles this request path).
func getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Tell client json is sent back in response, so client understands it.

	query := r.URL.Query() // Get all query params sent with URL.
	ID := query.Get("id")  // Get the id query param.

	intID, err := strconv.Atoi(ID) // Parse query string as Int.
	if err != nil {                //Error: ID query string passed is not a number.
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "ID is not a number!"}`))
		return
	}

	// Iterate through queryContacts (all contact records), where index is assigned the full range of indexes in Contacts db.
	for _, queryContacts := range Contacts {
		if queryContacts.ID == intID { // Check id against queried id.
			json.NewEncoder(w).Encode(queryContacts) // Encode into json string, and write as part of response 'w'.
		}
	}
	w.WriteHeader(http.StatusOK) // Respond with status 200, to indicate successful request.
}

// Handler function/method, that staisfies Handler interface. Called on "/all" GET Request, (handles this request path).
func getAllContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Tell client json is sent back in response, so client understands it.
	json.NewEncoder(w).Encode(Contacts)                // Write all contacts in response.
	w.WriteHeader(http.StatusOK)                       // Respond with status 200, to indicate successful request.
}

// Handler function/method, that staisfies Handler interface. Called on "/update?id=<id>" PUT Request, (handles this request path). JSON Data to be sent in request body.
func putUpdateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Tell client json is sent back in response, so client understands it.

	query := r.URL.Query() // Get all query params sent with URL.
	ID := query.Get("id")  // Get the id query param.

	intID, err := strconv.Atoi(ID) // Parse query string as Int.
	if err != nil {                //Error: ID query string passed is not a number.
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "ID is not a number!"}`))
		return
	}

	decoder := json.NewDecoder(r.Body) // New decoder, that reads request, to decode JSON in the request body.

	var updateData Contact // Create new empty contact record.

	// Decodes request body to be stored in 'updateData', and error handling if cannot be decoded.
	err = decoder.Decode(&updateData)
	if err != nil {
		panic(err) // Stop current go-routine (break out), to log error.
	}

	for index, contact := range Contacts { // Iterate through all contact records.
		if contact.ID == intID { // Locate client specified id record.

			updateData.ID = contact.ID                                 // Assign current ID as updated contacts ID number.
			Contacts = append(Contacts[:index], Contacts[index+1:]...) // Delete specified contact.
			Contacts = append(Contacts, updateData)                    // Append new contacts list with updated contact, for full complete list.
			json.NewEncoder(w).Encode(Contacts)                        // Encode into json string, and write as part of response 'w' sent.
		}
	}
}

// Handler function/method, that staisfies Handler interface. Called on "/delete?id=<id>" DELETE Request, (handles this request path).
func deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Tell client json is sent back in response, so client understands it.

	query := r.URL.Query() // Get all query params sent with URL.
	ID := query.Get("id")  // Get the id query param.

	intID, err := strconv.Atoi(ID) // Parse query string as Int.
	if err != nil {                //Error: ID query string passed is not a number.
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "ID is not a number!"}`))
		return
	}

	for index, contact := range Contacts { // Iterate through all contact records.
		if contact.ID == intID { // Locate client specified id record.

			Contacts = append(Contacts[:index], Contacts[index+1:]...) // Delete specified contact.
			json.NewEncoder(w).Encode(Contacts)                        // Encode into json string, and write as part of response 'w' sent.
		}
	}
	w.WriteHeader(http.StatusOK) // Respond with status 200, to indicate successful request.
}

// Handler function/method, that staisfies Handler interface. Called on "/deleteAll" DELETE Request, (handles this request path).
func deleteAllContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Tell client json is sent back in response, so client understands it.
	Contacts = nil                                     // Delete all Contacts data.
	json.NewEncoder(w).Encode(Contacts)                // Encode into json string, and write as part of response 'w' sent.
	w.WriteHeader(http.StatusOK)                       // Respond with status 200, to indicate successful request.
}

// Handler function/method for POST Request (Create), that staisfies Handler interface.
func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Tell client json is sent back in response, so client understands it.
	w.WriteHeader(http.StatusCreated)                  // Respond with status 201, to indicate successful creation request.
	w.Write([]byte(`{"message": "http POST Request (Create)"}`))
}

// Handler function/method for GET Request (Read), that staisfies Handler interface.
func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Tell client json is sent back in response, so client understands it.
	w.WriteHeader(http.StatusOK)                       // Respond with status 200, to indicate successful request.
	w.Write([]byte(`{"message": "http GET Request (Read)"}`))
}

// Handler function/method for POST Request (Update), that staisfies Handler interface.
func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Tell client json is sent back in response, so client understands it.
	w.WriteHeader(http.StatusAccepted)                 // Respond with status 202, to indicate accpeted request, but not acted upon.
	w.Write([]byte(`{"message": "http PUT Request (Update)"}`))
}

// Handler function/method for DELETE Request (Delete), that staisfies Handler interface.
func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Tell client json is sent back in response, so client understands it.
	w.WriteHeader(http.StatusOK)                       // Respond with status 200, to indicate successful request.
	w.Write([]byte(`{"message": "http DELETE Request (Delete)"}`))
}

// Handler function/method for any exceptions/errors, that staisfies Handler interface.
func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")               // Tell client json is sent back in response, so client understands it.
	w.WriteHeader(http.StatusNotFound)                               // 404 error.
	w.Write([]byte(`{"message": "not found: Error 404 - Caught!"}`)) // Error handling/catching.
}

// Main entry point into API App, packaged. Provides all routing, and initialise server.
func main() {

	// Init API infrastructure.
	r := mux.NewRouter()                         // Init instance of request router/handler (interface?).
	api := r.PathPrefix("/api/v0-1").Subrouter() // New sub-router. "/api/v1", i.e.: "localhost:8080/api/v1/{endpoint}".

	// Handles given pattern/end-points/routes, to execute the given handler function, for each given request type received.
	api.HandleFunc("", get).Methods(http.MethodGet)       // localhost:8080/api/v0-1/
	api.HandleFunc("", post).Methods(http.MethodPost)     // localhost:8080/api/v0-1/
	api.HandleFunc("", put).Methods(http.MethodPut)       // localhost:8080/api/v0-1/
	api.HandleFunc("", delete).Methods(http.MethodDelete) // localhost:8080/api/v0-1/
	api.HandleFunc("", notFound)                          // localhost:8080/api/v0-1/

	api.HandleFunc("/contacts/home", homePage).Methods(http.MethodGet)                  // localhost:8080/api/v0-1/contacts/home
	api.HandleFunc("/contacts/new", postNewContact).Methods(http.MethodPost)            // localhost:8080/api/v0-1/contacts/new
	api.HandleFunc("/contacts/all", getAllContacts).Methods(http.MethodGet)             // localhost:8080/api/v0-1/contacts/all
	api.HandleFunc("/contacts/find", getContact).Methods(http.MethodGet)                // localhost:8080/api/v0-1/contacts/find?id=<id>
	api.HandleFunc("/contacts/update", putUpdateContact).Methods(http.MethodPut)        // localhost:8080/api/v0-1/contacts/update?id=<id>
	api.HandleFunc("/contacts/delete", deleteContact).Methods(http.MethodDelete)        // localhost:8080/api/v0-1/contacts/delete?id=<id>
	api.HandleFunc("/contacts/deleteAll", deleteAllContacts).Methods(http.MethodDelete) // localhost:8080/api/v0-1/contacts/deleteAll

	log.Fatal(http.ListenAndServe(":8080", r)) // Listens for incoming TCP connection requests. Set port, to open local server on and handle the incoming requests.
}
