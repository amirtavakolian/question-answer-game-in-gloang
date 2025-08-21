package profile

import (
	"QA-Game/repository/contracts"
	"QA-Game/repository/mysql"
	"QA-Game/response/richerror"
	"QA-Game/response/successresponse"
	"QA-Game/services/jwttoken"
	"net/http"
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

func (profileService ProfileService) GetPlayerProfile(req *http.Request) string {

	if req.Method != http.MethodGet {
		return profileService.ErrorResponse.SetMessage("Only GET method is allowed").Buid()
	}

	getJwtResult, getJwtMsg, claims := profileService.JwtService.Get(req.Header.Get("Authorization"))

	if !getJwtResult {
		return profileService.ErrorResponse.SetMessage(getJwtMsg).Buid()
	}

	playerProfile := profileService.ProfileRepo.GetPlayerProfile(claims.PhoneNumber)

	if !playerProfile.Status {
		return profileService.ErrorResponse.SetMessage(playerProfile.Message).SetStatus(404).Buid()
	}

	return profileService.SuccessResponse.SetData(playerProfile.Data).Buid()
}
