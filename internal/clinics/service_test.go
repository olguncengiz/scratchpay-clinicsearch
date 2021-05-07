package clinics

import (
	"context"
	"errors"

	"github.com/olguncengiz/scratchpay-clinicsearch/internal/entity"
)

var errCRUD = errors.New("error crud")

type mockRepository struct {
	dentalClinics []entity.DentalClinic
	vetClinics    []entity.VetClinic
}

func (m mockRepository) QueryDentalClinics(ctx context.Context, dc entity.DentalClinic) ([]entity.DentalClinic, error) {
	return m.dentalClinics, nil
}

func (m mockRepository) QueryVetClinics(ctx context.Context, vc entity.VetClinic) ([]entity.VetClinic, error) {
	return m.vetClinics, nil
}
