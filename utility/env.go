package utility

import (
	"errors"
	"os"
	"strconv"
)

const (
	timberData = ".timber"
	timberPort = 36036
	maxStreams = 32
)

// Environment holds sanitized environment variables
type Environment struct {
	TimberData string
	TimberPort int
	MaxStreams int
}

// GetEnv reads environment variables, sanitizes them and returns an Environment struct
func GetEnv() (*Environment, error) {
	// Creating environment struct
	env := &Environment{
		TimberData: timberData,
		TimberPort: timberPort,
		MaxStreams: maxStreams,
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

	// Read MAX_STREAMS environment
	streamsS := os.Getenv("MAX_STREAMS")
	if streamsS != "" {
		streams, err := strconv.Atoi(streamsS)
		if err != nil {
			return nil, errors.New("MAX_STREAMS must be an integer")
		}
		if streams <= 0 {
			return nil, errors.New("MAX_STREAMS must be an greater than 0")
		}
		env.MaxStreams = streams
	}

	return env, nil
}
