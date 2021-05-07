package clinics

import (
	"context"
	"testing"

	"github.com/olguncengiz/scratchpay-clinicsearch/internal/entity"

	"github.com/olguncengiz/scratchpay-clinicsearch/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	logger, _ := log.NewForTest()
	repo := NewRepository("../../dental-clinics.json", "../../vet-clinics.json", logger)

	ctx := context.Background()

	// query
	dc := entity.DentalClinic{}
	dc.StateName = "California"
	dcs, err := repo.QueryDentalClinics(ctx, dc)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(dcs))
}
