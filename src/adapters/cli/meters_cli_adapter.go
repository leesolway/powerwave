package clidapters

import (
	"fmt"

	"github.com/leesolway/powerwave/src/domain"
	"github.com/spf13/cobra"
)

// AdapterForMetersByCustomer adapts CLI requests to domain logic for fetching meters by customer
func AdapterForMetersByCustomer(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: powerwave getmeters [customer]")
		return
	}
	customer := args[0]

	meters, err := domain.GetMetersByCustomerName(customer)
	if err != nil {
		fmt.Printf("Error fetching meters: %v\n", err)
		return
	}

	if len(meters) == 0 {
		fmt.Println("No meters found")
		return
	}

	for _, meter := range meters {
		fmt.Printf("SerialID: %s, Building: %s\n", meter.SerialID, meter.Building)
	}
}