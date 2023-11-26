package dberrors

type ConflicError struct{}

func (e *ConflicError) Error() string {
	return "attempt to create a record with a pre-existing key"
}
