package mysql

import (
	"QA-Game/dto/playerdto"
	"QA-Game/entity"
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

func (p *Player) Store(playerDTO playerdto.PlayerRegister) (entity.Player, error) {

	playerEntity := entity.Player{}

	result, err := p.Connection.DB.Exec("INSERT INTO players (name, phone_number, password, avatar) VALUES (?, ?, ?, ?)", playerDTO.Name, playerDTO.PhoneNumber, playerDTO.Password, playerDTO.Avatar)

	if err != nil {
		return playerEntity, err
	}

	if id, err := result.LastInsertId(); err == nil {
		playerEntity.Id = uint(id)
		playerEntity.Name = playerDTO.Name
		playerEntity.PhoneNumber = playerDTO.PhoneNumber
		playerEntity.Password = playerDTO.Password
		playerEntity.Avatar = playerDTO.Avatar
	}

	return playerEntity, nil
}

func (p *Player) FindPlayerByPhoneNumber(phoneNumber string) (string, string, error) {

	result := p.Connection.DB.QueryRow("SELECT phone_number, password FROM players WHERE phone_number = ?", phoneNumber)

	var phone_number, password string

	scanResult := result.Scan(&phone_number, &password)

	if scanResult != nil {
		return "", "", fmt.Errorf("Player not found.")
	}

	return phone_number, password, nil
}
