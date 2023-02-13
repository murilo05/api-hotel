package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gitlab.engdb.com.br/apigin/domain/entities"
	"gitlab.engdb.com.br/apigin/domain/usecase"
	"gitlab.engdb.com.br/apigin/infrastructure/handler"
	"gitlab.engdb.com.br/apigin/infrastructure/repository"
)

func Config(env entities.Env, r *gin.Engine) {
	db := connection(env)
	handlerCfg(db, r)
}

func connection(env entities.Env) *sql.DB {

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", env.DBUSER, env.DBPASSAWORD, env.DBHOST, env.DBPORT, env.DBNAME)
	dbConn, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatalf("Error while connecting with database %s", err)
	}

	repository.NewMysqlScheduleRepository(dbConn)

	return dbConn

}

func handlerCfg(dbConn *sql.DB, r *gin.Engine) {
	scheduleRepo := repository.NewMysqlScheduleRepository(dbConn)

	uu := usecase.NewScheduleUsecase(scheduleRepo)

	handler.NewProjectHandler(r, uu)

}
