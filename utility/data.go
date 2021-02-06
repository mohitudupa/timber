package utility

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
)

// InitDataDirectory prepares the folder TIMBER_DATA
func InitDataDirectory(env *Environment) error {
	// Check if path exists and create if not found
	_, err := os.Stat(env.TimberData)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(env.TimberData, os.ModePerm)
		if err != nil {
			return errors.New("Error creating data directory at, " + env.TimberData)
		}
	} else if err != nil {
		return errors.New("Error reading data directory at, " + env.TimberData)
	}

	return nil
}

// GetLogFiles returns the list of logfiles in TIMBER_DATA
func GetLogFiles(env *Environment) ([]string, error) {
	// Read TIMBER_DATA folder
	files, err := ioutil.ReadDir(env.TimberData)
	if err != nil {
		return nil, errors.New("Error reading data directory at, " + env.TimberData)
	}

	// Populate log paths
	logFiles := []string{}
	for _, file := range files {
		filePath := path.Join(env.TimberData, file.Name())

		fileStat, err := os.Stat(filePath)
		if err != nil {
			continue
		}
		if fileStat.Mode().IsRegular() {
			logFiles = append(logFiles, file.Name())
		}
	}

	return logFiles, nil
}
