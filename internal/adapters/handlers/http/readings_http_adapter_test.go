package httpadapters

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/leesolway/powerwave/internal/core/domain"
)

// AdapterForMeterReadingTestSuite is the struct for our test suite.
type AdapterForMeterReadingTestSuite struct {
	suite.Suite
	mockService *domain.MockPowerMeterService
	router      *gin.Engine
}

// SetupSuite runs once before the test suite starts.
func (suite *AdapterForMeterReadingTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	suite.router = gin.New()
	suite.mockService = new(domain.MockPowerMeterService)
	suite.router.GET("/readings/:serialID/:date", AdapterForMeterReading(suite.mockService))
}

// SetupTest runs before each test in the suite.
func (suite *AdapterForMeterReadingTestSuite) SetupTest() {
	suite.mockService.ExpectedCalls = nil
	suite.mockService.Calls = nil
}

// AfterTest runs after each test in the suite.
func (suite *AdapterForMeterReadingTestSuite) AfterTest(_, _ string) {
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *AdapterForMeterReadingTestSuite) TestAdapterForMeterReading() {
	tests := []struct {
		description      string
		serialID         string
		date             string
		setupMock        func()
		expectedStatus   int
		expectedResponse string
	}{
		{
			description: "Valid reading fetch",
			serialID:    "1111-1111-1111",
			date:        "2022-01-01",
			setupMock: func() {
				suite.mockService.On("GetMeterReadingBySerialIDAndDate", "1111-1111-1111", mock.AnythingOfType("time.Time")).Return(domain.MeterReading{
					SerialID:    "1111-1111-1111",
					Date:        "2022-01-01",
					KWhForDay:   20,
					KWhForMonth: 620,
				}, nil)
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"serialID":"1111-1111-1111","reading":{"SerialID":"1111-1111-1111","Date":"2022-01-01","KWhForDay":20,"KWhForMonth":620}}`,
		},
		{
			description:      "Invalid date format",
			serialID:         "1111-1111-1111",
			date:             "01-01-2022", // Incorrect date format
			setupMock:        func() {},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `{"error":"Invalid date format"}`,
		},
		{
			description: "Service error on fetching reading",
			serialID:    "1111-1111-1111",
			date:        "2022-01-01",
			setupMock: func() {
				// Simulate an error returned by the service when attempting to fetch the meter reading
				suite.mockService.On("GetMeterReadingBySerialIDAndDate", "1111-1111-1111", mock.AnythingOfType("time.Time")).Return(domain.MeterReading{}, errors.New("internal error"))
			},
			expectedStatus:   http.StatusInternalServerError,
			expectedResponse: `{"error":"Failed to fetch reading"}`,
		},
	}

	for _, test := range tests {
		suite.T().Run(test.description, func(t *testing.T) {
			// Reset the mock before setting it up for each test
			suite.mockService.ExpectedCalls = nil
			suite.mockService.Calls = nil

			test.setupMock()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/readings/%s/%s", test.serialID, test.date), nil)
			suite.router.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatus, w.Code)
			assert.JSONEq(t, test.expectedResponse, w.Body.String())
		})
	}
}

// TestAdapterForMeterReadingTestSuite runs the test suite.
func TestAdapterForMeterReadingTestSuite(t *testing.T) {
	suite.Run(t, new(AdapterForMeterReadingTestSuite))
}
