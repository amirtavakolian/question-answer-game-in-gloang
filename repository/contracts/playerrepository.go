package contracts

import (
	"QA-Game/dto/playerdto"
	"QA-Game/entity"
)

type PlayerRepository interface {
	IsPhoneNumberExist(phoneNumber string) (bool, error)
	Store(playerDTO playerdto.PlayerRegister) (entity.Player, error)
}
