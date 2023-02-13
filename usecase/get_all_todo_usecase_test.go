package usecase_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/domain/mocks"
	"todo-clean/usecase"
	"todo-clean/util/mockdata"
)

func TestGetAllTodoUseCase(t *testing.T) {
	suite.Run(t, new(TestGetAllUseCaseTestSuite))
}

type TestGetAllUseCaseTestSuite struct {
	suite.Suite

	sqlMock    sqlmock.Sqlmock
	sqlMockDB  *sql.DB
	mockGormDB *gorm.DB

	repositoryMock *mocks.TodoRepository
	useCase        domain.TodoUseCase

	mockDataModel []domain.GetTodoEntity
}

func (suite *TestGetAllUseCaseTestSuite) SetupSuite() {

}

func (suite *TestGetAllUseCaseTestSuite) TeardownSuite() {

}

func (suite *TestGetAllUseCaseTestSuite) SetupTest() {
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

func (suite *TestGetAllUseCaseTestSuite) TearDownTest() {

}

func (suite *TestGetAllUseCaseTestSuite) Test_Happy() {
	tx := context.Background()
	suite.repositoryMock.On("GetAllTodoRepository", tx).Return(suite.mockDataModel, nil)

	result, err := suite.useCase.GetAllTodoUseCase(tx)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), suite.mockDataModel, result)
}

func (suite *TestGetAllUseCaseTestSuite) Test_Error_Something_Went_Wrong() {
	tx := context.Background()

	suite.repositoryMock.On("GetAllTodoRepository", tx).Return(nil, common.ErrGetAllTodo)

	result, err := suite.useCase.GetAllTodoUseCase(tx)
	log.Println("result =>>", result)
	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), common.ErrGetAllTodo.Error())
}
