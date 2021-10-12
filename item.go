package ddd_go

import "github.com/google/uuid"

// Item is a model entity
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}
