package repository_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/repository"
	mockdata "todo-clean/util/mockdata"
)

func TestDeleteTodoRepository(t *testing.T) {
	suite.Run(t, new(TestDeleteRepositoryTestSuite))
}

type TestDeleteRepositoryTestSuite struct {
	suite.Suite
	sqlMock       sqlmock.Sqlmock
	sqlMockDB     *sql.DB
	mockGormDB    *gorm.DB
	repository    domain.TodoRepository
	mockDataModel []domain.GetTodoEntity
}

func (suite *TestDeleteRepositoryTestSuite) SetupSuite() {

}

func (suite *TestDeleteRepositoryTestSuite) TearDownSuite() {

}

func (suite *TestDeleteRepositoryTestSuite) SetupTest() {
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
	suite.repository = repository.NewRepository(suite.mockGormDB)

}

func (suite *TestDeleteRepositoryTestSuite) TearDownTest() {

}

func (suite *TestDeleteRepositoryTestSuite) Test_Happy() {
	mockQuery := mockdata.QueryDeleteTodoEntityMockData()
	suite.sqlMock.ExpectBegin()
	suite.sqlMock.ExpectExec(regexp.QuoteMeta(`UPDATE "`+common.TodoTable+`"`+` SET `+`"`+common.TodoDeletedAtCol+`"`+`=$1`+
		` WHERE `+common.TodoIDCol+` = $2 AND "`+common.TodoTable+`"."`+common.TodoDeletedAtCol+`" = $3`)).
		WithArgs(
			time.Now().Unix(),
			mockQuery.ID,
			0,
		).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := suite.repository.DeleteTodoRepository(context.Background(), suite.mockGormDB.Begin(), mockQuery)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
}

func (suite *TestDeleteRepositoryTestSuite) Test_Error_Something_Went_Wrong() {
	mockQuery := mockdata.QueryDeleteTodoEntityMockData()
	suite.sqlMock.ExpectBegin()
	suite.sqlMock.ExpectExec(regexp.QuoteMeta(`UPDATE "`+common.TodoTable+`"`+` SET `+`"`+common.TodoDeletedAtCol+`"`+`=$1`+
		` WHERE `+common.TodoIDCol+` = $2 AND "`+common.TodoTable+`"."`+common.TodoDeletedAtCol+`" = $3`)).
		WithArgs(
			time.Now().Unix(),
			mockQuery.ID,
			0,
		).WillReturnError(common.ErrDBDeleteTodo)

	result, err := suite.repository.DeleteTodoRepository(context.Background(), suite.mockGormDB.Begin(), mockQuery)
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), common.ErrDBDeleteTodo.Error())

}
