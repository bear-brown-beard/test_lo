package main

import (
	"log"
	"test_lo/internal/di"
)

func main() {
	container, err := di.CreateContainer()
	if err != nil {
		log.Fatal("Failed to initialize container:", err)
	}
	defer container.Logger.Close()

	container.ServerService.Start()
	container.ServerService.HandleStop()
}
