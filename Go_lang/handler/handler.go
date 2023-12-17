package handler

import (
	model "Go_lang/models"
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
	"gofr.dev/pkg/gofr"

	"go_lang/datastore"
)

type PatientHandler struct {
	store datastore.HospitalPatientInterface
}

func New(s datastore.HospitalPatientInterface) PatientHandler {
	return PatientHandler{store: s}
}

func (h PatientHandler) GetPatientByID(ctx *gofr.Context) (interface{}, error) {
	// Extract and validate ID
	idStr := ctx.PathParam("id")
	id, err := validateID(idStr)
	if err != nil {
		return nil, errors.New("id", err.Error())
	}

	// Fetch patient from store
	patient, err := h.store.GetPatientByID(ctx, id)
	if err != nil {
		return nil, handleError(err)
	}

	// Mask sensitive data before returning response
	patient.MaskSensitiveData()

	return patient, nil
}

func (h PatientHandler) CreatePatient(ctx *gofr.Context) (interface{}, error) {
	// Bind JSON data to model
	var patient model.HospitalPatient
	err := json.NewDecoder(ctx.Request().Body).Decode(&patient)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	// Validate patient data
	err = validatePatientData(&patient)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	// Add patient to store
	newPatient, err := h.store.AddPatient(ctx, &patient)
	if err != nil {
		return nil, handleError(err)
	}

	// Mask sensitive data before returning response
	newPatient.MaskSensitiveData()

	return newPatient, nil
}

// ... Implement handler methods for ListPatients, UpdatePatient, UpdatePatientStatus, and DischargePatient ...

func (h PatientHandler) DeletePatient(ctx *gofr.Context) (interface{}, error) {
	// Extract and validate ID
	idStr := ctx.PathParam("id")
	id, err := validateID(idStr)
	if err != nil {
		return nil, errors.NewParam("id", err.Error())
	}

	// Delete patient from store
	err = h.store.DischargePatient(ctx, id)
	if err != nil {
		return nil, handleError(err)
	}

	return "Patient discharged successfully", nil
}

func validateID(idStr string) (int, error) {
	return strconv.Atoi(idStr)
}

func validatePatientData(patient *model.HospitalPatient) error {
	// Implement specific data validation checks for patient fields like name, age, symptoms, admission date, etc.
	// You can use libraries like `validate` or custom validation functions.
	return nil
}

func handleError(err error) error {
	// Map internal errors to generic messages and error codes for public API responses.
	// Avoid disclosing sensitive details in error messages.
	switch err.(type) {
	case errors.EntityNotFound:
		return errors.NotFound{Message: "Patient not found"}
	default:
		return errors.InternalServerError{Message: "Internal server error"}
	}
}

// ... Implement additional utility functions like patient data masking ...
