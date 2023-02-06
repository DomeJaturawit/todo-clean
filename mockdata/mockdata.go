package mockdata

import (
	"errors"
	"github.com/google/uuid"
	"github.com/goombaio/namegenerator"
	"math/rand"
	"time"
	"todo-clean/domain"
)

const (
	MockDBDSN  = "mockDB"
	DriverName = "postgres"
)

var (
	RepositoryError = errors.New("RepositoryError")
)

func genSeed() (seed int64) {
	seed = int64(rand.Intn(99))
	return seed
}
func NewTitle() string {
	titleGenerator := namegenerator.NewNameGenerator(genSeed())
	title := titleGenerator.Generate()
	return title
}

func NewDescription() string {
	descriptionGenerator := namegenerator.NewNameGenerator(genSeed())
	description := descriptionGenerator.Generate()
	return description
}

func NewStatus() string {
	statusGenerator := namegenerator.NewNameGenerator(genSeed())
	status := statusGenerator.Generate()
	return status
}

func CreateTodoEntityMockData() domain.CreateTodoEntity {
	return domain.CreateTodoEntity{
		ID:          uuid.New(),
		Title:       NewTitle(),
		Description: NewDescription(),
		Status:      NewStatus(),
		CreatedAt:   time.Now(),
	}
}
