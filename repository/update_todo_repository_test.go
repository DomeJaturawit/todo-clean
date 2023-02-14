package repository_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"regexp"
	"testing"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/repository"
	"todo-clean/util"
	"todo-clean/util/mockdata"
)

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

func TestUpdateTodoRepository(t *testing.T) {
	suite.Run(t, new(TestUpdateRepositoryTestSuite))
}

type TestUpdateRepositoryTestSuite struct {
	suite.Suite
	sqlMock    sqlmock.Sqlmock
	sqlMockDB  *sql.DB
	mockGormDB *gorm.DB
	repository domain.TodoRepository
}

func (suite *TestUpdateRepositoryTestSuite) SetupSuite() {

}

func (suite *TestUpdateRepositoryTestSuite) TearDownSuite() {

}

func (suite *TestUpdateRepositoryTestSuite) SetupTest() {

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

func (suite *TestUpdateRepositoryTestSuite) TearDownTest() {

}

func (suite *TestUpdateRepositoryTestSuite) Test_Happy() {

	mockEntity := mockdata.UpdateTodoEntityMockData()
	mockQuery := mockdata.QueryUpdateTodoEntityMockData()

	suite.sqlMock.ExpectBegin()

	suite.sqlMock.ExpectExec(regexp.QuoteMeta(`UPDATE "`+common.TodoTable+`" SET `+`"`+common.TodoTitleCol+`"`+`=$1,`+`"`+common.TodoDescriptionCol+`"`+`=$2,`+
		`"`+common.TodoStatusCol+`"`+`=$3,`+`"`+common.TodoUpdatedAtCol+`"`+
		`=$4 `+` WHERE `+common.TodoIDCol+` = $5 `)).WithArgs(
		mockEntity.Title,
		mockEntity.Description,
		mockEntity.Status,
		util.TestMatchTime{},
		mockQuery.ID,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := suite.repository.UpdateTodoRepository(context.Background(), suite.mockGormDB.Begin(), mockQuery, mockEntity)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
}

func (suite *TestUpdateRepositoryTestSuite) Test_Error_Something_Went_Wrong() {
	mockEntity := mockdata.UpdateTodoEntityMockData()
	mockQuery := mockdata.QueryUpdateTodoEntityMockData()

	suite.sqlMock.ExpectBegin()

	suite.sqlMock.ExpectExec(regexp.QuoteMeta(`UPDATE "`+common.TodoTable+`" SET `+`"`+common.TodoTitleCol+`"`+`=$1,`+`"`+common.TodoDescriptionCol+`"`+`=$2,`+
		`"`+common.TodoStatusCol+`"`+`=$3,`+`"`+common.TodoUpdatedAtCol+`"`+
		`=$4 `+` WHERE `+common.TodoIDCol+` = $5 `)).WithArgs(
		mockEntity.Title,
		mockEntity.Description,
		mockEntity.Status,
		util.TestMatchTime{},
		mockQuery.ID,
	).WillReturnError(common.ErrDBUpdateTodo)

	result, err := suite.repository.UpdateTodoRepository(context.Background(), suite.mockGormDB.Begin(), mockQuery, mockEntity)
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), common.ErrDBUpdateTodo.Error())

}
