package main

// import "github.com/davecgh/go-spew/spew"

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var cmdConvert = &cobra.Command{
		Use:   "convert SOURCE_KML DESTIMATION_JSON",
		Short: "Convert feminicide-related KML into JSON",
		Long:  `Convert feminicide-related KML into JSON (long)`,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			kmlFile := args[0]
			jsonFile := args[1]
			fmt.Printf("Converting KML %s into JSON %s...\n", kmlFile, jsonFile)
			convert(kmlFile, jsonFile)
		},
	}

	var cmdFetch = &cobra.Command{
		Use:   "fetch YEAR",
		Short: "Fetch KML data for feminicide in specified year",
		Long:  "Fetch KML data for feminicide in specified year",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			year := args[0]
			fmt.Printf("Fetching KML for year %s\n", year)
			fetch(year)
		},
	}
	var rootCmd = &cobra.Command{Use: "fi-cli"}
	rootCmd.AddCommand(cmdConvert, cmdFetch)
	rootCmd.Execute()
}
