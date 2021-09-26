package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/mohitudupa/timber/logger"
)

func main() {

	log.Println("Starting Timber...")

	// Load config
	c := logger.NewConfig()
	_, err := os.Stat("./timberconf.json")
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error reading config file. Error: %v.\nExiting\n", err)
	}

	err = c.Load()
	if err != nil {
		log.Println("Error reading ./timberconf.json. Using default configs instead.")
	}

	log.Printf("Config:\nTimberData: %s\nTimberPort: %d\nTimberLogs: %v", c.Data, c.Port, c.Logs)

	// Initialize and perpare DATA directory
	err = logger.InitDataDirectory(c.Data)
	if err != nil {
		log.Fatalf("Error loading log directory at %s. Error: %v.\nExiting.\n", c.Data, err)
	}

	// Setup StreamHandler
	sh := logger.StreamHandler{}

	// Create all log streams
	for _, lf := range c.Logs {
		lp := path.Join(c.Data, lf)
		err := sh.Add(lf, lp)
		if err != nil {
			log.Println(err)
		}
	}
	if len(sh) == 0 {
		log.Fatal("Error no logs open.\nExiting.\n")
	}

	// Closing streams while exiting
	for _, s := range sh {
		defer s.Stop()
	}

	// Starting server
	http.HandleFunc("/log/", sh.ServeHTTP)
	err = http.ListenAndServe(":"+strconv.Itoa(c.Port), nil)
	if err != nil {
		log.Fatalf("Error starting log server. Error: %v.\nExiting\n", err)
	}
}
