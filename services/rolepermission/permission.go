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

type PermissionService struct {
	PermissionValidationService rolepermvalidation.PermissionValidation
	ErrorResponse               response.Response
	SuccessResponse             response.Response
	PermissionRepository        contracts.PermissionRepository
	AssignValidation            rolepermvalidation.AssignValidation
}

func NewPermissionService() PermissionService {
	return PermissionService{
		PermissionValidationService: rolepermvalidation.NewPermissionValidation(),
		ErrorResponse:               richerror.NewErrorResponse(),
		SuccessResponse:             successresponse.NewSuccessResponse(),
		PermissionRepository:        mysql.NewPermissionRepo(),
		AssignValidation:            rolepermvalidation.NewAssignValidation(),
	}
}

func (s PermissionService) Store(permissionParam rolepermissionparam.StorePermissionParam) response.Response {

	validatePermissionResult, validatePermissionData := s.PermissionValidationService.Validate(permissionParam)

	if !validatePermissionResult {
		return s.ErrorResponse.SetData(validatePermissionData).SetStatus(http.StatusBadRequest).Build()
	}

	storePermissionResult := s.PermissionRepository.Store(permissionParam)

	return s.ErrorResponse.SetData(storePermissionResult).SetStatus(http.StatusBadRequest).Build()
}

func (s PermissionService) PermToRole(assignPermToRoleParams rolepermissionparam.AssignPermissionToRoleParam) response.Response {

	paramsValidationResult, paramsValidationData := s.AssignValidation.Validate(assignPermToRoleParams)

	if !paramsValidationResult {
		return s.ErrorResponse.SetData(paramsValidationData).Build()
	}

	assignResult := s.PermissionRepository.AssignPermToRole(assignPermToRoleParams)

	return s.ErrorResponse.SetMessage(assignResult.Message).Build()
}
