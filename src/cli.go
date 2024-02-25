package main

import (
	"log"

	clidapters "github.com/leesolway/powerwave/src/adapters/cli"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "powerwave",
	Short: "Powerwave is a tool for managing power meters.",
	Long:  `Powerwave is a comprehensive tool for managing power meters through CLI and HTTP server.`,
}

// ExecuteCLI executes the root command, which in turn handles the execution of all child commands.
func ExecuteCLI() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal("Error executing CLI command:", err)
	}
}

func init() {
	rootCommand.AddCommand(metersByCustomerCmd)
	rootCommand.AddCommand(meterReadingByDateCmd)
}

var metersByCustomerCmd = &cobra.Command{
	Use:   "getmeters [customer]",
	Short: "Get meters by customer name",
	Args:  cobra.ExactArgs(1),
	Run:   clidapters.AdapterForMetersByCustomer,
}

var meterReadingByDateCmd = &cobra.Command{
	Use:   "getreading [serialID] [date]",
	Short: "Get meter reading by serial ID and date",
	Args:  cobra.ExactArgs(2),
	Run:   clidapters.AdapterForMeterReading,
}
