package model

import (
	"github.com/gofrs/uuid"
)

// Customer is a model of the "customers" table.
type Customer struct {
	ID   uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Name string    `bun:",notnull"`
}
