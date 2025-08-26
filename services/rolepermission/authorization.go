package rolepermission

import (
	"QA-Game/param/roleparam"
	"QA-Game/repository/contracts"
	"QA-Game/repository/mysql"
	"QA-Game/response"
	"QA-Game/response/richerror"
	"QA-Game/response/successresponse"
	"QA-Game/validation/rolepermvalidation"
	"net/http"
)

type Service struct {
	RoleValidationService rolepermvalidation.RoleValidation
	ErrorResponse         response.Response
	SuccessResponse       response.Response
	RoleRepository        contracts.RoleRepository
}

func New() Service {
	return Service{
		RoleValidationService: rolepermvalidation.New(),
		ErrorResponse:         richerror.NewErrorResponse(),
		SuccessResponse:       successresponse.NewSuccessResponse(),
		RoleRepository:        mysql.NewRoleRepo(),
	}
}

func (s Service) Store(roleParam roleparam.StoreRoleParam) response.Response {

	validateRoleResult, validateRoleData := s.RoleValidationService.Validate(roleParam)

	if !validateRoleResult {
		return s.ErrorResponse.SetData(validateRoleData).SetStatus(http.StatusBadRequest).Build()
	}

	storeRoleResult := s.RoleRepository.Store(roleParam)

	return s.ErrorResponse.SetData(storeRoleResult).SetStatus(http.StatusBadRequest).Build()
}
