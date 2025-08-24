package playerhandler

import (
	"QA-Game/dto/playerdto"
	"QA-Game/response"
	"QA-Game/response/richerror"
	"QA-Game/response/successresponse"
	"QA-Game/services/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	ErrorResponse   response.Response
	SuccessResponse response.Response
}

func New() *Handler {
	return &Handler{
		ErrorResponse:   richerror.NewErrorResponse(),
		SuccessResponse: successresponse.NewSuccessResponse(),
	}
}

func (h *Handler) PlayerRegister(c echo.Context) error {

	registerPlayerDTO := playerdto.PlayerRegisterRequest{}

	if err := c.Bind(&registerPlayerDTO); err != nil {
		//todo => extend the error response struct to hold & log system errors for developers, then remove them from users
		result := h.ErrorResponse.SetMessage(err.Error()).SetStatus(http.StatusInternalServerError).Build()

		return c.JSON(result.GetStatus(), result)
	}

	authService := auth.NewAuthService()

	result := authService.Register(registerPlayerDTO)

	return c.JSON(result.GetStatus(), result)
}

func (h *Handler) PlayerLogin(c echo.Context) error {

	playerLoginDTO := playerdto.PlayerLoginRequest{}

	if err := c.Bind(&playerLoginDTO); err != nil {
		//todo => extend the error response struct to hold & log system errors for developers, then remove them from users
		result := h.ErrorResponse.SetMessage(err.Error()).SetStatus(500).Build()

		return c.JSON(result.GetStatus(), result)
	}

	authService := auth.NewAuthService()

	result := authService.Login(playerLoginDTO)

	return c.JSON(result.GetStatus(), result)
}
