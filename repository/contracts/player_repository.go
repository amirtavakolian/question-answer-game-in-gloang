package contracts

import (

"QA-Game/entity"
"QA-Game/param/playerparam"
"QA-Game/repository/dbresponses"

)

type PlayerRepository interface {
	IsPhoneNumberExist(phoneNumber string) (bool, error)
	Store(playerDTO playerparam.PlayerRegisterRequest) (entity.Player, error)
	FindPlayerByPhoneNumber(phoneNumber string) (dbresponses.Player, error)
}
