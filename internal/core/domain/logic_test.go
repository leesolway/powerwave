package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// PowerMeterServiceTestSuite is the test suite for the PowerMeterService.
type PowerMeterServiceTestSuite struct {
	suite.Suite
	powerMeterService PowerMeterService
}

// SetupTest initializes the test suite for the PowerMeterService.
func (suite *PowerMeterServiceTestSuite) SetupTest() {
	suite.powerMeterService = &DefaultPowerMeterService{}
}

// TestPowerMeterServiceTestSuite runs the test suite for the PowerMeterService.
func TestPowerMeterServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PowerMeterServiceTestSuite))
}

func (suite *PowerMeterServiceTestSuite) TestGetMetersByCustomerName() {
	tests := []struct {
		name           string
		customerName   string
		expectedLength int
		expectedError  bool
	}{
		{"Existing Customer", "Aquaflow", 2, false},
		{"Non-existing Customer", "Nonexistent", 0, false},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			meters, err := suite.powerMeterService.GetMetersByCustomerName(tt.customerName)
			if tt.expectedError {
				assert.Error(suite.T(), err)
			} else {
				assert.NoError(suite.T(), err)
				assert.Len(suite.T(), meters, tt.expectedLength)
			}
		})
	}
}

func (suite *PowerMeterServiceTestSuite) TestGetMeterReadingBySerialIDAndDate() {
	date := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name            string
		serialID        string
		date            time.Time
		expectError     bool
		expectedReading MeterReading
	}{
		{
			name:        "Existing Meter",
			serialID:    "1111-1111-1111",
			date:        date,
			expectError: false,
			expectedReading: MeterReading{
				SerialID:    "1111-1111-1111",
				Date:        "2022-01-01",
				KWhForDay:   20,
				KWhForMonth: 620,
			},
		},
		{
			name:        "Fake Meter",
			serialID:    "fake-serial-id",
			date:        date,
			expectError: true,
		},
	}

	for _, test := range tests {
		suite.Run(test.serialID, func() {
			reading, err := suite.powerMeterService.GetMeterReadingBySerialIDAndDate(test.serialID, test.date)
			if test.expectError {
				suite.Error(err)
			} else {
				suite.NoError(err)

				suite.Assert().Equal(test.expectedReading.SerialID, reading.SerialID)
				suite.Assert().Equal(test.expectedReading.Date, reading.Date)
				suite.Assert().Equal(test.expectedReading.KWhForDay, reading.KWhForDay)
				suite.Assert().Equal(test.expectedReading.KWhForMonth, reading.KWhForMonth)
			}
		})
	}
}
