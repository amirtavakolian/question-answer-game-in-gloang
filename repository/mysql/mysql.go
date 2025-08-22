package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Mysql struct {
	DB *sql.DB
}

func NewMysql() *Mysql {

	// todo => remove user pass (prevent from hard-coding)
	db, err := sql.Open("mysql", "root:@(localhost:3306)/question-game?parseTime=true")

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &Mysql{DB: db}
}
