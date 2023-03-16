/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

const (
	PidFile = "./child.pid"
)

// parentCmd represents the parent command
var parentCmd = &cobra.Command{
	Use:   "parent",
	Short: "Parent Process",
	Long:  `Call the child process using exec.Command().Start()`,
	Run: func(cmd *cobra.Command, args []string) {
		pidFH, _ := os.Create(PidFile)
		defer pidFH.Close()

		childCmdName := "child"
		childCmd := exec.Command("./cobra-fork", childCmdName)
		err := childCmd.Start()
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", childCmd)

		pid := childCmd.Process.Pid
		fmt.Fprintf(pidFH, "%v\n", pid)
		fmt.Printf("started child process: PID: %v\n", pid)
	},
}

func init() {
	rootCmd.AddCommand(parentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
