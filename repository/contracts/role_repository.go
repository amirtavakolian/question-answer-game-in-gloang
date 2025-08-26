package contracts

import (
	"QA-Game/param/roleparam"
	"QA-Game/repository/dbresponses"
)

type RoleRepository interface {
	Store(role roleparam.StoreRoleParam) dbresponses.RoleResponse
}
