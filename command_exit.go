package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing Pokerepl")
	os.Exit(0)
	return nil
}
