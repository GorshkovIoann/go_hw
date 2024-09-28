package idgenerators

import (
	"math/rand"

	"github.com/google/uuid"
)

func GenerateRandomID() int {
	return rand.Intn(1000000)
}
func GenerateUUID() int {
	uuidObj := uuid.New()
	return int(uuidObj.ID())
}
