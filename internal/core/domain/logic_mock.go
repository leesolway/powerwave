package domain

import (
	"time"

	"github.com/stretchr/testify/mock"
)

// MockPowerMeterService mocks the PowerMeterService for testing purposes.
type MockPowerMeterService struct {
	mock.Mock
}

// GetMetersByCustomerName mocks the GetMetersByCustomerName method of PowerMeterService.
func (m *MockPowerMeterService) GetMetersByCustomerName(customerName string) ([]PowerMeter, error) {
	args := m.Called(customerName)
	return args.Get(0).([]PowerMeter), args.Error(1)
}

// GetMeterReadingBySerialIDAndDate mocks the GetMeterReadingBySerialIDAndDate method of PowerMeterService.
func (m *MockPowerMeterService) GetMeterReadingBySerialIDAndDate(serialID string, date time.Time) (MeterReading, error) {
	args := m.Called(serialID, date)
	return args.Get(0).(MeterReading), args.Error(1)
}
