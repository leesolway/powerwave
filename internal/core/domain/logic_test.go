package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetMetersByCustomerName(t *testing.T) {
	tests := []struct {
		name           string
		customerName   string
		expectedMeters []PowerMeter
	}{
		{
			name:         "Existing Customer",
			customerName: "Aquaflow",
			expectedMeters: []PowerMeter{
				{SerialID: "1111-1111-1111", Building: "Treatment Plant A", Customer: "Aquaflow", DailyKWh: 20},
				{SerialID: "1111-1111-2222", Building: "Treatment Plant B", Customer: "Aquaflow", DailyKWh: 30},
			},
		},
		{
			name:           "Non-existing Customer",
			customerName:   "Nonexistent",
			expectedMeters: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			meters, err := GetMetersByCustomerName(test.customerName)
			assert.NoError(t, err)
			assert.Equal(t, test.expectedMeters, meters)
		})
	}
}

func TestGetMeterReadingBySerialIDAndDate(t *testing.T) {
	date := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name         string
		serialID     string
		expectedKWh  float64
		expectedErr  bool
		errorMessage string
	}{
		{
			name:        "Existing Meter",
			serialID:    "1111-1111-1111",
			expectedKWh: 620, // Assuming 31 days in January
			expectedErr: false,
		},
		{
			name:         "Non-existing Meter",
			serialID:     "Nonexistent",
			expectedKWh:  0,
			expectedErr:  true,
			errorMessage: "meter not found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reading, err := GetMeterReadingBySerialIDAndDate(test.serialID, date)
			if test.expectedErr {
				assert.Error(t, err)
				assert.EqualError(t, err, test.errorMessage)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedKWh, reading.KWhForMonth)
			}
		})
	}
}
