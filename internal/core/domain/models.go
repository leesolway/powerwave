package domain

// PowerMeter represents a power meter device with its associated data.
type PowerMeter struct {
	SerialID string
	Building string
	Customer string
	DailyKWh float64
}

// MeterReading represents a reading for a power meter on a specific date.
type MeterReading struct {
	SerialID    string
	Date        string
	KWhForDay   float64
	KWhForMonth float64
}

// Mock data
var meters = []PowerMeter{
	{SerialID: "1111-1111-1111", Building: "Treatment Plant A", Customer: "Aquaflow", DailyKWh: 20},
	{SerialID: "1111-1111-2222", Building: "Treatment Plant B", Customer: "Aquaflow", DailyKWh: 30},
	{SerialID: "1111-1111-3333", Building: "Student Halls", Customer: "Albers Facilities Management", DailyKWh: 40},
}
