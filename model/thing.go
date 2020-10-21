package model


//Thing represents a thing
type Thing struct{
	ID int64		`json:"ID"`
	Name string 	`json:"name"`
	Address string	`json:"address"`
}