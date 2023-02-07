package usecase_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"todo-clean/domain"
	"todo-clean/domain/mocks"
	"todo-clean/usecase"
	"todo-clean/util/mockdata"
)

func TestCreateTodoUseCase(t *testing.T) {
	suite.Run(t, new(TestCreateUseCaseTestSuite))
}

type TestCreateUseCaseTestSuite struct {
	suite.Suite

	sqlMock    sqlmock.Sqlmock
	sqlMockDB  *sql.DB
	mockGormDB *gorm.DB

	repositoryMock *mocks.TodoRepositoryInterface
	useCase        domain.TodoUseCaseInterface

	createEntityModel domain.CreateTodoEntity
	request           domain.CreateTodoEntityRequest
}

func (suite *TestCreateUseCaseTestSuite) SetupSuite() {

}

func (suite *TestCreateUseCaseTestSuite) TeardownSuite() {

}

func (suite *TestCreateUseCaseTestSuite) SetupTest() {
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

	suite.repositoryMock = new(mocks.TodoRepositoryInterface)
	suite.useCase = usecase.NewUseCase(suite.repositoryMock)

	suite.request = mockdata.CreateTodoEntityRequestMockData()
	suite.createEntityModel = mockdata.CreateTodoUseCaseEntityMockData(suite.request)

}

func (suite *TestCreateUseCaseTestSuite) TearDownTest() {

}

func (suite *TestCreateUseCaseTestSuite) Test_Happy() {

	appCtx := context.Background()
	mockEntity := mockdata.CreateTodoUseCaseEntityMockData(suite.request)

	tx := suite.mockGormDB.Begin()

	suite.repositoryMock.On("Begin", appCtx).Return(tx, nil).
		On("CreateTodoRepository", tx, mock.AnythingOfType("domain.CreateTodoEntity")).Return(&mockEntity, nil).
		On("Commit").Return(tx, nil)

	result, err := suite.useCase.CreateTodoUseCase(appCtx, suite.request)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), suite.createEntityModel.Title, result.Title)
	assert.Equal(suite.T(), suite.createEntityModel.Status, result.Status)
	assert.Equal(suite.T(), suite.createEntityModel.Description, result.Description)
	//TODO Should Equal All Time
	assert.Equal(suite.T(), suite.createEntityModel.CreatedAt.Day(), result.CreatedAt.Day())

}

func (suite *TestCreateUseCaseTestSuite) Test_Error_Something_Went_Wrong() {
	appCtx := context.Background()
	tx := suite.mockGormDB.Begin()

	suite.repositoryMock.On("Begin", appCtx).Return(tx, nil).
		On("CreateTodoRepository", tx, mock.AnythingOfType("domain.CreateTodoEntity")).Return(nil, mockdata.UseCaseError).
		On("RollBack").Return(nil, mockdata.UseCaseError)

	result, err := suite.useCase.CreateTodoUseCase(appCtx, suite.request)

	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)

	assert.Contains(suite.T(), err.Error(), mockdata.UseCaseError.Error())
}
