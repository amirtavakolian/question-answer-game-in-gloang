package contracts

import "QA-Game/repository/dbresponses"

type ProfileRepository interface {
	GetPlayerProfile(phoneNumber string) dbresponses.ProfileResponse
}
