package delivery_test

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-clean/common"
	"todo-clean/delivery"
	"todo-clean/domain"
	"todo-clean/domain/mocks"
	"todo-clean/lib/errorLib"
	"todo-clean/util/mockdata"
)

func TestGetTodoDelivery(t *testing.T) {
	suite.Run(t, new(TestGetDeliveryTestSuite))
}

type TestGetDeliveryTestSuite struct {
	suite.Suite
	ginEngine *gin.Engine

	useCaseMock *mocks.TodoUseCase
	res         *httptest.ResponseRecorder
}

func (suite *TestGetDeliveryTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	suite.ginEngine = gin.New()
	suite.useCaseMock = new(mocks.TodoUseCase)
	delivery.NewHandler(suite.ginEngine, suite.useCaseMock)
}

func (suite *TestGetDeliveryTestSuite) TearDownSuite() {

}

func (suite *TestGetDeliveryTestSuite) SetupTest() {
	suite.res = httptest.NewRecorder()

}

func (suite *TestGetDeliveryTestSuite) TearDownTest() {
	suite.useCaseMock.ExpectedCalls = []*mock.Call{}
}

func (suite *TestGetDeliveryTestSuite) Test_Happy_With_Key() {
	id := uuid.New().String()
	mockEntity := mockdata.GetTodoEntityMockData()
	suite.useCaseMock.On("GetTodoUseCase", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*uuid.UUID")).
		Return([]domain.GetTodoEntity{mockEntity}, nil)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, common.APIGroup+common.APIAllTodoGetPath+id, nil)
	log.Println("req", req)
	assert.NoError(suite.T(), err)
	suite.ginEngine.ServeHTTP(suite.res, req)
	assert.Equal(suite.T(), http.StatusOK, suite.res.Code)

}

func (suite *TestGetDeliveryTestSuite) Test_Happy_With_No_Key() {
	mockEntity := mockdata.GetTodoEntityMockData()

	suite.useCaseMock.On("GetTodoUseCase", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*uuid.UUID")).
		Return([]domain.GetTodoEntity{mockEntity}, nil)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, common.APIGroup+common.APIAllTodoGetPath, nil)
	assert.NoError(suite.T(), err)
	suite.ginEngine.ServeHTTP(suite.res, req)
	assert.Equal(suite.T(), http.StatusOK, suite.res.Code)
}

func (suite *TestGetDeliveryTestSuite) Test_Error_With_Key() {
	var err error
	expectedError := errorLib.WrapError(common.ErrGetAllTodo.Error(), err)

	suite.useCaseMock.On("GetTodoUseCase", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*uuid.UUID")).
		Return(nil, expectedError)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, common.APIGroup+common.APIAllTodoGetPath, nil)

	assert.NoError(suite.T(), err)
	suite.ginEngine.ServeHTTP(suite.res, req)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.res.Code)
}

func (suite *TestGetDeliveryTestSuite) Test_Error_With_No_Key() {
	id := uuid.New().String()
	var err error
	expectedError := errorLib.WrapError(common.ErrGetAllTodo.Error(), err)

	suite.useCaseMock.On("GetTodoUseCase", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*uuid.UUID")).
		Return(nil, expectedError)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, common.APIGroup+common.APIAllTodoGetPath+id, nil)

	assert.NoError(suite.T(), err)
	suite.ginEngine.ServeHTTP(suite.res, req)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.res.Code)
}
