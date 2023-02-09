package repository_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"todo-clean/domain"
	"todo-clean/repository"
	"todo-clean/util/mockdata"
)

func TestRollBackCreateTodoRepository(t *testing.T) {
	suite.Run(t, new(TestRollBackCreateRepositoryTestSuite))
}

type TestRollBackCreateRepositoryTestSuite struct {
	suite.Suite
	sqlMock    sqlmock.Sqlmock
	sqlMockDB  *sql.DB
	mockGormDB *gorm.DB
	repository domain.TodoRepositoryInterface
}

func (suite *TestRollBackCreateRepositoryTestSuite) SetupSuite() {

}

func (suite *TestRollBackCreateRepositoryTestSuite) TearDownSuite() {

}

func (suite *TestRollBackCreateRepositoryTestSuite) SetupTest() {
	var err error
	suite.sqlMockDB, suite.sqlMock, err = sqlmock.New()
	assert.NoError(suite.T(), err)

	// Init Gorm Driver Configuration
	suite.mockGormDB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  mockdata.MockDBDSN,
		DriverName:           mockdata.DriverName,
		Conn:                 suite.sqlMockDB,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	assert.NoError(suite.T(), err)
	suite.repository = repository.NewRepository(suite.mockGormDB)

}

func (suite *TestRollBackCreateRepositoryTestSuite) TearDownTest() {

}

// TODO Happy Test Case
func (suite *TestRollBackCreateRepositoryTestSuite) Test_Happy() {

}

// TODO: Fail Test Case
func (suite *TestRollBackCreateRepositoryTestSuite) Test_Error_Something_Went_Wrong() {

}
