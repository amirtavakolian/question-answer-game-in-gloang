package contracts

import (
	"QA-Game/param/permissionparam"
	"QA-Game/repository/dbresponses"
)

type PermissionRepository interface {
	Store(permission permissionparam.StorePermissionParam) dbresponses.PermissionResponse
}
