package clinics

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}
	res.logger.Info("r: %s", r)

	r.Post("/dentalClinics", res.queryDentalClinics)
	r.Post("/vetClinics", res.queryVetClinics)

	//r.Use(authHandler)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) queryDentalClinics(c *routing.Context) error {
	var input ClinicSearchRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	dcs, err := r.service.QueryDentalClinics(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(dcs, http.StatusOK)
}

func (r resource) queryVetClinics(c *routing.Context) error {
	ctx := c.Request.Context()
	var input ClinicSearchRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	vcs, err := r.service.QueryVetClinics(ctx, input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(vcs, http.StatusOK)
}
