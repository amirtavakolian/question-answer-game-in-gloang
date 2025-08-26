package rolepermvalidation

import (
	"QA-Game/param/roleparam"
	"QA-Game/repository/mysql"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"strings"
)

type RoleValidation struct {
}

func New() RoleValidation {
	return RoleValidation{}
}

func (r RoleValidation) Validate(roleParam roleparam.StoreRoleParam) (bool, map[string]interface{}) {

	err := validation.ValidateStruct(&roleParam,
		validation.Field(&roleParam.Name,
			validation.Required,
			validation.Length(2, 50),
			validation.By(UniqueNameRule(mysql.NewPlayerRepo())),
		),

		validation.Field(&roleParam.Description,
			validation.Length(3, 200),
			validation.NilOrNotEmpty,
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

func UniqueNameRule(db *mysql.Player) validation.RuleFunc {
	return func(value interface{}) error {
		roleName, ok := value.(string)
		if !ok {
			return validation.NewError("validation_invalid_type", "role name must be a string")
		}

		var exists bool

		err := db.Connection.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM roles WHERE name = ?)", roleName).Scan(&exists)

		if err != nil {
			return err
		}

		if exists {
			return validation.NewError("validation_name_exists", "role name already exists")
		}
		return nil
	}
}
