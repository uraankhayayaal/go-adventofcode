package helpers

import (
	"errors"
	"fmt"
	"os"
)

func IsPathOrFileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	fmt.Printf("Schrodinger's file: error occurred but it might exist. Error: %v\n", err)
	return false
}
