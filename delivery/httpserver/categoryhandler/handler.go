package categoryhandler

import (
	"QA-Game/param/categoryparam"
	"QA-Game/response"
	"QA-Game/response/richerror"
	"QA-Game/response/successresponse"
	"QA-Game/services/categoryservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	ErrorResponse   response.Response
	SuccessResponse response.Response
	CategoryService categoryservice.Service
}

func NewCategoryHandler() Handler {
	return Handler{
		ErrorResponse:   richerror.NewErrorResponse(),
		SuccessResponse: successresponse.NewSuccessResponse(),
		CategoryService: categoryservice.NewCategoryService(),
	}
}

func (h Handler) Store(c echo.Context) error {

	categoryStoreParam := categoryparam.CategoryStore{}

	if err := c.Bind(&categoryStoreParam); err != nil {
		result := h.ErrorResponse.SetMessage(err.Error()).SetStatus(http.StatusInternalServerError).Build()

		return c.JSON(result.GetStatus(), result)
	}

	result := h.CategoryService.Store(categoryStoreParam)

	return c.JSON(result.GetStatus(), result)
}
