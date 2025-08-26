package mysql

import (
	"QA-Game/entity"
	"QA-Game/param/playerparam"
	"QA-Game/repository/dbresponses"
	"database/sql"
	"fmt"
)

type Player struct {
	Connection *Mysql
}

func NewPlayerRepo() *Player {
	return &Player{
		Connection: NewMysql(),
	}
}

func (p *Player) IsPhoneNumberExist(phoneNumber string) (bool, error) {

	row := p.Connection.DB.QueryRow("SELECT phone_number FROM players WHERE phone_number = ?", phoneNumber)

	var mobile string
	err1 := row.Scan(&mobile)

	if err1 != nil && err1 == sql.ErrNoRows {
		return true, err1
		//todo => extend the error response struct to hold & log system errors for developers, then remove them from users
	}

	return false, fmt.Errorf("Phone number is available")
}

func (p *Player) Store(playerDTO playerparam.PlayerRegisterRequest) (entity.Player, error) {

	playerEntity := entity.Player{}

	result, err := p.Connection.DB.Exec("INSERT INTO players (name, phone_number, password, avatar) VALUES (?, ?, ?, ?)", playerDTO.Name, playerDTO.PhoneNumber, playerDTO.Password, playerDTO.Avatar)

	if err != nil {
		return playerEntity, err
	}

	if id, err := result.LastInsertId(); err == nil {
		playerEntity.Id = uint(id)
		playerEntity.Name = playerDTO.Name
		playerEntity.PhoneNumber = playerDTO.PhoneNumber
		playerEntity.Avatar = playerDTO.Avatar
	}

	return playerEntity, nil
}

func (p *Player) FindPlayerByPhoneNumber(phoneNumber string) (dbresponses.Player, error) {

	result := p.Connection.DB.QueryRow("SELECT id, phone_number, password FROM players WHERE phone_number = ?", phoneNumber)

	var response dbresponses.Player

	scanResult := result.Scan(&response.UserId, &response.PhoneNumber, &response.Password)

	if scanResult != nil {
		return response, fmt.Errorf("Player not found.")
	}

	return response, nil
}
