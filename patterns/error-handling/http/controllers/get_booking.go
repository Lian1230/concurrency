package controllers

import (
	"context"
	"fmt"
	"github.com/steevehook/http/logging"
	"net/http"

	"github.com/google/uuid"

	"github.com/steevehook/http/models"
	"github.com/steevehook/http/transport"
)

type bookingGetter interface {
	GetBooking(ctx context.Context, req models.GetBookingRequest) (models.Booking, error)
}

func getBooking(service bookingGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logging.Logger()
		id := routeParam(r, "id")
		_, err := uuid.Parse(id)
		if err != nil {
			e := models.FormatValidationError{
				Message: fmt.Sprintf("invalid uuid: %s", id),
			}
			transport.SendHTTPError(w, e)
			return
		}

		req := models.GetBookingRequest{
			ID: id,
		}
		booking, err := service.GetBooking(r.Context(), req)
		if err != nil {
			transport.SendHTTPError(w, err)
			return
		}

		logger.Info("successfully fetched booking")
		transport.SendJSON(w, http.StatusOK, booking)
	}
}
