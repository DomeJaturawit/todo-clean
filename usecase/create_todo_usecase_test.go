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
	"time"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/domain/mocks"
	"todo-clean/lib/error_lib"
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
		On("CreateTodoRepository", appCtx, tx, mock.AnythingOfType("domain.CreateTodoEntity")).Return(&mockEntity, nil).
		On("Commit", tx).Return(nil)

	result, err := suite.useCase.CreateTodoUseCase(appCtx, suite.request)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), suite.createEntityModel.Title, result.Title)
	assert.Equal(suite.T(), suite.createEntityModel.Status, result.Status)
	assert.Equal(suite.T(), suite.createEntityModel.Description, result.Description)
	assert.WithinDuration(suite.T(), suite.createEntityModel.CreatedAt, suite.createEntityModel.CreatedAt, time.Second)
}

func (suite *TestCreateUseCaseTestSuite) Test_Error_Something_Went_Wrong() {
	appCtx := context.Background()
	tx := suite.mockGormDB.Begin()

	expectedError := error_lib.WrapError(common.ErrUseCaseCreateTodo.Error(), common.ErrDBCreateTodoRepo)

	suite.repositoryMock.On("Begin", appCtx).Return(tx, nil).
		On("CreateTodoRepository", appCtx, tx, mock.AnythingOfType("domain.CreateTodoEntity")).Return(nil, expectedError).
		On("Commit").Return(expectedError).On("RollBack").Return(expectedError)

	result, err := suite.useCase.CreateTodoUseCase(appCtx, suite.request)

	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), expectedError.Error())
}
