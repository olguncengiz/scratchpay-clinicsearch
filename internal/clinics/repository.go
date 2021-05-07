package clinics

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/olguncengiz/scratchpay-clinicsearch/internal/entity"
	"github.com/olguncengiz/scratchpay-clinicsearch/pkg/log"
)

// Repository encapsulates the logic to access dental clinics from the JSON file.
type Repository interface {
	// QueryDentalClinic returns the list of dental clinics with the given dental clinic parameters.
	QueryDentalClinics(ctx context.Context, dc entity.DentalClinic) ([]entity.DentalClinic, error)
	// QueryVetClinic returns the list of vet clinics with the given vet clinic parameters.
	QueryVetClinics(ctx context.Context, vc entity.VetClinic) ([]entity.VetClinic, error)
}

// repository persists clinics in JSON files
type repository struct {
	dentalClinics []*entity.DentalClinic
	vetClinics    []*entity.VetClinic
	logger        log.Logger
}

// NewRepository creates a new clinics repository
func NewRepository(dcFile, vcFile string, logger log.Logger) Repository {
	// Open our jsonFile for dental clinics
	jsonFileDC, err := os.Open(dcFile)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFileDC.Close()

	// Open our jsonFile for vet clinics
	jsonFileVC, err := os.Open(vcFile)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFileDC.Close()

	byteValue, _ := ioutil.ReadAll(jsonFileDC)

	var dcResult []*entity.DentalClinic
	json.Unmarshal([]byte(byteValue), &dcResult)

	fmt.Println(dcResult)

	byteValue, _ = ioutil.ReadAll(jsonFileVC)
	var vcResult []*entity.VetClinic
	json.Unmarshal([]byte(byteValue), &vcResult)

	fmt.Println(vcResult)

	return repository{dcResult, vcResult, logger}
}

// QueryDentalClinic retrieves the dental clinic records with the specified parameters from the JSON file.
func (r repository) QueryDentalClinics(ctx context.Context, dc entity.DentalClinic) ([]entity.DentalClinic, error) {
	var dcs []entity.DentalClinic
	for _, d := range r.dentalClinics {
		if compareDental(dc, *d) {
			dcs = append(dcs, *d)
		}
	}
	return dcs, nil
}

// QueryVetClinic retrieves the vet clinic records with the specified parameters from the JSON file.
func (r repository) QueryVetClinics(ctx context.Context, vc entity.VetClinic) ([]entity.VetClinic, error) {
	var vcs []entity.VetClinic
	for _, v := range r.vetClinics {
		if compareVet(vc, *v) {
			vcs = append(vcs, *v)
		}
	}
	return vcs, nil
}

func compareDental(query, item entity.DentalClinic) bool {
	if query.Name != "" && item.Name != query.Name {
		return false
	}

	if query.StateName != "" && item.StateName != query.StateName {
		return false
	}

	if query.Availability.From != "" && (item.Availability.From > query.Availability.From || item.Availability.To < query.Availability.From) {
		return false
	}

	if query.Availability.To != "" && (item.Availability.To < query.Availability.To || item.Availability.From > query.Availability.To) {
		return false
	}

	return true
}

func compareVet(query, item entity.VetClinic) bool {
	if query.ClinicName != "" && item.ClinicName != query.ClinicName {
		return false
	}

	if query.StateCode != "" && item.StateCode != query.StateCode {
		return false
	}

	if query.Opening.From != "" && (item.Opening.From > query.Opening.From || item.Opening.To < query.Opening.From) {
		return false
	}

	if query.Opening.To != "" && (item.Opening.To < query.Opening.To || item.Opening.From > query.Opening.To) {
		return false
	}

	return true
}
