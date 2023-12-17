package datastore

import (
	"errors"

	"gofr.dev/pkg/gofr"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"gopkg.in/validator.v2"
)

// HospitalPatient struct defines the schema for hospital patients
type HospitalPatient struct {
	ID            int    `db:"id"`
	Name          string `db:"name"`
	Age           int    `db:"age"`
	Symptoms      string `db:"symptoms"`
	AdmissionDate string `db:"admission_date"`
	Status        string `db:"status"` // e.g., "admitted", "in_treatment", "discharged"
	DischargeDate string `db:"discharge_date"`
	// ... Add any additional fields as needed
}

func New() *HospitalPatient {
	return &HospitalPatient{}
}

// AddPatient adds a new patient entry to the database
func (p *HospitalPatient) AddPatient(ctx *gofr.Context, patient *HospitalPatient) (*HospitalPatient, error) {

	if err := validator.Validate(patient); err != nil {
		return nil, errors.Wrap(err, "invalid patient data")
	}
	query := `INSERT INTO patients (name, age, symptoms, admission_date, status)
			VALUES (?, ?, ?, ?, ?)`
	result, err := ctx.DB().ExecContext(ctx, query, patient.Name, patient.Age,
		patient.Symptoms, patient.AdmissionDate, patient.Status)
	if err != nil {
		return nil, gofr.Wrap(err, "error adding patient")
	}

	// Get the newly generated ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, gofr.Wrap(err, "error getting patient ID")
	}

	patient.ID = int(id)
	return patient, nil
}

// ListPatients retrieves a list of all currently admitted patients
func (p *HospitalPatient) ListPatients(ctx *gofr.Context) ([]*HospitalPatient, error) {
	query := `SELECT * FROM patients WHERE status = ?`
	patients := []*HospitalPatient{}
	rows, err := ctx.DB().QueryContext(ctx, query, "admitted")
	if err != nil {
		return nil, gofr.Wrap(err, "error listing patients")
	}
	defer rows.Close()

	for rows.Next() {
		var patient HospitalPatient
		err := rows.Scan(&patient.ID, &patient.Name, &patient.Age, &patient.Symptoms, &patient.AdmissionDate, &patient.Status, &patient.DischargeDate)
		if err != nil {
			return nil, gofr.Wrap(err, "error scanning patient data")
		}
		patients = append(patients, &patient)
	}

	return patients, nil
}

// UpdatePatientStatus updates a patient's status based on treatment progress
func (p *HospitalPatient) UpdatePatientStatus(ctx *gofr.Context, id int, newStatus string) error {
	query := `UPDATE patients SET status = ? WHERE id = ?`
	_, err := ctx.DB().ExecContext(ctx, query, newStatus, id)
	if err != nil {
		return gofr.Wrapf(err, "error updating patient status for ID %d", id)
	}

	return nil
}

// DischargePatient removes a patient's entry from the database upon discharge
func (p *HospitalPatient) DischargePatient(ctx *gofr.Context, id int, dischargeDate string) error {
	query := `UPDATE patients SET status = ?, discharge_date = ? WHERE id = ?`
	_, err := ctx.DB().ExecContext(ctx, query, "discharged", dischargeDate, id)
	if err != nil {
		return gofr.Wrapf(err, "error discharging patient with ID %d", id)
	}

	return nil
}

// Additional functions and methods can be implemented here based on your specific needs
