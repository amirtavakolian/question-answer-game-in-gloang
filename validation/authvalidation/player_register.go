package authvalidation

import (
	"QA-Game/dto/playerdto"
	"QA-Game/repository/mysql"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"regexp"
	"strings"
)

type PlayerRegisterValidation struct{}

func (r PlayerRegisterValidation) Validate(registerPlayerDTO playerdto.PlayerRegisterRequest) (bool, map[string]interface{}) {

	err := validation.ValidateStruct(&registerPlayerDTO,

		validation.Field(&registerPlayerDTO.Name,
			validation.Required,
			validation.Length(5, 200),
		),

		validation.Field(&registerPlayerDTO.PhoneNumber,
			validation.Required,
			validation.Length(6, 50),
			validation.Match(regexp.MustCompile("^09\\d{9}$")).Error("phone number format is not correct"),
			validation.By(UniquePhoneRule(mysql.NewPlayerRepo())),
		),

		validation.Field(&registerPlayerDTO.Password,
			validation.Required,
			validation.Length(6, 200),
		),

		validation.Field(&registerPlayerDTO.Avatar,
			is.Alphanumeric,
		),
	)

	if err != nil {
		splitErrors := strings.Split(err.Error(), ";")

		c := make(map[string]interface{})

		c["errors"] = splitErrors

		return false, c
	}

	return true, nil
}

func UniquePhoneRule(db *mysql.Player) validation.RuleFunc {
	return func(value interface{}) error {
		phoneNumber, ok := value.(string)
		if !ok {
			return validation.NewError("validation_invalid_type", "phone number must be a string")
		}

		var exists bool
		err := db.Connection.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM players WHERE phone_number = ?)", phoneNumber).Scan(&exists)
		if err != nil {
			return err
		}

		if exists {
			return validation.NewError("validation_phone_exists", "phone number already exists")
		}
		return nil
	}
}
