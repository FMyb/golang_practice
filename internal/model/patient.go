package model

import "github.com/google/uuid"

type Patient struct {
	PatientID uuid.UUID `db:"patient_id"`
	Name      string    `db:"name"`
	Age       int32     `db:"age"`
}

type PatientStats struct {
	HospitalID       uuid.UUID `db:"id"`
	Name             string    `db:"name"`
	NumberOfPatients int32     `db:"number_of_patients"`
	AverageAge       float32   `db:"average_age"`
	NumberOfRegister int32     `db:"number_of_register"`
	NumberOfDiscard  int32     `db:"number_of_discard"`
}
