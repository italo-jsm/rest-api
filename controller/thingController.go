package controller

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
	"github.com/italosm/rest-api/model"
	"github.com/italosm/rest-api/repository"
	"net/http"
)

//ThingController represents a controller of things
type ThingController struct {
	Repo repository.ThingRepositoryInterface 
}

//CreateThing handles a POST request
func (thingController *ThingController)CreateThing(w http.ResponseWriter, r *http.Request) {
	var newThing model.Thing
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error)
	}
	json.Unmarshal(reqBody, &newThing)
	createdThing := thingController.Repo.CreateThing(&newThing)
	if(createdThing != nil){
		w.WriteHeader(http.StatusCreated)
	}else{
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(createdThing)
}

//GetAllThings returns all things in the base
func (thingController *ThingController) GetAllThings(w http.ResponseWriter, r *http.Request) {
	things := thingController.Repo.FindAllThings()
	json.NewEncoder(w).Encode(things)
}

//GetOneThing returns a element by a given id
func (thingController *ThingController) GetOneThing(w http.ResponseWriter, r *http.Request) {
	thingID := mux.Vars(r)["id"]
	thing := thingController.Repo.FindOneThing(thingID)
	if (thing == nil){
		w.WriteHeader(http.StatusNoContent)
	}
	json.NewEncoder(w).Encode(thing)
}

//UpdateThing updates a thing
func (thingController *ThingController) UpdateThing(w http.ResponseWriter, r *http.Request) {
	thingID := mux.Vars(r)["id"]
	var thingToUpdate model.Thing

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(reqBody, &thingToUpdate)
	success, updatedThing := thingController.Repo.UpdateOneThing(thingID, &thingToUpdate)
	if !success {
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(updatedThing)
}

//DeleteThing deletes a thing by a given id
func (thingController *ThingController) DeleteThing(w http.ResponseWriter, r *http.Request) {
	thingID := mux.Vars(r)["id"]
	success := thingController.Repo.DeleteOneThing(thingID)
	if !success {
		w.WriteHeader(http.StatusNotFound)
	}
}

