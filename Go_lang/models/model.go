package model

import (
	"time"
)

type HospitalPatient struct {
	ID                  int       `json:"id,omitempty"`
	MedicalRecordNumber string    `json:"medicalRecordNumber"`
	Name                string    `json:"name"`
	DateOfBirth         time.Time `json:"dateOfBirth,omitempty"`
	Gender              string    `json:"gender,omitempty"`
	NextOfKin           struct {
		Name               string `json:"nextOfKinName"`
		ContactInformation string `json:"nextOfKinContact"`
	} `json:"nextOfKin"`
	AdmissionDate     time.Time  `json:"admissionDate"`
	DischargeDate     *time.Time `json:"dischargeDate,omitempty"`
	CurrentStatus     string     `json:"currentStatus"`
	MedicalHistory    string     `json:"medicalHistory,omitempty" sensitive`
	CurrentMedication []string   `json:"currentMedication,omitempty" sensitive`
}
