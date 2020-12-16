package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest_api/database"
	"strconv"
)

type PersonAPiHandler struct {
	DBHandler *database.PersonHandler
}

func NewPersonAPiHandler(db *database.DBHandler) *PersonAPiHandler {
	handler, err := database.NewPersonHandler(db)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &PersonAPiHandler{DBHandler: handler}
}

func (p *PersonAPiHandler) GetAllPersons(res http.ResponseWriter, req *http.Request) {
	people, err := p.DBHandler.GeTAllPeople()
	if err != nil {
		res.Write([]byte("somethings wrong"))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(people)
}

func (p *PersonAPiHandler) GetPersonById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("somethings wrong"))
		return
	}
	person, err := p.DBHandler.GetPersonById(id)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Not Found"))
		return
	}
	json.NewEncoder(res).Encode(person)
	res.WriteHeader(http.StatusOK)
}

func (p *PersonAPiHandler) GetPersonByName(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]
	person, err := p.DBHandler.GetPersonByName(name)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Not found"))
		return
	}
	json.NewEncoder(res).Encode(person)
}

func (p *PersonAPiHandler) AddPerson(res http.ResponseWriter, req *http.Request) {
	var person database.Person
	json.NewDecoder(req.Body).Decode(&person)
	if err := p.DBHandler.AddPerson(&person); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("invalid data type"))
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("ok"))

}

func (p *PersonAPiHandler) UpdatePerson(res http.ResponseWriter, req *http.Request) {
	var person *database.Person
	if err := json.NewDecoder(req.Body).Decode(&person); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("invalid data"))
		return
	}
	err := p.DBHandler.UpdatePerson(person)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("somethings wrong"))
		return
	}
	json.NewEncoder(res).Encode(person)
	res.WriteHeader(http.StatusAccepted)
}

func (p *PersonAPiHandler) DeletePerson(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)["id"]
	id, err := strconv.Atoi(vars)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("somethings wrong"))
		return
	}
	person, err := p.DBHandler.GetPersonById(id)
	if err != nil {
		res.Write([]byte("Not Found"))
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := p.DBHandler.DeletePerson(person); err != nil {
		res.Write([]byte("somethings wrong"))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Write([]byte("deleted successfully"))
	res.WriteHeader(http.StatusOK)
}
