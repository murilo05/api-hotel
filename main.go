package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gitlab.engdb.com.br/apigin/domain/entities"
	"gitlab.engdb.com.br/apigin/infrastructure"
	"gitlab.engdb.com.br/apigin/utils/env"
)

func main() {
	r := gin.Default()

	env.ReadEnvFile()

	HOST := viper.GetString("HOST")
	PORT := viper.GetString("PORT")
	DBHOST := viper.GetString("DBHOST")
	DBPORT := viper.GetString("DBPORT")
	DBUSER := viper.GetString("DBUSER")
	DBPASSAWORD := viper.GetString("DBPASSAWORD")
	DBNAME := viper.GetString("DBNAME")
	TIMEOUT := viper.GetInt("TIMEOUT")

	env := entities.Env{
		HOST:        HOST,
		PORT:        PORT,
		DBHOST:      DBHOST,
		DBPORT:      DBPORT,
		DBUSER:      DBUSER,
		DBPASSAWORD: DBPASSAWORD,
		DBNAME:      DBNAME,
		TIMEOUT:     TIMEOUT,
	}

	infrastructure.Config(env, r)

	r.Run(HOST + ":" + PORT)
}
