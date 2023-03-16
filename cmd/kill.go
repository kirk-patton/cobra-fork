/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// killCmd represents the kill command
var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kill the process created by parent",
	Long: `Read the pid file created by the parent
create a new exec.Command().Process() from it.`,

	Run: func(cmd *cobra.Command, args []string) {
		p, _ := os.ReadFile(PidFile)
		// trim off the newline from the pid
		p = p[:len(p)-1]

		pid, _ := strconv.Atoi(string(p))

		// Find the pid and cread a Process from it
		process, _ := os.FindProcess(pid)

		// Nuke the process
		process.Kill()
		fmt.Printf("killing: %v\n", pid)
	},
}

func init() {
	rootCmd.AddCommand(killCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// killCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// killCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
