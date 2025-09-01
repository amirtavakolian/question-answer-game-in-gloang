package mysql

import (
	"QA-Game/param/rolepermissionparam"
	"QA-Game/repository/dbresponses"
	"strings"
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

func (permission Permission) AssignPermToRole(assignPermToRoleParams rolepermissionparam.AssignPermissionToRoleParam) dbresponses.PermissionResponse {

	query := "INSERT INTO permission_role (role_id, permission_id) VALUES "

	values := []interface{}{}

	placeholders := []string{}

	for _, permissionID := range assignPermToRoleParams.Permission_id {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, assignPermToRoleParams.Role_id, permissionID)
	}

	query += strings.Join(placeholders, ", ")

	_, err := permission.Connection.DB.Exec("DELETE FROM permission_role WHERE role_id = ?", assignPermToRoleParams.Role_id)

	if err != nil {
		return dbresponses.PermissionResponse{
			Status:  false,
			Message: err.Error(),
		}
	}

	_, err = permission.Connection.DB.Exec(query, values...)

	if err != nil {
		return dbresponses.PermissionResponse{
			Status:  false,
			Message: err.Error(),
		}
	}

	return dbresponses.PermissionResponse{
		Status:  true,
		Message: "Permissions assigned successfully",
	}

}
