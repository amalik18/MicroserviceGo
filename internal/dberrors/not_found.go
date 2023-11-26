package dberrors

import "fmt"

type NotFoundError struct {
	Entity string
	Id     string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Unable to find %s with Id %s", e.Entity, e.Id)
}
