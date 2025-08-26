package contracts

import (
		"QA-Game/param/rolepermissionparam"
"QA-Game/repository/dbresponses"
)

type PermissionRepository interface {
	Store(permission rolepermissionparam.StorePermissionParam) dbresponses.PermissionResponse
	AssignPermToRole(assignPermToRoleParams rolepermissionparam.AssignPermissionToRoleParam) dbresponses.PermissionResponse
}
