package clinics

import (
	"net/http"
	"testing"

	"github.com/olguncengiz/scratchpay-clinicsearch/internal/auth"
	"github.com/olguncengiz/scratchpay-clinicsearch/internal/test"
	"github.com/olguncengiz/scratchpay-clinicsearch/pkg/log"
)

func TestAPI(t *testing.T) {
	logger, _ := log.NewForTest()
	router := test.MockRouter(logger)

	repo := NewRepository("../../dental-clinics.json", "../../vet-clinics.json", logger)

	RegisterHandlers(router.Group(""), NewService(repo, logger), auth.MockAuthHandler, logger)
	header := auth.MockAuthHeader()

	tests := []test.APITestCase{
		{"query dental invalid name", "POST", "/dentalClinics", `{"name":"test"}`, header, http.StatusOK, "null"},
		{"query dental new york", "POST", "/dentalClinics", `{"state":"New York"}`, header, http.StatusOK, "[\r\n    {\r\n        \"name\": \"Cleveland Clinic\",\r\n        \"stateName\": \"New York\",\r\n        \"availability\": {\r\n            \"from\": \"11:00\",\r\n            \"to\": \"22:00\"\r\n        }\r\n    }\r\n]"},
		{"query vet kansas", "POST", "/vetClinics", `{"state":"KS"}`, header, http.StatusOK, "[\r\n    {\r\n        \"clinicName\": \"German Pets Clinics\",\r\n        \"stateCode\": \"KS\",\r\n        \"opening\": {\r\n            \"from\": \"08:00\",\r\n            \"to\": \"20:00\"\r\n        }\r\n    }\r\n]"},
	}
	for _, tc := range tests {
		test.Endpoint(t, router, tc)
	}
}
