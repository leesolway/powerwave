package main

import (
	"log"

	"github.com/spf13/cobra"

	clidapters "github.com/leesolway/powerwave/internal/adapters/handlers/cli"
	"github.com/leesolway/powerwave/internal/core/domain"
)

func main() {
	powerMeterService := &domain.DefaultPowerMeterService{}

	var rootCommand = &cobra.Command{
		Use:   "powerwave",
		Short: "Powerwave is a tool for managing power meters.",
		Long:  `Powerwave is a comprehensive tool for managing power meters through CLI.`,
	}

	initCommands(rootCommand, powerMeterService)

	if err := rootCommand.Execute(); err != nil {
		log.Fatal("Error executing CLI command:", err)
	}
}

func initCommands(rootCommand *cobra.Command, powerMeterService domain.PowerMeterService) {
	metersByCustomerCmd := &cobra.Command{
		Use:   "getmeters [customer]",
		Short: "Get meters by customer name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			clidapters.AdapterForMetersByCustomer(powerMeterService, cmd, args)
		},
	}

	meterReadingByDateCmd := &cobra.Command{
		Use:   "getreading [serialID] [date]",
		Short: "Get meter reading by serial ID and date",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			clidapters.AdapterForMeterReading(powerMeterService, cmd, args)
		},
	}

	rootCommand.AddCommand(metersByCustomerCmd, meterReadingByDateCmd)
}
