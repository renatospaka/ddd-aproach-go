package tavern

import "github.com/google/uuid"

// Person is a entity that represents a person in all Domains
type Person struct {
	ID uuid.UUID
	Name string
	Age int
}
