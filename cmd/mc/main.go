package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"os"
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

	err := mcCmd.Execute()
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
