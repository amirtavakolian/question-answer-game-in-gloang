package mysql

import (
	"QA-Game/param/categoryparam"
	"QA-Game/repository/dbresponses"
	)

type Category struct {
	Connection *Mysql
}

func NewCategoryRepo() *Category {
	return &Category{
		Connection: NewMysql(),
	}
}

func (c Category) Store(categoryParam categoryparam.CategoryStore) dbresponses.CategoryResponse {

	_, err := c.Connection.DB.Exec("INSERT INTO categories (title) VALUES (?)", categoryParam.Title)

	if err != nil {
		response := dbresponses.CategoryResponse{
			Status:  false,
			Message: err.Error(),
		}

		return response
	}


	return dbresponses.CategoryResponse{
		Status:  true,
		Message: "Category created successfully",
	}
}
