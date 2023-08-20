package config

import (
	"log"
	"os"

	"github.com/subosito/gotenv"
)

type AppCon struct {
	Srv *srvCon
	Db  *dbCon
	Token *token
}
type srvCon struct {
	Port string
}
type token struct{
	JwtSecret string
}

type dbCon struct {
	DbHost string
	DbUser string
	DbPass string
	DbName string
	DbPort string
}

func LoadConfig() *AppCon {
	if err := gotenv.Load(".env"); err != nil {
		log.Fatal("failed load env")
	}
	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT")
	srvPort := os.Getenv("PORT")
	jwtSecret := os.Getenv("JWT_SECRET")

	DbEnv := &dbCon{
		DbHost: dbHost,
		DbUser: dbUser,
		DbPass: dbPass,
		DbName: dbName,
		DbPort: dbPort,
	}
	SrvEnv := &srvCon{
		Port: srvPort,
	}
	token := &token{
		JwtSecret: jwtSecret,
	}
	conf := AppCon{
		Db:  DbEnv,
		Srv: SrvEnv,
		Token: token,
	}
	return &conf
}
