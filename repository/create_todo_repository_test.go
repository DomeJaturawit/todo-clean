package repository

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"regexp"
	"testing"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/mockdata"
)

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

func TestCreateTodoRepository(t *testing.T) {
	suite.Run(t, new(TestCreateRepositoryTestSuite))
}

type TestCreateRepositoryTestSuite struct {
	suite.Suite
	sqlMock    sqlmock.Sqlmock
	sqlMockDB  *sql.DB
	mockGormDB *gorm.DB
	repository domain.TodoRepositoryInterface
}

func (suite *TestCreateRepositoryTestSuite) SetupSuite() {

}

func (suite *TestCreateRepositoryTestSuite) TearDownSuite() {

}

func (suite *TestCreateRepositoryTestSuite) SetupTest() {

	var err error

	suite.sqlMockDB, suite.sqlMock, err = sqlmock.New()
	assert.NoError(suite.T(), err)

	suite.mockGormDB, err = gorm.Open(postgres.New(postgres.Config{
		DriverName:           mockdata.DriverName,
		DSN:                  mockdata.MockDBDSN,
		PreferSimpleProtocol: true,
		Conn:                 suite.sqlMockDB,
	}), &gorm.Config{})
	assert.NoError(suite.T(), err)
	suite.repository = NewRepository(suite.mockGormDB)

}

func (suite *TestCreateRepositoryTestSuite) TearDownTest() {

	assert.NoError(suite.T(), suite.sqlMock.ExpectationsWereMet())
}

func (suite *TestCreateRepositoryTestSuite) Test_Happy() {

	mockEntity := mockdata.CreateTodoEntityMockData()
	suite.sqlMock.MatchExpectationsInOrder(true)

	suite.sqlMock.ExpectBegin()

	suite.sqlMock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "`+common.TodoTable+`"`)).WithArgs(
		mockEntity.ID,
		mockEntity.Title,
		mockEntity.Description,
		mockEntity.Status,
		mockEntity.CreatedAt,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := suite.repository.CreateTodoRepository(suite.mockGormDB.Begin(), mockEntity)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
}

func (suite *TestCreateRepositoryTestSuite) Test_Error_Something_Went_Wrong() {

	mockEntity := mockdata.CreateTodoEntityMockData()
	suite.sqlMock.MatchExpectationsInOrder(true)
	suite.sqlMock.ExpectBegin()

	suite.sqlMock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "`+common.TodoTable+`"`)).WithArgs(
		mockEntity.ID,
		mockEntity.Title,
		mockEntity.Description,
		mockEntity.Status,
		mockEntity.CreatedAt,
	).WillReturnError(mockdata.RepositoryError)

	result, err := suite.repository.CreateTodoRepository(suite.mockGormDB.Begin(), mockEntity)

	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), mockdata.RepositoryError.Error())
}
