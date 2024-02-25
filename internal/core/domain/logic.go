package domain

import (
	"errors"
	"time"

	"github.com/leesolway/powerwave/internal/core/utils"
)

// Mock data
var meters = []PowerMeter{
	{SerialID: "1111-1111-1111", Building: "Treatment Plant A", Customer: "Aquaflow", DailyKWh: 20},
	{SerialID: "1111-1111-2222", Building: "Treatment Plant B", Customer: "Aquaflow", DailyKWh: 30},
	{SerialID: "1111-1111-3333", Building: "Student Halls", Customer: "Albers Facilities Management", DailyKWh: 40},
}

// GetMetersByCustomerName retrieves meters for a given customer name.
func GetMetersByCustomerName(customerName string) ([]PowerMeter, error) {
	var result []PowerMeter
	for _, meter := range meters {
		if meter.Customer == customerName {
			result = append(result, meter)
		}
	}

	return result, nil
}

// GetDailyKWhBySerialIDAndDate calculates the kWh consumed for a given meter on a specific date.
func GetMeterReadingBySerialIDAndDate(serialID string, date time.Time) (MeterReading, error) {
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
