package logger

import (
	"errors"
	"os"
)

func removeBlanks(s []string) []string {
	resp := []string{}
	for _, ele := range s {
		if ele != "" {
			resp = append(resp, ele)
		}
	}
	return resp
}

// InitDataDirectory prepares the folder TIMBER_DATA
func InitDataDirectory(d string) error {
	// Check if path exists and create if not found
	_, err := os.Stat(d)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(d, os.ModePerm)
		if err != nil {
			return errors.New("Error creating data directory at, " + d)
		}
	} else if err != nil {
		return errors.New("Error reading data directory at, " + d)
	}

	return nil
}
