package model

//State represents a state
type State struct {
	Thing Thing      `json:"thing"`
	ID int64         `json:"ID"`
	PortNumber int64 `json:"portNumber"`
	Status bool     `json:"status"`
}
