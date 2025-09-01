package mysql

import (
	"QA-Game/param/rolepermissionparam"
	"QA-Game/repository/dbresponses"
)

type Role struct {
	Connection *Mysql
}

func NewRoleRepo() Role {
	return Role{
		Connection: NewMysql(),
	}
}

func (role Role) Store(roleParam rolepermissionparam.StoreRoleParam) dbresponses.RoleResponse {

	_, err := role.Connection.DB.Exec("INSERT INTO roles (name, description) values (?, ?)", roleParam.Name, roleParam.Description)

	if err != nil {
		response := dbresponses.RoleResponse{
			Status:  false,
			Message: err.Error(),
		}
		return response
	}

	return dbresponses.RoleResponse{
		Status:  true,
		Message: "Role created successfully",
	}
}
