package main

import (
	"log"

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

	err = utility.InitDataDirectory(env)
	if err != nil {
		log.Fatal(err)
	}

	logPaths, err := utility.GetLogPaths(env)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(logPaths)
}
