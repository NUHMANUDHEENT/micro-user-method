package main

import (
	"golang-application/internal/config"
)

func main() {
	config.LoadEnv()
	config.GrpcSetup()

}
