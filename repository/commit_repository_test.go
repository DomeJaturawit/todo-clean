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

func TestCommitCreateTodoRepository(t *testing.T) {
	suite.Run(t, new(TestCommitCreateRepositoryTestSuite))
}

type TestCommitCreateRepositoryTestSuite struct {
	suite.Suite
	sqlMock    sqlmock.Sqlmock
	sqlMockDB  *sql.DB
	mockGormDB *gorm.DB
	repository domain.TodoRepository
}

func (suite *TestCommitCreateRepositoryTestSuite) SetupSuite() {

}

func (suite *TestCommitCreateRepositoryTestSuite) TearDownSuite() {

}

func (suite *TestCommitCreateRepositoryTestSuite) SetupTest() {
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

func (suite *TestCommitCreateRepositoryTestSuite) TearDownTest() {

}

func (suite *TestCommitCreateRepositoryTestSuite) Test_Happy() {
	suite.sqlMock.MatchExpectationsInOrder(true)
	suite.sqlMock.ExpectBegin()
	suite.sqlMock.ExpectCommit()

	tx := suite.mockGormDB.Begin()

	err := suite.repository.Commit(tx)
	assert.Nil(suite.T(), err)

}

// TODO: Fail Test Case
func (suite *TestCommitCreateRepositoryTestSuite) Test_Error_Something_Went_Wrong() {

}
