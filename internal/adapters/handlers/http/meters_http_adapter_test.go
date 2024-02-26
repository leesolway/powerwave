package httpadapters

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/leesolway/powerwave/internal/core/domain"
)

// AdapterForMetersByCustomerTestSuite defines the suite for testing AdapterForMetersByCustomer.
type AdapterForMetersByCustomerTestSuite struct {
	suite.Suite
	mockService *domain.MockPowerMeterService
	router      *gin.Engine
}

// SetupSuite runs once before the test suite starts.
func (suite *AdapterForMetersByCustomerTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	suite.router = gin.Default()
	suite.mockService = new(domain.MockPowerMeterService)
	suite.router.GET("/meters/:customer", AdapterForMetersByCustomer(suite.mockService))
}

// BeforeTest is run before each test in the suite.
func (suite *AdapterForMetersByCustomerTestSuite) BeforeTest(suiteName, testName string) {
	suite.mockService.ExpectedCalls = nil
	suite.mockService.Calls = nil
}

func (suite *AdapterForMetersByCustomerTestSuite) TestAdapterForMetersByCustomer() {
	tests := []struct {
		description    string
		customerName   string
		mockReturn     []domain.PowerMeter
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			description:  "Successful fetch",
			customerName: "Aquaflow",
			mockReturn: []domain.PowerMeter{
				{SerialID: "1111-1111-1111", Building: "Treatment Plant A", Customer: "Aquaflow", DailyKWh: 20},
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"SerialID":"1111-1111-1111","Building":"Treatment Plant A","Customer":"Aquaflow","DailyKWh":20}]`,
		},
		{
			description:    "Customer not found",
			customerName:   "Nonexistent",
			mockReturn:     nil,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"message":"No meters found"}`,
		},
		{
			description:    "Internal server error",
			customerName:   "ErrorCase",
			mockError:      errors.New("internal error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"Failed to fetch meters"}`,
		},
	}

	for _, test := range tests {
		suite.Run(test.description, func() {
			suite.mockService.On("GetMetersByCustomerName", test.customerName).Return(test.mockReturn, test.mockError)

			req := httptest.NewRequest("GET", "/meters/"+test.customerName, nil)
			w := httptest.NewRecorder()

			suite.router.ServeHTTP(w, req)

			suite.Equal(test.expectedStatus, w.Code)
			if test.expectedBody != "" {
				suite.JSONEq(test.expectedBody, w.Body.String())
			}
		})
	}
}

func TestAdapterForMetersByCustomerTestSuite(t *testing.T) {
	suite.Run(t, new(AdapterForMetersByCustomerTestSuite))
}
