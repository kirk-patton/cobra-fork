/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// childCmd represents the child command
var childCmd = &cobra.Command{
	Use:   "child",
	Short: "Some daemon process...",
	Long: `Some long running daemon like process...
Typically, something that opens and listens on a socket

Using sleep...`,

	Run: func(cmd *cobra.Command, args []string) {
		sleepSec := 300
		fmt.Printf("child called: sleeping %v sec", sleepSec)
		time.Sleep(time.Second * time.Duration(sleepSec))
	},
}

func init() {
	rootCmd.AddCommand(childCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// childCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// childCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
