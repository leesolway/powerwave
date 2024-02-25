package domain

import (
	"errors"
	"time"

	"github.com/leesolway/powerwave/internal/core/utils"
)

// PowerMeterService defines the interface for operations on power meters.
type PowerMeterService interface {
	GetMetersByCustomerName(customerName string) ([]PowerMeter, error)
	GetMeterReadingBySerialIDAndDate(serialID string, date time.Time) (MeterReading, error)
}

// DefaultPowerMeterService is a concrete implementation of PowerMeterService.
type DefaultPowerMeterService struct{}

// GetMetersByCustomerName retrieves meters for a given customer name.
func (s *DefaultPowerMeterService) GetMetersByCustomerName(customerName string) ([]PowerMeter, error) {
	var result []PowerMeter
	for _, meter := range meters {
		if meter.Customer == customerName {
			result = append(result, meter)
		}
	}

	return result, nil
}

// GetDailyKWhBySerialIDAndDate calculates the kWh consumed for a given meter on a specific date.
func (s *DefaultPowerMeterService) GetMeterReadingBySerialIDAndDate(serialID string, date time.Time) (MeterReading, error) {
	for _, meter := range meters {
		if meter.SerialID == serialID {

			totalDays := utils.DaysInMonth(date)
			totalKWhForMonth := float64(totalDays) * meter.DailyKWh

			return MeterReading{
				SerialID:    meter.SerialID,
				Date:        date.Format("2006-01-02"),
				KWhForDay:   meter.DailyKWh,
				KWhForMonth: totalKWhForMonth,
			}, nil
		}
	}

	return MeterReading{}, errors.New("meter not found")
}
