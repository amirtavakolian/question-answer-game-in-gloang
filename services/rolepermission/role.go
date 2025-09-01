package rolepermission

import (
	"QA-Game/param/rolepermissionparam"
	"QA-Game/repository/contracts"
	"QA-Game/repository/mysql"
	"QA-Game/response"
	"QA-Game/response/richerror"
	"QA-Game/response/successresponse"
	"QA-Game/validation/rolepermvalidation"
	"net/http"
)

type RoleService struct {
	RoleValidationService rolepermvalidation.RoleValidation
	ErrorResponse         response.Response
	SuccessResponse       response.Response
	RoleRepository        contracts.RoleRepository
	AssignValidation      rolepermvalidation.AssignValidation
}

func New() RoleService {
	return RoleService{
		RoleValidationService: rolepermvalidation.NewRoleValidation(),
		ErrorResponse:         richerror.NewErrorResponse(),
		SuccessResponse:       successresponse.NewSuccessResponse(),
		RoleRepository:        mysql.NewRoleRepo(),
	}
}

func (s RoleService) Store(roleParam rolepermissionparam.StoreRoleParam) response.Response {

	validateRoleResult, validateRoleData := s.RoleValidationService.Validate(roleParam)

	if !validateRoleResult {
		return s.ErrorResponse.SetData(validateRoleData).SetStatus(http.StatusBadRequest).Build()
	}

	storeRoleResult := s.RoleRepository.Store(roleParam)

	return s.ErrorResponse.SetData(storeRoleResult).SetStatus(http.StatusBadRequest).Build()
}

