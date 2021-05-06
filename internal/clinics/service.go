package clinics

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Service encapsulates usecase logic for clinics.
type Service interface {
	QueryDentalClinics(ctx context.Context, input ClinicSearchRequest) ([]entity.DentalClinic, error)
	QueryVetClinics(ctx context.Context, input ClinicSearchRequest) ([]entity.VetClinic, error)
}

/*
// DentalClinic represents the data about a dental clinic.
type DentalClinic struct {

}

type VetClinic struct {

}
*/

// ClinicSearchRequest represents a clinic search request.
type ClinicSearchRequest struct {
	Name   string `json:"name" required:"true"`
	State  string `json:"state"`
	Opens  string `json:"opens"`
	Closes string `json:"closes"`
}

// Validate validates the ClinicSearchRequest fields.
func (m ClinicSearchRequest) Validate() error {
	if m.Opens > m.Closes {
		return validation.NewError("200", "Opens can not be later than Closes")
	}
	return nil
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new clinics service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// QueryDentalClinics returns the dental clinics with the specified parameters.
func (s service) QueryDentalClinics(ctx context.Context, req ClinicSearchRequest) ([]entity.DentalClinic, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	dc := entity.DentalClinic{
		Name:      req.Name,
		StateName: req.State,
		Availability: entity.Availability{
			From: req.Opens,
			To:   req.Closes,
		},
	}
	items, err := s.repo.QueryDentalClinics(ctx, dc)
	if err != nil {
		return nil, err
	}
	return items, nil
}

// QueryVetClinics returns the vet clinics with the specified parameters.
func (s service) QueryVetClinics(ctx context.Context, req ClinicSearchRequest) ([]entity.VetClinic, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	vc := entity.VetClinic{
		ClinicName: req.Name,
		StateCode:  req.State,
		Opening: entity.Opening{
			From: req.Opens,
			To:   req.Closes,
		},
	}
	items, err := s.repo.QueryVetClinics(ctx, vc)
	if err != nil {
		return nil, err
	}
	return items, nil
}
