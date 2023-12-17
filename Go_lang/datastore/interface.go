package datastore

import (
	model "Go_lang/models"

	"gofr.dev/pkg/gofr"
)

// HospitalPatientInterface defines the operations for managing hospital patients
type HospitalPatientInterface interface {
	// AddPatient adds a new patient record to the database
	AddPatient(ctx *gofr.Context, patient *model.HospitalPatient) (*model.HospitalPatient, error)

	// GetPatientByID retrieves a patient record based on its ID
	GetPatientByID(ctx *gofr.Context, id int) (*model.HospitalPatient, error)

	// ListPatients retrieves a list of all currently admitted patients
	ListPatients(ctx *gofr.Context) ([]*model.HospitalPatient, error)

	// UpdatePatient updates an existing patient record with the provided information
	UpdatePatient(ctx *gofr.Context, patient *model.HospitalPatient) (*model.HospitalPatient, error)

	// UpdatePatientStatus updates a patient's status based on treatment progress
	UpdatePatientStatus(ctx *gofr.Context, id int, status string) error

	// DischargePatient removes a patient record from the database upon discharge
	DischargePatient(ctx *gofr.Context, id int) error
}
