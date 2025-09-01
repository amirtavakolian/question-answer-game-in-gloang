package contracts

import (
	"QA-Game/param/rolepermissionparam"
	"QA-Game/repository/dbresponses"
)

type RoleRepository interface {
	Store(role rolepermissionparam.StoreRoleParam) dbresponses.RoleResponse
}
