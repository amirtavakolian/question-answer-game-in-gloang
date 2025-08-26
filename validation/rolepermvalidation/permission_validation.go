package rolepermvalidation

import (
	"QA-Game/param/permissionparam"
	"QA-Game/repository/mysql"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"strings"
)

type PermissionValidation struct {
}

func NewPermissionValidation() PermissionValidation {
	return PermissionValidation{}
}

func (p PermissionValidation) Validate(permissionParam permissionparam.StorePermissionParam) (bool, map[string]interface{}) {

	err := validation.ValidateStruct(&permissionParam,
		validation.Field(&permissionParam.Name,
			validation.Required,
			validation.Length(2, 50),
			validation.By(UniquePermissionNameRule(mysql.NewPermissionRepo())),
		),

		validation.Field(&permissionParam.Description,
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

func UniquePermissionNameRule(db *mysql.Permission) validation.RuleFunc {
	return func(value interface{}) error {
		permissionName, ok := value.(string)
		if !ok {
			return validation.NewError("validation_invalid_type", "permission name must be a string")
		}

		var exists bool

		err := db.Connection.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM permissions WHERE name = ?)", permissionName).Scan(&exists)

		if err != nil {
			return err
		}

		if exists {
			return validation.NewError("validation_name_exists", "permission name already exists")
		}
		return nil
	}
}
