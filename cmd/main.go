package main

import (
	"api/api-hotel/domain/entities"
	"api/api-hotel/infrastructure"
	"api/api-hotel/utils/env"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := gin.Default()

	env.ReadEnvFile()

	HOST := viper.GetString("HOST")
	PORT := viper.GetString("PORT")
	DBHOST := viper.GetString("DBHOST")
	DBPORT := viper.GetInt("DBPORT")
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
