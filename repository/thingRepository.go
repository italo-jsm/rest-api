package repository

import (
	"github.com/italosm/rest-api/db"
	"github.com/italosm/rest-api/model"
	"fmt"
)

//ThingRepository represents a repository of things
type ThingRepository struct {
}

//CreateThing saves a new thing in the database
func (repository *ThingRepository) CreateThing(t *model.Thing) *model.Thing {
	database := db.ConnectDatabase()
	insert, err := database.Prepare("insert into thing (name, address) values (?, ?)")
	if err != nil {
		panic(err.Error)
	}
	result, err := insert.Exec(t.Name, t.Address)
	if err != nil {
		return nil
	}
	lastID, err := result.LastInsertId()
	t.ID = lastID
	defer database.Close()
	return t
}

// FindAllThings returns all the things
func (repository *ThingRepository) FindAllThings() []model.Thing {
	db := db.ConnectDatabase()
	allThings, err := db.Query("select * from thing")
	if err != nil {
		panic(err.Error)
	}
	t := model.Thing{}
	things := []model.Thing{}

	for allThings.Next() {
		var ID int64
		var name, address string
		err = allThings.Scan(&ID, &name, &address)
		if err != nil {
			panic(err.Error)
		}
		t.ID = ID
		t.Name = name
		t.Address = address
		things = append(things, t)
	}
	defer db.Close()
	return things
}

//FindOneThing finds one thing
func (repository *ThingRepository) FindOneThing(id string) *model.Thing {
	database := db.ConnectDatabase()
	sqlStatement := "select * from thing where id=?"
	row := database.QueryRow(sqlStatement, id)
	var thing model.Thing
	err := row.Scan(&thing.ID, &thing.Name, &thing.Address)
	if err != nil {
		return nil
	}
	defer database.Close()
	return &thing
}

//DeleteOneThing deletes a thing from the database
func (repository *ThingRepository) DeleteOneThing(id string) bool {
	database := db.ConnectDatabase()
	thingToDelete := repository.FindOneThing(id)
	if thingToDelete != nil {
		sqlStatement := "delete from thing where id=?"
		_, err := database.Exec(sqlStatement, id)
		if err != nil {
			return false
		}
		return true
	}
	defer database.Close()
	return false
}

//UpdateOneThing updates a thing in the database
func (repository *ThingRepository) UpdateOneThing(id string, t *model.Thing) (bool, *model.Thing) {
	database := db.ConnectDatabase()
	thingToUpdate := repository.FindOneThing(id)
	if thingToUpdate != nil {
		sqlStatement := "update thing set Name=?, Address=? where id=?"
		_, err := database.Exec(sqlStatement, t.Name, t.Address, id)
		if err != nil {
			fmt.Println(err)
			return false, nil
		}
		return true, t
	}
	defer database.Close()
	return false, nil
}
