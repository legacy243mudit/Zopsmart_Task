package main

import (
	"go_lang/datastore"
	"go_lang/handler"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	s := datastore.New()
	h := handler.New(s)

	app.GET("/students/{id}", h.GetPatientByID)
	app.POST("/students", h.CreatePatient)
	app.PUT("/students/{id}", h.validatePatientData)
	app.DELETE("/students/{id}", h.DeletePatient)

	// starting the server on a custom port
	app.Server.HTTP.Port = 9092
	app.Start()
}
