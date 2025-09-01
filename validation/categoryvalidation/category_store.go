package categoryvalidation

import (
	"QA-Game/param/categoryparam"
	"QA-Game/repository/mysql"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"strings"
)

type CategoryStore struct {
}

func (s CategoryStore) Validate(categoryStore categoryparam.CategoryStore) (bool, map[string]interface{}) {

	err := validation.ValidateStruct(&categoryStore,
		validation.Field(&categoryStore.Title,
			validation.Required,
			validation.Length(3, 250),
			validation.By(isCategoryTitleExist(mysql.NewPermissionRepo())),
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

func isCategoryTitleExist(db *mysql.Permission) validation.RuleFunc {
	return func(value interface{}) error {
		categoryTitle, ok := value.(string)
		if !ok {
			return validation.NewError("validation_invalid_type", "category must be string")
		}

		var exists bool

		err := db.Connection.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM categories WHERE title = ?)", categoryTitle).Scan(&exists)

		if err != nil {
			return err
		}

		if exists {
			return validation.NewError("validation_name_exists", "title is exist")
		}

		return nil
	}
}
