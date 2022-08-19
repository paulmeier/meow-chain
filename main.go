package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Account string

type Tx struct {
	From  Account `json:"from"`
	To    Account `json:"to"`
	Value uint    `json:"value"`
	Data  string  `json:"data"`
}

func (t Tx) IsReward() bool {
	return t.Data == "reward"
}

type State struct {
	Balances map[Account]uint
	txMempool []Tx

	dbFile *os.File
}

func NewStateFromDisk() (*State, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	generateFilePath := filepath.Join(cwd, "database", "genesis.json")
	gen, err := loadGenesis(generateFilePath)
	if err != nil {
		return nil, err
	}

	balances := make(map[Account]uint)
	for account, balance := range gen.Balances {
		balances[account] = balance
	}
}

func main() {

}
