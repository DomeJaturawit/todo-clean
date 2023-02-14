package repository_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
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

func TestGetTodoRepository(t *testing.T) {
	suite.Run(t, new(TestGetRepositoryTestSuite))
}

type TestGetRepositoryTestSuite struct {
	suite.Suite
	sqlMock       sqlmock.Sqlmock
	sqlMockDB     *sql.DB
	mockGormDB    *gorm.DB
	repository    domain.TodoRepository
	mockDataModel []domain.GetTodoEntity
}

func (suite *TestGetRepositoryTestSuite) SetupSuite() {

}

func (suite *TestGetRepositoryTestSuite) TearDownSuite() {

}

func (suite *TestGetRepositoryTestSuite) SetupTest() {
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
	suite.mockDataModel = []domain.GetTodoEntity{
		mockdata.GetTodoEntityMockData(),
		mockdata.GetTodoEntityMockData(),
	}

	suite.mockDataModel[1].ID = uuid.New()
	suite.mockDataModel[1].Title = mockdata.NewString()
	suite.mockDataModel[1].Status = mockdata.NewString()
	suite.mockDataModel[1].Description = mockdata.NewString()
	suite.mockDataModel[1].CreatedAt = time.Now()

}

func (suite *TestGetRepositoryTestSuite) TearDownTest() {
	assert.NoError(suite.T(), suite.sqlMock.ExpectationsWereMet())
}

func (suite *TestGetRepositoryTestSuite) Test_Happy_With_Key() {
	suite.sqlMock.MatchExpectationsInOrder(true)
	suite.sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "` + common.TodoTable + `" WHERE ` + common.TodoIDCol + ` = $1`)).WithArgs(suite.mockDataModel[0].ID).
		WillReturnRows(sqlmock.NewRows([]string{
			common.TodoIDCol,
			common.TodoTitleCol,
			common.TodoDescriptionCol,
			common.TodoStatusCol,
			common.TodoCreatedAtCol,
		}).AddRow(
			suite.mockDataModel[0].ID,
			suite.mockDataModel[0].Title,
			suite.mockDataModel[0].Description,
			suite.mockDataModel[0].Status,
			suite.mockDataModel[0].CreatedAt,
		).AddRow(
			suite.mockDataModel[1].ID,
			suite.mockDataModel[1].Title,
			suite.mockDataModel[1].Description,
			suite.mockDataModel[1].Status,
			suite.mockDataModel[1].CreatedAt))
	tx := context.Background()

	result, err := suite.repository.GetTodoRepository(tx, &suite.mockDataModel[0].ID)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), len(result), len(suite.mockDataModel))

	assert.Equal(suite.T(), suite.mockDataModel, result)

}

func (suite *TestGetRepositoryTestSuite) Test_Error_Something_Went_Wrong_With_Key() {
	suite.sqlMock.MatchExpectationsInOrder(true)
	suite.sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "` + common.TodoTable + `" WHERE ` + common.TodoIDCol + ` = $1`)).WithArgs(suite.mockDataModel[0].ID).
		WillReturnError(common.ErrDBGetTodo)

	tx := context.Background()
	result, err := suite.repository.GetTodoRepository(tx, &suite.mockDataModel[0].ID)
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), common.ErrDBGetTodo.Error())

}

func (suite *TestGetRepositoryTestSuite) Test_Happy_With_No_Key() {
	suite.sqlMock.MatchExpectationsInOrder(true)
	suite.sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "` + common.TodoTable + `"`)).
		WillReturnRows(sqlmock.NewRows([]string{
			common.TodoIDCol,
			common.TodoTitleCol,
			common.TodoDescriptionCol,
			common.TodoStatusCol,
			common.TodoCreatedAtCol,
		}).AddRow(
			suite.mockDataModel[0].ID,
			suite.mockDataModel[0].Title,
			suite.mockDataModel[0].Description,
			suite.mockDataModel[0].Status,
			suite.mockDataModel[0].CreatedAt,
		).AddRow(
			suite.mockDataModel[1].ID,
			suite.mockDataModel[1].Title,
			suite.mockDataModel[1].Description,
			suite.mockDataModel[1].Status,
			suite.mockDataModel[1].CreatedAt,
		))
	tx := context.Background()

	result, err := suite.repository.GetTodoRepository(tx, nil)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), len(result), len(suite.mockDataModel))

	assert.Equal(suite.T(), suite.mockDataModel, result)
}

func (suite *TestGetRepositoryTestSuite) Test_Error_Something_Went_Wrong_With_No_Key() {
	suite.sqlMock.MatchExpectationsInOrder(true)
	suite.sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "` + common.TodoTable + `"`)).
		WillReturnError(common.ErrDBGetTodo)
	tx := context.Background()
	result, err := suite.repository.GetTodoRepository(tx, nil)
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), common.ErrDBGetTodo.Error())
}
