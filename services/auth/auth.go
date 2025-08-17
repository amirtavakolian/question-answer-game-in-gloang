package auth

import (
	"QA-Game/dto/playerdto"
	"QA-Game/repository/contracts"
	"QA-Game/repository/mysql"
	"QA-Game/response/richerror"
	"QA-Game/response/successresponse"
	"QA-Game/validation"
	"encoding/json"
	"io"
	"net/http"
)

type AuthService struct {
	PlayerRepo      contracts.PlayerRepository
	ErrorResponse   richerror.ErrorResponse
	SuccessResponse successresponse.SuccessResponse
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

}
