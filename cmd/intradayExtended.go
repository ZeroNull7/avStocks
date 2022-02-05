/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// IntradayExtendedCmd represents the extended command
var IntradayExtendedCmd = &cobra.Command{
	Use:   "extended",
	Short: "Intraday extended api calls",
}

func init() {
	intradayCmd.AddCommand(IntradayExtendedCmd)
}
