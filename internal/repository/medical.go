package repository

import (
	"awesomeProject/internal/model"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"time"
)

var ErrNotFound = errors.New("entity not found")

type MedicalRepo interface {
	CreateHospital(name string) (*model.Hospital, error)
	CreatePatient(name string, age int32) (*model.Patient, error)
	RegisterPatient(patientID uuid.UUID, hospitalID uuid.UUID, registerTime time.Time) error
	DiscardPatient(patientID uuid.UUID, hospitalID uuid.UUID, discardTime time.Time) error
	GetStatOfPatient(from time.Time, to time.Time) ([]model.PatientStats, error)
}

type Implementation struct {
	db *sqlx.DB
}

func NewMedicalRepo(db *sqlx.DB) MedicalRepo {
	return &Implementation{db: db}
}

func (i *Implementation) CreateHospital(name string) (*model.Hospital, error) {
	hospital := model.Hospital{
		HospitalID: uuid.New(),
		Name:       name,
	}
	_, err := i.db.NamedExec("insert into hospitals (hospital_id, name) values (:hospital_id, :name)", &hospital)
	if err != nil {
		log.Errorf("failed to create hospital: %v", err)
		return nil, err
	}
	return &hospital, nil
}

func (i *Implementation) CreatePatient(name string, age int32) (*model.Patient, error) {
	patient := model.Patient{
		PatientID: uuid.New(),
		Name:      name,
		Age:       age,
	}
	_, err := i.db.NamedExec("insert into patients (patient_id, name, age) VALUES (:patient_id, :name, :age)", &patient)
	if err != nil {
		log.Errorf("failed to create patient: %v", err)
		return nil, err
	}
	return &patient, err
}

func (i *Implementation) RegisterPatient(patientID uuid.UUID, hospitalID uuid.UUID, registerTime time.Time) error {
	_, err := i.db.Exec(
		"insert into registered_patients (patient_id, hospital_id, register_time) values ($1, $2, $3)",
		patientID, hospitalID, registerTime,
	)
	if err != nil {
		log.Errorf("failed to register patient: %v", err)
		return err
	}
	return nil
}

func (i *Implementation) DiscardPatient(patientID uuid.UUID, hospitalID uuid.UUID, discardTime time.Time) error {
	res, err := i.db.Exec(
		"update registered_patients set discard_time = $3 where patient_id = $1 and hospital_id = $2",
		patientID, hospitalID, discardTime,
	)
	if err != nil {
		return err
	}
	if aff, err := res.RowsAffected(); err != nil || aff == 0 {
		return ErrNotFound
	}
	return nil
}

func (i *Implementation) GetStatOfPatient(from time.Time, to time.Time) ([]model.PatientStats, error) {
	stat := make([]model.PatientStats, 0)
	err := i.db.Select(&stat, "select * from patients_stat($1, $2)", from, to)

	if err != nil {
		log.Errorf("fail to get stat: %v", err)
		return nil, err
	}
	return stat, nil
}
