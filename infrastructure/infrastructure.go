package infrastructure

import (
	"database/sql"
	"fmt"

	"api/api-hotel/domain/entities"
	"api/api-hotel/domain/usecase"
	"api/api-hotel/infrastructure/handler"
	"api/api-hotel/infrastructure/repository"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Config(env entities.Env, r *gin.Engine) {
	db := connection(env)
	handlerCfg(db, r)
}

func connection(env entities.Env) *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		env.DBHOST, env.DBPORT, env.DBUSER, env.DBPASSAWORD, env.DBNAME)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	repository.NewMysqlHotelRepository(db)

	return db

}

func handlerCfg(dbConn *sql.DB, r *gin.Engine) {
	hotelRepo := repository.NewMysqlHotelRepository(dbConn)

	uu := usecase.NewHotelUsecase(hotelRepo)

	handler.NewProjectHandler(r, uu)

}
