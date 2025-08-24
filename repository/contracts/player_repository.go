package contracts

import (
	"QA-Game/param/playerparam"
	"QA-Game/entity"
)

type PlayerRepository interface {
	IsPhoneNumberExist(phoneNumber string) (bool, error)
	Store(playerDTO playerparam.PlayerRegisterRequest) (entity.Player, error)
	FindPlayerByPhoneNumber(phoneNumber string) (string, string, error)
}
