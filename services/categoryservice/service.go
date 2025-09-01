package categoryservice

import (
	"QA-Game/param/categoryparam"
	"QA-Game/repository/contracts"
	"QA-Game/repository/mysql"
	"QA-Game/response"
	"QA-Game/response/richerror"
	"QA-Game/response/successresponse"
	"QA-Game/validation/categoryvalidation"
)

type Service struct {
	CategoryValidation categoryvalidation.CategoryStore
	ErrorResponse      response.Response
	SuccessResponse    response.Response
	CategoryRepository contracts.CategoryRepository
}

func NewCategoryService() Service {
	return Service{
		CategoryValidation: categoryvalidation.CategoryStore{},
		ErrorResponse:      richerror.NewErrorResponse(),
		SuccessResponse:    successresponse.NewSuccessResponse(),
		CategoryRepository: mysql.NewCategoryRepo(),
	}
}

func (s Service) Store(categoryStore categoryparam.CategoryStore) response.Response {

	validationResult, validationData := s.CategoryValidation.Validate(categoryStore)

	if !validationResult {
		return s.ErrorResponse.SetData(validationData).Build()
	}

	result := s.CategoryRepository.Store(categoryStore)

	if !result.Status {
		return s.ErrorResponse.SetData(result.Data).Build()
	}

	return s.SuccessResponse.SetMessage("Category created successfully").Build()
}
