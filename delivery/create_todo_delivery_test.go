package delivery_test

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-clean/common"
	"todo-clean/delivery"
	"todo-clean/domain"
	"todo-clean/domain/mocks"
	"todo-clean/lib/error_lib"
	"todo-clean/util/mockdata"
)

func TestCreateTodoDelivery(t *testing.T) {
	suite.Run(t, new(TestCreateDeliveryTestSuite))
}

type TestCreateDeliveryTestSuite struct {
	suite.Suite
	ginEngine *gin.Engine

	handler           domain.TodoUseCaseInterface
	useCaseMock       *mocks.TodoUseCaseInterface
	request           domain.CreateTodoEntityRequest
	createEntityModel domain.CreateTodoEntity
}

func (suite *TestCreateDeliveryTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	suite.ginEngine = gin.New()
	suite.useCaseMock = new(mocks.TodoUseCaseInterface)
	delivery.NewHandler(suite.ginEngine, suite.useCaseMock)
}

func (suite *TestCreateDeliveryTestSuite) TearDownSuite() {

}

func (suite *TestCreateDeliveryTestSuite) SetupTest() {

}

func (suite *TestCreateDeliveryTestSuite) TearDownTest() {

}

func (suite *TestCreateDeliveryTestSuite) Test_Happy() {
	var err error
	res := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(res)
	mockEntity := mockdata.CreateTodoEntityMockData()
	suite.useCaseMock.On("CreateTodoUseCase", context.Background(), mock.AnythingOfType("domain.CreateTodoEntityRequest")).Return(&mockEntity, nil)

	reqBody, err := json.Marshal(mockdata.CreateTodoDeliveryRequestMockData())
	assert.NoError(suite.T(), err)

	reader := strings.NewReader(string(reqBody))

	c.Request, err = http.NewRequest(http.MethodPost, common.APIGroup+common.APITodoCreatPath, reader)
	assert.NoError(suite.T(), err)
	suite.ginEngine.ServeHTTP(res, c.Request)
	assert.Equal(suite.T(), http.StatusCreated, res.Code)
}

func (suite *TestCreateDeliveryTestSuite) Test_Error_Something_Went_Wrong() {
	var err error
	res := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(res)

	expectedError := error_lib.WrapError(common.ErrBadRequest.Error(), err)

	suite.useCaseMock.On("CreateTodoUseCase", context.Background(), mock.AnythingOfType("domain.CreateTodoEntityRequest")).Return(nil, expectedError)

	reqBody, err := json.Marshal(mockdata.CreateTodoDeliveryRequestMockData())
	assert.NoError(suite.T(), err)
	reader := strings.NewReader(string(reqBody))

	c.Request, err = http.NewRequest(http.MethodPost, common.APIGroup+common.APITodoCreatPath, reader)
	assert.NoError(suite.T(), err)
	suite.ginEngine.ServeHTTP(res, c.Request)
	assert.Equal(suite.T(), http.StatusBadRequest, res.Code)
}
