package infra

import (
	"fmt"
	"log"
	"thresher/utils"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type DbConnector struct {
	Conn *gorm.DB
}

func NewPostgresConnector() *DbConnector {
	conf := utils.LoadConfig()
	dsn := dbConnInfo(*conf)
	Psql, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &DbConnector{
		Conn: Psql,
	}
}
func dbConnInfo(con utils.AppCon) string {
	databaseSourceName := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		con.Db.DbHost,
		con.Db.DbUser,
		con.Db.DbPass,
		con.Db.DbName,
		con.Db.DbPort,
	)
	return databaseSourceName
}
