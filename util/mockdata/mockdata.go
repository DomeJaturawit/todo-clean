package mockdata

import (
	"github.com/goombaio/namegenerator"
	"math/rand"
)

const (
	MockDBDSN  = "mockDB"
	DriverName = "postgres"
)

func GenSeed() (seed int64) {
	seed = int64(rand.Intn(99))
	return seed
}

func NewString() string {
	statusGenerator := namegenerator.NewNameGenerator(GenSeed())
	status := statusGenerator.Generate()
	return status
}
