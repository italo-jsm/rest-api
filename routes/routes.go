package routes

import (
	"time"
	"fmt"
	"net/http"
	"github.com/italosm/rest-api/repository"
	"github.com/italosm/rest-api/controller"
	"github.com/gorilla/mux"
)

func middlewareHandler(handler http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("before handler; middleware start")
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("middleware finished; %s", time.Since(start))
	})
}
//CreateRoutes create the routes
func CreateRoutes() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	thingControler := new(controller.ThingController)
	thingControler.Repo = &repository.ThingRepository{}
	stateController := new(controller.StateController)
	stateController.StateRepo = &repository.StateRepository{}
	stateController.ThingRepo = &repository.ThingRepository{}
	router.HandleFunc("/thing", thingControler.CreateThing).Methods("POST")
	router.Handle("/things", middlewareHandler(http.HandlerFunc(thingControler.GetAllThings))).Methods("GET")
	router.HandleFunc("/things/{id}", thingControler.GetOneThing).Methods("GET")
	router.HandleFunc("/things/{id}", thingControler.DeleteThing).Methods("DELETE")
	router.HandleFunc("/things/{id}", thingControler.UpdateThing).Methods("PUT")

	router.HandleFunc("/state", stateController.CreateState).Methods("POST")
	return router
}