package delivery_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-clean/common"
	"todo-clean/delivery"
	"todo-clean/domain/mocks"
	"todo-clean/util/mockdata"
)

func TestUpdateTodoDelivery(t *testing.T) {
	suite.Run(t, new(TestUpdateDeliveryTestSuite))
}

type TestUpdateDeliveryTestSuite struct {
	suite.Suite
	ginEngine *gin.Engine

	useCaseMock *mocks.TodoUseCase
	res         *httptest.ResponseRecorder
}

func (suite *TestUpdateDeliveryTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	suite.ginEngine = gin.New()
	suite.useCaseMock = new(mocks.TodoUseCase)
	delivery.NewHandler(suite.ginEngine, suite.useCaseMock)
}

func (suite *TestUpdateDeliveryTestSuite) TearDownSuite() {

}

func (suite *TestUpdateDeliveryTestSuite) SetupTest() {
	suite.res = httptest.NewRecorder()

}

func (suite *TestUpdateDeliveryTestSuite) TearDownTest() {
	suite.useCaseMock.ExpectedCalls = []*mock.Call{}
}

func (suite *TestUpdateDeliveryTestSuite) Test_Happy() {
	var err error
	id := uuid.New().String()
	c, _ := gin.CreateTestContext(suite.res)
	mockEntity := mockdata.UpdateTodoEntityMockData()
	suite.useCaseMock.On("UpdateTodoUseCase", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("domain.UpdateTodoQueryEntity"), mock.AnythingOfType("*domain.UpdateTodoEntity")).Return(&mockEntity, nil)
	reqBody, err := json.Marshal(mockdata.UpdateTodoDeliveryRequestMockData())
	assert.NoError(suite.T(), err)

	reader := strings.NewReader(string(reqBody))

	c.Request, err = http.NewRequest(http.MethodPatch, common.APIGroup+common.APIUpdateTodoTestPath+id, reader)
	assert.NoError(suite.T(), err)
	suite.ginEngine.ServeHTTP(suite.res, c.Request)
	assert.Equal(suite.T(), http.StatusOK, suite.res.Code)
}

func (suite *TestUpdateDeliveryTestSuite) Test_Error_Bad_Request() {
	var err error
	id := uuid.New().String()

	suite.useCaseMock.On("UpdateTodoUseCase", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("domain.UpdateTodoQueryEntity"),
		mock.AnythingOfType("*domain.UpdateTodoEntity")).Return(nil, common.ErrBadRequest)

	c, _ := gin.CreateTestContext(suite.res)
	reqBody, err := json.Marshal("")
	assert.NoError(suite.T(), err)

	reader := strings.NewReader(string(reqBody))

	c.Request, err = http.NewRequest(http.MethodPatch, common.APIGroup+common.APIUpdateTodoTestPath+id, reader)

	assert.NoError(suite.T(), err)
	suite.ginEngine.ServeHTTP(suite.res, c.Request)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.res.Code)
}

func (suite *TestUpdateDeliveryTestSuite) Test_Error_Internal_Server() {
	var err error
	c, _ := gin.CreateTestContext(suite.res)
	id := uuid.New().String()

	suite.useCaseMock.On("UpdateTodoUseCase", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("domain.UpdateTodoQueryEntity"),
		mock.AnythingOfType("*domain.UpdateTodoEntity")).Return(nil, common.ErrInternal)

	reqBody, err := json.Marshal(mockdata.CreateTodoDeliveryRequestMockData())
	assert.NoError(suite.T(), err)
	reader := strings.NewReader(string(reqBody))

	c.Request, err = http.NewRequest(http.MethodPatch, common.APIGroup+common.APIUpdateTodoTestPath+id, reader)
	assert.NoError(suite.T(), err)
	suite.ginEngine.ServeHTTP(suite.res, c.Request)

	assert.Equal(suite.T(), http.StatusInternalServerError, suite.res.Code)
}
