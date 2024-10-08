package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing Pokerepl")
	os.Exit(0)
	return nil
}
