package controller

import (
	"strconv"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/italosm/rest-api/model"
	"github.com/italosm/rest-api/repository"
)

//StateController represents a state controller
type StateController struct {
	StateRepo repository.StateRepositoryInterface
	ThingRepo repository.ThingRepositoryInterface
}
//CreateState handles a POST request
func (stateController *StateController) CreateState(w http.ResponseWriter, r *http.Request) {
	var newState model.State
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error)
	}
	json.Unmarshal(reqBody, &newState)
	thingID := newState.Thing.ID
	thing := stateController.ThingRepo.FindOneThing(strconv.FormatInt(thingID, 10))
	if (thing == nil){
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newState.Thing = *thing
	createdState := stateController.StateRepo.CreateState(&newState)
	if(createdState != nil){
		w.WriteHeader(http.StatusCreated)
	}else{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(createdState)
}