package main

// import "github.com/davecgh/go-spew/spew"

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var convertOptions ConvertOptions
	var convertCmd = &cobra.Command{
		Use:   "convert SOURCE_KML",
		Short: "Convert feminicide-related KML into JSON",
		Long:  `Convert feminicide-related KML into JSON (long)`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			convertOptions.InputKml = args[0]
			err := convert(convertOptions)
			if err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
				os.Exit(1)
			}
		},
	}
	convertCmd.Flags().StringVarP(&convertOptions.OutputJson, "output", "o", "-", "Output file name for KML")

	var fetchOptions FetchOptions
	var fetchCmd = &cobra.Command{
		Use:   "fetch YEAR",
		Short: "Fetch KML data for feminicide in specified year",
		Long:  "Fetch KML data for feminicide in specified year",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fetchOptions.Year = args[0]
			err := fetch(fetchOptions)
			if err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
				os.Exit(1)
			}
		},
	}
	fetchCmd.Flags().StringVarP(&fetchOptions.OutputKml, "output", "o", "-", "Output file name for JSON")

	var rootCmd = &cobra.Command{Use: "fi-cli"}
	rootCmd.AddCommand(convertCmd, fetchCmd)
	rootCmd.Execute()
}
