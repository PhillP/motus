package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/phillp/motus/streamapi/app"
)

func main() {
	// Create service
	service := goa.New("Stream Statistics API")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "OrdinalValues" controller
	c := NewOrdinalValuesController(service)
	app.MountOrdinalValuesController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
