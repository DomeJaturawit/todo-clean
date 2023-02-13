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

func TestBeginCreateTodoRepository(t *testing.T) {
	suite.Run(t, new(TestBeginCreateRepositoryTestSuite))
}

type TestBeginCreateRepositoryTestSuite struct {
	suite.Suite
	sqlMock    sqlmock.Sqlmock
	sqlMockDB  *sql.DB
	mockGormDB *gorm.DB
	repository domain.TodoRepository
}

func (suite *TestBeginCreateRepositoryTestSuite) SetupSuite() {

}

func (suite *TestBeginCreateRepositoryTestSuite) TearDownSuite() {

}

func (suite *TestBeginCreateRepositoryTestSuite) SetupTest() {
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

func (suite *TestBeginCreateRepositoryTestSuite) TearDownTest() {

}

func (suite *TestBeginCreateRepositoryTestSuite) Test_Happy() {
	suite.sqlMock.MatchExpectationsInOrder(true)
	suite.sqlMock.ExpectBegin()

	result, err := suite.repository.Begin()

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)

}

// TODO: Fail Test Case
func (suite *TestBeginCreateRepositoryTestSuite) Test_Error_Something_Went_Wrong() {

}
