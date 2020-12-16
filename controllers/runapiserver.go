package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"rest_api/database"
)

func InitialApiServer(db *database.DBHandler) error {
	r := mux.NewRouter()
	RunPersonApiOnRouter(r, db)
	return http.ListenAndServe(":8000", r)

}

func RunPersonApiOnRouter(r *mux.Router, db *database.DBHandler) {
	personHandler := NewPersonAPiHandler(db)
	person := r.PathPrefix("/v1/person").Subrouter()
	person.Methods("GET").Path("/all_persons").HandlerFunc(personHandler.GetAllPersons)
	person.Methods("GET").Path("/{name}").HandlerFunc(personHandler.GetPersonByName)
	person.Methods("GET").Path("/person_by_id/{id}").HandlerFunc(personHandler.GetPersonById)
	person.Methods("POST").Path("/create_person").HandlerFunc(personHandler.AddPerson)
	person.Methods("PUT").Path("/update_person").HandlerFunc(personHandler.UpdatePerson)
	person.Methods("DELETE").Path("/delete_person/{id}").HandlerFunc(personHandler.DeletePerson)
}
