package model

import "github.com/google/uuid"

type Hospital struct {
	HospitalID uuid.UUID `db:"hospital_id"`
	Name       string    `db:"name"`
}
