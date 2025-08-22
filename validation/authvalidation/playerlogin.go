package authvalidation

import (
	"QA-Game/dto/playerdto"
	"github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
	"strings"
)

type PlayerLogin struct{}

func (r PlayerLogin) Validate(loginPlayerDto playerdto.PlayerLogin) (bool, map[string]interface{}) {

	err := validation.ValidateStruct(&loginPlayerDto,

		validation.Field(&loginPlayerDto.PhoneNumber,
			validation.Required,
			validation.Length(6, 50),
			validation.Match(regexp.MustCompile("^09\\d{9}$")).Error("phone number format is not correct"),
		),

		validation.Field(&loginPlayerDto.Password,
			validation.Required,
			validation.Length(6, 200),
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
