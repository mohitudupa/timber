package utility

import (
	"errors"
	"os"
	"strconv"
)

const (
	timberData = "~/.timber"
	timberPort = 36036
)

// Environment holds sanitized environment variables
type Environment struct {
	TimberData string
	TimberPort int
}

// GetEnv reads environment variables, sanitizes them and returns an Environment struct
func GetEnv() (*Environment, error) {
	// Creating environment struct
	env := &Environment{
		TimberData: timberData,
		TimberPort: timberPort,
	}
	// Read TIMBER_DATA enviroment
	data := os.Getenv("TIMBER_DATA")
	if data != "" {
		env.TimberData = data
	}

	// Read TIMBER_PORT enviroment
	portS := os.Getenv("TIMBER_PORT")
	if portS != "" {
		port, err := strconv.Atoi(portS)
		if err != nil {
			return nil, errors.New("TIMBER_PORT must be an integer")
		}
		if port <= 0 || port >= 65535 {
			return nil, errors.New("TIMBER_PORT must be an within the range of 0 - 65535")
		}
		env.TimberPort = port
	}

	return env, nil
}
