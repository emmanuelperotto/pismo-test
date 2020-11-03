package main

import (
	"github.com/emmanuelperotto/pismo-test/api/config"
)

// TODO: Check all exported stuff and see it it's really necessary
func main() {
	config.SetupDB()
	config.RunServer()
}
