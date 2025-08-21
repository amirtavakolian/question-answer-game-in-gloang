package auth

import (
	"QA-Game/dto/playerdto"
	"QA-Game/repository/contracts"
	"QA-Game/repository/mysql"
	"QA-Game/response/richerror"
	"QA-Game/response/successresponse"
	"QA-Game/services/jwttoken"
	"QA-Game/validation/authvalidation"
	"encoding/json"
	"io"
	"net/http"
)

type AuthService struct {
	PlayerRepo      contracts.PlayerRepository
	ErrorResponse   richerror.ErrorResponse
	SuccessResponse successresponse.SuccessResponse
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewAuthService() AuthService {
	return AuthService{
		PlayerRepo:      mysql.NewPlayerRepo(),
		ErrorResponse:   *richerror.NewErrorResponse(),
		SuccessResponse: *successresponse.NewSuccessResponse(),
	}
}

func (auth AuthService) Register(req *http.Request) string {

	if req.Method != http.MethodPost {
		return "Post method only."
	}

	registerPlayerDTO := playerdto.PlayerRegister{}

	registerPlayerValidation := authvalidation.PlayerRegister{}

	requestData, err := io.ReadAll(req.Body)

	if err != nil {
		//todo => extend the error response struct to hold & log system errors for developers, then remove them from users
		return auth.ErrorResponse.SetMessage(err.Error()).SetStatus(500).Buid()
	}

	json.Unmarshal(requestData, &registerPlayerDTO)

	validationResult := registerPlayerValidation.ValidatePhoneNumber(registerPlayerDTO.PhoneNumber)

	if !validationResult {
		return auth.ErrorResponse.SetMessage("Phone number is not valid").SetStatus(402).Buid()
	}

	res, msg := auth.PlayerRepo.IsPhoneNumberExist(registerPlayerDTO.PhoneNumber)

	if !res {
		return auth.ErrorResponse.SetMessage(msg.Error()).SetStatus(400).Buid()
	}
	// todo => validate first name => should not be empty
	// todo => hash password before storing in database
	// todo => move IsPhoneNumberExist to validation package (registerPlayerValidation)

	playerEntity, storeError := auth.PlayerRepo.Store(registerPlayerDTO)

	if storeError != nil {
		return auth.ErrorResponse.SetMessage(storeError.Error()).SetStatus(400).Buid()
	}

	return auth.SuccessResponse.SetMessage("Player created successfully").SetStatus(200).SetData(playerEntity).Buid()
}

func (auth AuthService) Login(req *http.Request) string {

	if req.Method != http.MethodPost {
		return "Post method only."
	}

	playerLoginDTO := playerdto.PlayerLogin{}

	playerLoginValidation := authvalidation.PlayerLogin{}

	RequestBodyData, RequestBodyDataErr := io.ReadAll(req.Body)

	if RequestBodyDataErr != nil {
		return auth.ErrorResponse.SetMessage("System error happened").SetStatus(500).Buid()
	}

	json.Unmarshal(RequestBodyData, &playerLoginDTO)

	result := playerLoginValidation.ValidatePhoneNumber(playerLoginDTO.PhoneNumber)

	if !result {
		return auth.ErrorResponse.SetMessage("phone number must be 11 charecters.").Buid()
	}

	result = playerLoginValidation.ValidatePassword(playerLoginDTO.Password)

	if !result {
		return auth.ErrorResponse.SetMessage("Password must more then 5 charecters.").Buid()
	}

	phoneNumber, password, err := auth.PlayerRepo.FindPlayerByPhoneNumber(playerLoginDTO.PhoneNumber)

	if err != nil {
		return auth.ErrorResponse.SetMessage(err.Error()).SetStatus(404).Buid()
	}

	// todo => the password must be hashed
	if phoneNumber != playerLoginDTO.PhoneNumber || password != playerLoginDTO.Password {
		return auth.ErrorResponse.SetMessage("phone number or password is wrong").Buid()
	}

	loginResponse := LoginResponse{
		AccessToken:  jwttoken.NewJwtToken().CreateAccessToken(playerLoginDTO.PhoneNumber),
		RefreshToken: jwttoken.NewJwtToken().CreateRefreshToken(playerLoginDTO.PhoneNumber),
	}

	return auth.SuccessResponse.SetData(loginResponse).SetStatus(200).Buid()
}
