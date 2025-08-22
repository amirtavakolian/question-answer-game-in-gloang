package auth

import (
	"QA-Game/dto/playerdto"
	"QA-Game/repository/contracts"
	"QA-Game/repository/mysql"
	"QA-Game/response"
	"QA-Game/response/richerror"
	"QA-Game/response/successresponse"
	"QA-Game/services/jwttoken"
	"QA-Game/validation/authvalidation"
	"github.com/labstack/echo/v4"
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

func (auth AuthService) Register(c echo.Context) response.Response {

	registerPlayerDTO := playerdto.PlayerRegister{}
	registerPlayerValidation := authvalidation.PlayerRegister{}

	if err := c.Bind(&registerPlayerDTO); err != nil {
		//todo => extend the error response struct to hold & log system errors for developers, then remove them from users
		return auth.ErrorResponse.SetMessage(err.Error()).SetStatus(http.StatusInternalServerError).Build()
	}

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

func (auth AuthService) Login(c echo.Context) response.Response {

	playerLoginDTO := playerdto.PlayerLogin{}

	playerLoginValidation := authvalidation.PlayerLogin{}

	if err := c.Bind(&playerLoginDTO); err != nil {
		//todo => extend the error response struct to hold & log system errors for developers, then remove them from users
		return auth.ErrorResponse.SetMessage(err.Error()).SetStatus(500).Build()
	}

	validationResult, validationMessage := playerLoginValidation.Validate(playerLoginDTO)

	if !validationResult {
		return auth.ErrorResponse.SetData(validationMessage).SetStatus(http.StatusBadRequest).Build()
	}

	phoneNumber, password, err := auth.PlayerRepo.FindPlayerByPhoneNumber(playerLoginDTO.PhoneNumber)

	if err != nil {
		return auth.ErrorResponse.SetMessage(err.Error()).SetStatus(404).Build()
	}

	// todo => the password must be hashed
	if phoneNumber != playerLoginDTO.PhoneNumber || password != playerLoginDTO.Password {
		return auth.ErrorResponse.SetMessage("phone number or password is wrong").Build()
	}

	loginResponse := LoginResponse{
		AccessToken:  jwttoken.NewJwtToken().CreateAccessToken(playerLoginDTO.PhoneNumber),
		RefreshToken: jwttoken.NewJwtToken().CreateRefreshToken(playerLoginDTO.PhoneNumber),
	}

	return auth.SuccessResponse.SetData(loginResponse).SetStatus(200).Build()
}
