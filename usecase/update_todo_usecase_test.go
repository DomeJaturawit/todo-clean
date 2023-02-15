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
	"todo-clean/usecase"
	"todo-clean/util/mockdata"
)

func TestUpdateTodoUseCase(t *testing.T) {
	suite.Run(t, new(TestUpdateUseCaseTestSuite))
}

type TestUpdateUseCaseTestSuite struct {
	suite.Suite

	sqlMock    sqlmock.Sqlmock
	sqlMockDB  *sql.DB
	mockGormDB *gorm.DB

	repositoryMock    *mocks.TodoRepository
	useCase           domain.TodoUseCase
	keyID             domain.UpdateTodoQueryEntity
	updateEntityModel domain.UpdateTodoEntity
	mockDataModel     []domain.GetTodoEntity
}

func (suite *TestUpdateUseCaseTestSuite) SetupSuite() {

}

func (suite *TestUpdateUseCaseTestSuite) TeardownSuite() {

}

func (suite *TestUpdateUseCaseTestSuite) SetupTest() {
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

	suite.updateEntityModel = mockdata.UpdateTodoEntityMockData()
	suite.keyID = mockdata.QueryUpdateTodoEntityMockData()
	suite.mockDataModel = []domain.GetTodoEntity{
		mockdata.GetTodoEntityMockData(),
	}
}

func (suite *TestUpdateUseCaseTestSuite) TearDownTest() {

}

func (suite *TestUpdateUseCaseTestSuite) Test_Happy() {

	appCtx := context.Background()
	mockEntity := mockdata.UpdateTodoEntityMockData()
	key := suite.mockDataModel[0].ID
	tx := suite.mockGormDB.Begin()

	suite.repositoryMock.On("Begin").Return(tx, nil).
		On("GetTodoRepository", context.Background(), &key).Return(suite.mockDataModel, nil).
		On("UpdateTodoRepository", appCtx, tx, mock.AnythingOfType("domain.UpdateTodoQueryEntity"), mock.AnythingOfType("domain.UpdateTodoEntity")).Return(&mockEntity, nil).
		On("Commit", tx).Return(nil)

	result, err := suite.useCase.UpdateTodoUseCase(appCtx, suite.keyID, &suite.updateEntityModel)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), suite.updateEntityModel.Title, result.Title)
	assert.Equal(suite.T(), suite.updateEntityModel.Status, result.Status)
	assert.Equal(suite.T(), suite.updateEntityModel.Description, result.Description)
	assert.WithinDuration(suite.T(), suite.updateEntityModel.UpdatedAt, result.UpdatedAt, time.Second)
}

func (suite *TestUpdateUseCaseTestSuite) Test_Update_Error() {
	appCtx := context.Background()
	key := suite.mockDataModel[0].ID
	tx := suite.mockGormDB.Begin()

	suite.repositoryMock.On("Begin").Return(tx, nil).
		On("GetTodoRepository", context.Background(), &key).Return(suite.mockDataModel, nil).
		On("UpdateTodoRepository", appCtx, tx, mock.AnythingOfType("domain.UpdateTodoQueryEntity"), mock.AnythingOfType("domain.UpdateTodoEntity")).Return(nil, common.ErrUseCaseUpdateTodo).
		On("Commit", tx).Return(nil)

	result, err := suite.useCase.UpdateTodoUseCase(appCtx, suite.keyID, &suite.updateEntityModel)

	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), common.ErrUseCaseUpdateTodo.Error())
}

func (suite *TestUpdateUseCaseTestSuite) Test_Update_Get_Data_Error() {
	appCtx := context.Background()
	key := suite.mockDataModel[0].ID
	tx := suite.mockGormDB.Begin()

	suite.repositoryMock.On("Begin").Return(tx, nil).
		On("GetTodoRepository", context.Background(), &key).Return(nil, common.ErrDBGetTodo)

	result, err := suite.useCase.UpdateTodoUseCase(appCtx, suite.keyID, &suite.updateEntityModel)

	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), common.ErrDBGetTodo.Error())
}
