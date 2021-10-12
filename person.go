package ddd_go

import "github.com/google/uuid"

// Person is a model entity
type Person struct {
	ID   uuid.UUID
	Name string
	Age  int
}
