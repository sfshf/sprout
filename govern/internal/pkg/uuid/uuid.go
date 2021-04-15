package uuid

import (
	"github.com/google/uuid"
)

// Temporary:

type UUID = uuid.UUID

func NewUUID() (UUID, error) {
	return uuid.NewRandom()
}
