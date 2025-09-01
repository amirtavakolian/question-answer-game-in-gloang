package auth

import (
	"QA-Game/param/playerparam"
	"QA-Game/repository/contracts"
	"QA-Game/repository/mysql"
	"QA-Game/response"
	"QA-Game/response/richerror"
	"QA-Game/response/successresponse"
	"QA-Game/services/jwttoken"
	"QA-Game/validation/authvalidation"
	"net/http"
)

type AuthService struct {
	PlayerRepo      contracts.PlayerRepository
	ErrorResponse   response.Response
	SuccessResponse response.Response
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewAuthService() AuthService {
	return AuthService{
		PlayerRepo:      mysql.NewPlayerRepo(),
		ErrorResponse:   richerror.NewErrorResponse(),
		SuccessResponse: successresponse.NewSuccessResponse(),
	}
}

func (auth AuthService) Register(registerPlayerDTO playerparam.PlayerRegisterRequest) response.Response {

	registerPlayerValidation := authvalidation.PlayerRegisterValidation{}

	validationResult, validationMessage := registerPlayerValidation.Validate(registerPlayerDTO)

	if !validationResult {
		return auth.ErrorResponse.SetData(validationMessage).SetStatus(http.StatusBadRequest).Build()
	}

	// todo => hash password before storing in database
	playerEntity, storeError := auth.PlayerRepo.Store(registerPlayerDTO)

	if storeError != nil {
		return auth.ErrorResponse.SetMessage(storeError.Error()).SetStatus(http.StatusBadRequest).Build()
	}

	return auth.SuccessResponse.SetMessage("Player created successfully").SetStatus(http.StatusOK).SetData(playerEntity).Build()
}

func (auth AuthService) Login(playerLoginDTO playerparam.PlayerLoginRequest) response.Response {

	playerLoginValidation := authvalidation.PlayerLoginValidation{}

	validationResult, validationMessage := playerLoginValidation.Validate(playerLoginDTO)

	if !validationResult {
		return auth.ErrorResponse.SetData(validationMessage).SetStatus(http.StatusBadRequest).Build()
	}

	result, err := auth.PlayerRepo.FindPlayerByPhoneNumber(playerLoginDTO.PhoneNumber)

	if err != nil {
		return auth.ErrorResponse.SetMessage(err.Error()).SetStatus(404).Build()
	}

	// todo => the password must be hashed
	if result.PhoneNumber != playerLoginDTO.PhoneNumber || result.Password != playerLoginDTO.Password {
		return auth.ErrorResponse.SetMessage("phone number or password is wrong").Build()
	}

	playerLoginDTO.PlayerId = result.PlayerId

	loginResponse := LoginResponse{
		AccessToken:  jwttoken.NewJwtToken().CreateAccessToken(playerLoginDTO),
		RefreshToken: jwttoken.NewJwtToken().CreateRefreshToken(playerLoginDTO),
	}

	return auth.SuccessResponse.SetData(loginResponse).SetStatus(200).Build()
}
