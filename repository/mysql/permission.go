package mysql

import (
		"QA-Game/param/rolepermissionparam"
"QA-Game/repository/dbresponses"
)

type Permission struct {
	Connection *Mysql
}

func NewPermissionRepo() *Permission {
	return &Permission{
		Connection: NewMysql(),
	}
}

func (permission Permission) Store(permissionParam rolepermissionparam.StorePermissionParam) dbresponses.PermissionResponse {

	_, err := permission.Connection.DB.Exec("INSERT INTO permissions (name, description) values (?, ?)", permissionParam.Name, permissionParam.Description)

	if err != nil {
		response := dbresponses.PermissionResponse{
			Status:  false,
			Message: err.Error(),
		}
		return response
	}

	return dbresponses.PermissionResponse{
		Status:  true,
		Message: "Permission created successfully",
	}
}
