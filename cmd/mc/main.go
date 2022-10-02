package main

import (
	"fmt"

	"os"

	"github.com/spf13/cobra"
)

func main() {
	mcCmd := &cobra.Command{
		Use:   "mc",
		Short: "MeowChain CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	mcCmd.AddCommand(versionCmd)
	mcCmd.AddCommand(balancesListCmd)
	mcCmd.AddCommand(txCmd())

	err := mcCmd.Execute()
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
