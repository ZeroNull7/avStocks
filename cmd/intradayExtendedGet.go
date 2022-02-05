/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/ZeroNull7/avStocks/pkg/config"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var intradayExtendedGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get list of stocks",
	Run: func(cmd *cobra.Command, args []string) {
		config = config.Get()
	},
}

func init() {
	intradayExtendedGetCmd.AddCommand(IntradayExtendedCmd)
}
