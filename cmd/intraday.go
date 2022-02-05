/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// intradayCmd represents the intraday command
var intradayCmd = &cobra.Command{
	Use:   "intraday",
	Short: "Intraday query api",
}

func init() {
	rootCmd.AddCommand(intradayCmd)

}
