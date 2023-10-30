package main

import (
	"go-base/bootstrap"
	"go-base/configs"
)

// main is the entry point of the Go program.
//
// It initializes the app using the bootstrap.App() function
// and then runs the app on the specified port.
func main() {
	app := bootstrap.App()

	app.Run(configs.Port)
}
