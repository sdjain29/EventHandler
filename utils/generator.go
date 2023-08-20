package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/pborman/uuid"
)

func SetSeedForRand() {
	rand.NewSource(time.Now().UTC().UnixNano())
}

func UuidGenerator() string {
	return strings.Replace(uuid.NewRandom().String(), "-", "", -1)
}
