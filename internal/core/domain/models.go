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
