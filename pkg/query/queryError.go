package query

import (
	"encoding/json"
)

//Error describe a database query error
type Error struct {
	ID   string
	Data ErrorString
}

//ErrorString describe the error string
type ErrorString struct {
	Error string
}

//NewError create a new query error
func NewError(id string, data string) Error {
	return Error{
		ID: id,
		Data: ErrorString{
			Error: data,
		},
	}
}

//ToString convert an Error to a string
func (e Error) ToString() string {
	s, _ := json.Marshal(e)
	return string(s)
}
