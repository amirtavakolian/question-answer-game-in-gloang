package profile

import (
	"QA-Game/repository/contracts"
	"QA-Game/repository/mysql"
	"QA-Game/response"
	"QA-Game/response/richerror"
	"QA-Game/response/successresponse"
	"QA-Game/services/jwttoken"
	"github.com/labstack/echo/v4"
)

type ProfileService struct {
	PlayerRepo      contracts.PlayerRepository
	ProfileRepo     contracts.ProfileRepository
	ErrorResponse   richerror.ErrorResponse
	SuccessResponse successresponse.SuccessResponse
	JwtService      jwttoken.JwtService
}

func NewProfileService() ProfileService {
	return ProfileService{
		PlayerRepo:      mysql.NewPlayerRepo(),
		ProfileRepo:     mysql.NewProfileRepo(),
		ErrorResponse:   *richerror.NewErrorResponse(),
		SuccessResponse: *successresponse.NewSuccessResponse(),
		JwtService:      jwttoken.NewJwtToken(),
	}
}

func (profileService ProfileService) GetPlayerProfile(c echo.Context) response.Response {

	getJwtResult, getJwtMsg, claims := profileService.JwtService.Get(c.Request().Header.Get(echo.HeaderAuthorization))

	if !getJwtResult {
		return profileService.ErrorResponse.SetMessage(getJwtMsg).Build()
	}

	playerProfile := profileService.ProfileRepo.GetPlayerProfile(claims.PhoneNumber)

	if !playerProfile.Status {
		return profileService.ErrorResponse.SetMessage(playerProfile.Message).SetStatus(404).Build()
	}

	return profileService.SuccessResponse.SetData(playerProfile.Data).Build()
}
