package http

import (
	"freefolks-fc/player_information/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type playerInformationHandler struct {
	pu usecase.PlayerInformationUsecase
}

func NewPlayerInformationHandler(e *echo.Echo, pu usecase.PlayerInformationUsecase) {
	handler := &playerInformationHandler{
		pu: pu,
	}

	playerInformation := e.Group("/player-information")

	playerInformation.GET("/fetch", handler.FindAll)
}

func (pih *playerInformationHandler) FindAll(e echo.Context) error {
	playerInformations, err := pih.pu.FindAll()

	if err != nil {
		panic(err)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   playerInformations,
	})
}
