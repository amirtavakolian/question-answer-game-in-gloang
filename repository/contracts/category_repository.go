package contracts

import (
	"QA-Game/param/categoryparam"
	"QA-Game/repository/dbresponses"
)

type CategoryRepository interface {
	Store(categoryParam categoryparam.CategoryStore) dbresponses.CategoryResponse
}
