package clidapters

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/leesolway/powerwave/internal/core/domain"
)

// AdapterForMeterReading adapts CLI requests to domain logic for fetching meter readings by date
func AdapterForMeterReading(powerMeterService domain.PowerMeterService, cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: powerwave getreading [serialID] [date]")
		return
	}

	serialID := args[0]
	date := args[1]

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println("Invalid date format:", date)
		return
	}

	reading, err := powerMeterService.GetMeterReadingBySerialIDAndDate(serialID, parsedDate)
	if err != nil {
		fmt.Println("Error fetching reading:", err)
		return
	}

	fmt.Printf("SerialID: %s, Reading: %+v\n", serialID, reading)
}
