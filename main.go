package main

import (
	"log"

	"github.com/mohitudupa/timber/logger"
	"github.com/mohitudupa/timber/utility"
)

func main() {

	log.Println("Hello World!")

	// Read enviroment variables
	env, err := utility.GetEnv()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(env)

	// Initialize and perpare TIMBER_DATA directory
	err = utility.InitDataDirectory(env)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch list of logfile paths
	logFiles, err := utility.GetLogFiles(env)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(logFiles)

	// Create log Pool and attach existing logfiles
	p := logger.NewPool(env)
	for _, logFile := range logFiles {
		p.Attach(logFile)
		defer p.Detach(logFile)
	}

}
