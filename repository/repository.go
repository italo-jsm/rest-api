package repository

import (
	"github.com/italosm/rest-api/model"
)
//ThingRepositoryInterface represents a things Repository interface
type ThingRepositoryInterface interface{
	CreateThing(t *model.Thing) (*model.Thing)
	FindAllThings()[]model.Thing
	FindOneThing(id string) *model.Thing
	DeleteOneThing(id string) bool
	UpdateOneThing(id string, Thing *model.Thing) (bool, *model.Thing)
}

//comentando algo

//StateRepositoryInterface represents a state Repository interface
type StateRepositoryInterface interface{
	CreateState(t *model.State) (*model.State)
	FindAllStates()[]model.State
	FindOneState(id string) *model.State
	DeleteOneState(id string) bool
	UpdateOneState(id string, State *model.State) (bool, *model.State)
}