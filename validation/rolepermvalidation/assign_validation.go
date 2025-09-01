package rolepermvalidation

import (
	"QA-Game/param/rolepermissionparam"
	"QA-Game/repository/mysql"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"strings"
)

type AssignValidation struct {
}

func NewAssignValidation() AssignValidation {
	return AssignValidation{}
}

func (assignValidation AssignValidation) Validate(assignPermToRoleParams rolepermissionparam.AssignPermissionToRoleParam) (bool, map[string]interface{}) {

	err := validation.ValidateStruct(&assignPermToRoleParams,
		validation.Field(&assignPermToRoleParams.Role_id,
			validation.Required,
			validation.By(RoleIdMustExistInDatabase(mysql.NewPermissionRepo())),
		),

		validation.Field(&assignPermToRoleParams.Permission_id,
			validation.Required,
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

func RoleIdMustExistInDatabase(db *mysql.Permission) validation.RuleFunc {
	return func(value interface{}) error {
		roleId, ok := value.(int)
		if !ok {
			return validation.NewError("validation_invalid_type", "role id must be integer")
		}

		var exists bool

		err := db.Connection.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM roles WHERE id = ?)", roleId).Scan(&exists)

		if err != nil {
			return err
		}

		if !exists {
			return validation.NewError("validation_name_exists", "role id is not exist")
		}
		return nil
	}
}
