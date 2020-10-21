package repository

import (
	"github.com/italosm/rest-api/db"
	"github.com/italosm/rest-api/model"
)

//StateRepository represents a stateRepository and implements Repository
type StateRepository struct {

}

//CreateState saves a new state in the database
func (repository *StateRepository) CreateState(s *model.State) *model.State {
	database := db.ConnectDatabase()
	insert, err := database.Prepare("insert into state (portNumber, status, thingId) values (?, ?, ?)")
	if err != nil {
		panic(err.Error)
	}
	result, err := insert.Exec(s.PortNumber, s.Status, s.Thing.ID)
	if err != nil {
		return nil
	}
	lastID, err := result.LastInsertId()
	s.ID = lastID
	defer database.Close()
	return s
}

// FindAllStates returns all the states
func (repository *StateRepository) FindAllStates() []model.State {
	db := db.ConnectDatabase()
	allStates, err := db.Query("select * from state")
	if err != nil {
		panic(err.Error)
	}
	s := model.State{}
	states := []model.State{}

	for allStates.Next() {
		var ID, portNumber int64
		var on bool
		err = allStates.Scan(&ID, &portNumber, &on)
		if err != nil {
			panic(err.Error)
		}
		s.ID = ID
		s.PortNumber = portNumber
		s.Status = on
		states = append(states, s)
	}
	defer db.Close()
	return states
}

//FindOneState finds one state
func (repository *StateRepository) FindOneState(id string) *model.State {
	database := db.ConnectDatabase()
	sqlStatement := "select * from state where id=?"
	row := database.QueryRow(sqlStatement, id)
	var state model.State
	err := row.Scan(&state.ID, &state.PortNumber, &state.Thing.ID)
	if err != nil {
		return nil
	}
	defer database.Close()
	return &state
}

//DeleteOneState deletes a state from the database
func (repository *StateRepository) DeleteOneState(id string) bool {
	database := db.ConnectDatabase()
	stateToDelete := repository.FindOneState(id)
	if stateToDelete != nil {
		sqlStatement := "delete from state where id=?"
		_, err := database.Exec(sqlStatement, id)
		if err != nil {
			return false
		}
		return true
	}
	defer database.Close()
	return false
}

//UpdateOneState updates a state in the database
func (repository *StateRepository) UpdateOneState(id string, s *model.State) (bool, *model.State) {
	database := db.ConnectDatabase()
	stateToDelete := repository.FindOneState(id)
	if stateToDelete != nil {
		sqlStatement := "update state set PortNumber=?, On=? where id=?"
		_, err := database.Exec(sqlStatement, s.PortNumber, s.Status, id)
		if err != nil {
			return false, nil
		}
		return true, s
	}
	defer database.Close()
	return false, nil
}
