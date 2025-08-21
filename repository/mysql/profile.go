package mysql

import (
	"QA-Game/repository/dbresponses"
	"database/sql"
	"errors"
)

type Profile struct {
	Connection *Mysql
}

type ProfileData struct {
	ID       int `json:"id"`
	PlayerID int `json:"player_id"`
}

func NewProfileRepo() *Profile {
	return &Profile{
		Connection: NewMysql(),
	}
}

func (p *Profile) GetPlayerProfile(phoneNumber string) dbresponses.ProfileResponse {

	profileData := ProfileData{}

	result := p.Connection.DB.QueryRow(`SELECT profiles.* FROM profiles JOIN players ON players.id = profiles.player_id WHERE players.phone_number = ?`, phoneNumber)

	scanResult := result.Scan(&profileData.ID, &profileData.PlayerID)

	if scanResult != nil && errors.Is(scanResult, sql.ErrNoRows) {
		return dbresponses.ProfileResponse{
			Status:  false,
			Message: "Record not found",
		}
	}

	return dbresponses.ProfileResponse{
		Status: true,
		Data:   profileData,
	}
}
