package usecase_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/domain/mocks"
	"todo-clean/usecase"
	"todo-clean/util/mockdata"
)

func TestGetTodoUseCase(t *testing.T) {
	suite.Run(t, new(TestGetUseCaseTestSuite))
}

type TestGetUseCaseTestSuite struct {
	suite.Suite

	sqlMock    sqlmock.Sqlmock
	sqlMockDB  *sql.DB
	mockGormDB *gorm.DB

	repositoryMock *mocks.TodoRepository
	useCase        domain.TodoUseCase

	mockDataModel []domain.GetTodoEntity
}

func (suite *TestGetUseCaseTestSuite) SetupSuite() {

}

func (suite *TestGetUseCaseTestSuite) TeardownSuite() {

}

func (suite *TestGetUseCaseTestSuite) SetupTest() {
	var err error

	// Init DB Mock Package
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

	suite.repositoryMock = new(mocks.TodoRepository)
	suite.useCase = usecase.NewUseCase(suite.repositoryMock)
	suite.mockDataModel = []domain.GetTodoEntity{
		mockdata.GetTodoEntityMockData(),
	}

}

func (suite *TestGetUseCaseTestSuite) TearDownTest() {

}

func (suite *TestGetUseCaseTestSuite) Test_Happy() {
	tx := context.Background()
	key := suite.mockDataModel[0].ID
	suite.repositoryMock.On("GetTodoRepository", tx, &key).Return(suite.mockDataModel, nil)

	result, err := suite.useCase.GetTodoUseCase(tx, &key)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), suite.mockDataModel, result)
}

func (suite *TestGetUseCaseTestSuite) Test_Error_Something_Went_Wrong() {
	tx := context.Background()
	key := suite.mockDataModel[0].ID
	suite.repositoryMock.On("GetTodoRepository", tx, &key).Return(nil, common.ErrGetTodo)

	result, err := suite.useCase.GetTodoUseCase(tx, &suite.mockDataModel[0].ID)

	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), common.ErrGetTodo.Error())
}
