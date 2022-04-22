package generators

import (
	"errors"

	"github.com/google/uuid"
)

type IDGenerator interface {
	NewString() string
}

var UUIDGenError = errors.New("UUID generator error")

type UUIDGenenerator struct{}

func NewUUIDGenenerator() *UUIDGenenerator {
	return &UUIDGenenerator{}
}

func (gen *UUIDGenenerator) NewString() string {
	return uuid.New().String()
}
