package main

import (
	"github.com/emmanuelperotto/pismo-test/api/config"
	"github.com/emmanuelperotto/pismo-test/api/repositories"
)

// TODO: Check all exported stuff and see it it's really necessary
func main() {
	repositories.SetupDB()
	config.RunServer()
}
