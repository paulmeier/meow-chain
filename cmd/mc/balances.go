package main

import (
	"fmt"
	"github.com/paulmeier/meow-chain/database"
	"os"

	"github.com/spf13/cobra"
)

func balancesCmd() *cobra.Command {
	var balancesCmd = &cobra.Command{
		Use:   "balances",
		Short: "Interact with balances (list ...)",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("bad")
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	balancesCmd.AddCommand(balancesListCmd)

	return balancesCmd
}

var balancesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all balances",
	Run: func(cmd *cobra.Command, args []string) {
		state, err := database.NewStateFromDisk()
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}
		defer state.Close()

		fmt.Println("Account Balances")
		fmt.Println("----------------")
		fmt.Println("")

		for account, balance := range state.Balances {
			fmt.Println(fmt.Sprintf("%s: %d", account, balance))
		}
	},
}
