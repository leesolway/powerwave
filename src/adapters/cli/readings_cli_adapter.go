package clidapters

import (
	"fmt"
	"time"

	"github.com/leesolway/powerwave/src/domain"
	"github.com/spf13/cobra"
)

// AdapterForMeterReading adapts CLI requests to domain logic for fetching meter readings by date
func AdapterForMeterReading(cmd *cobra.Command, args []string) {
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

	reading, err := domain.GetMeterReadingBySerialIDAndDate(serialID, parsedDate)
	if err != nil {
		fmt.Println("Error fetching reading:", err)
		return
	}

	fmt.Printf("SerialID: %s, Reading: %+v\n", serialID, reading)
}
