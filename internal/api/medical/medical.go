package medical

import (
	"awesomeProject/internal/model"
	"awesomeProject/internal/repository"
	"awesomeProject/pkg/generated/medical"
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/samber/lo"
)

type Implementation struct {
	medical.UnimplementedMedicalServiceServer
	medicalRepo repository.MedicalRepo
}

func NewImplementation(medicalRepo repository.MedicalRepo) *Implementation {
	return &Implementation{medicalRepo: medicalRepo}
}

func (i *Implementation) CreateHospital(ctx context.Context, request *medical.CreateHospitalRequest) (*medical.Hospital, error) {
	hospital, err := i.medicalRepo.CreateHospital(request.GetName())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create hospital")
	}
	return &medical.Hospital{
		HospitalId: hospital.HospitalID.String(),
		Name:       hospital.Name,
	}, nil
}

func (i *Implementation) CreatePatient(ctx context.Context, request *medical.CreatePatientRequest) (*medical.Patient, error) {
	patient, err := i.medicalRepo.CreatePatient(request.GetName(), request.GetAge())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create patient")
	}
	return &medical.Patient{
		PatientId: patient.PatientID.String(),
		Name:      patient.Name,
		Age:       patient.Age,
	}, nil
}

func (i *Implementation) RegisterPatient(ctx context.Context, request *medical.RegisterPatientRequest) (*emptypb.Empty, error) {
	err := i.medicalRepo.RegisterPatient(
		uuid.MustParse(request.GetPatientId()),
		uuid.MustParse(request.GetHospitalId()),
		request.GetRegisterTime().AsTime(),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to register patient")
	}
	return &emptypb.Empty{}, nil
}

func (i *Implementation) DischargePatient(ctx context.Context, request *medical.DiscardPatientRequest) (*emptypb.Empty, error) {
	err := i.medicalRepo.RegisterPatient(
		uuid.MustParse(request.GetPatientId()),
		uuid.MustParse(request.GetHospitalId()),
		request.GetDiscardTime().AsTime(),
	)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, "failed to register patient")
	}
	return &emptypb.Empty{}, nil
}

func (i *Implementation) GetStatOfPatients(ctx context.Context, request *medical.GetStatOfPatientsRequest) (*medical.PatientsStats, error) {
	stat, err := i.medicalRepo.GetStatOfPatient(
		request.GetFrom().AsTime(),
		request.GetTo().AsTime(),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get stat")
	}
	return &medical.PatientsStats{
		PatientStat: lo.Map(stat, func(it model.PatientStats, _ int) *medical.PatientsStat {
			return &medical.PatientsStat{
				HospitalId:       it.HospitalID.String(),
				Name:             it.Name,
				NumberOfPatients: it.NumberOfPatients,
				AverageAge:       it.AverageAge,
				NumberOfRegister: it.NumberOfRegister,
				NumberOfDiscard:  it.NumberOfDiscard,
			}
		}),
	}, nil
}
