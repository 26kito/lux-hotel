package service

import (
	"lux-hotel/entity"
	"lux-hotel/repository"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type MidtransService interface {
	HandleMidtransCallback(c echo.Context) error
}

type midtransService struct {
	MidtransRepository repository.MidtransRepository
}

func NewMidtransService(midtransRepository repository.MidtransRepository) MidtransService {
	return &midtransService{MidtransRepository: midtransRepository}
}

func (ms *midtransService) HandleMidtransCallback(c echo.Context) error {
	// Parse callback payload
	var payload entity.MidtransCallbackResponse

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid payload"})
	}

	// Identify transaction by order_id
	if strings.Contains(payload.OrderID, "TPUP") {
		ms.MidtransRepository.HandleTopUpCallback(payload)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Callback processed"})
}
