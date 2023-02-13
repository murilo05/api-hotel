package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.engdb.com.br/apigin/domain/entities"
	"gitlab.engdb.com.br/apigin/interfaces"
)

type mysqlRepo struct {
	DB *sql.DB
}

// NewMysqlAuthorRepository will create an implementation of author.Repository
func NewMysqlScheduleRepository(db *sql.DB) interfaces.ScheduleRepo {
	return &mysqlRepo{
		DB: db,
	}
}

func (m *mysqlRepo) GetSchedules(ctx context.Context) ([]entities.ResponseGetSchudeles, int, error) {

	query := "SELECT * FROM agendamentos"

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, 500, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, 500, err
	}

	schedules := []entities.ResponseGetSchudeles{}

	for rows.Next() {
		var schedule entities.ResponseGetSchudeles

		err = rows.Scan(&schedule.Company.CNPJ, &schedule.Hour, &schedule.FinalHour)
		if err != nil {
			fmt.Println("Err", err.Error())
			return nil, 500, err
		}

		resp, err := http.Get("https://receitaws.com.br/v1/cnpj/" + schedule.Company.CNPJ)
		if err != nil {
			schedule.AWS = "Não foi possivel consultar a razão social"
		} else {
			bodyAWS, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				schedule.AWS = "Não foi possivel consultar a razão social"
			} else {
				json.Unmarshal(bodyAWS, &schedule.AWS)

			}
		}

		schedules = append(schedules, schedule)
	}
	if len(schedules) == 0 {
		return schedules, 204, nil
	}

	return schedules, 200, nil
}
func (m *mysqlRepo) CreateSchedule(ctx context.Context, body entities.InputSchedule) (int, error) {

	finalHour, err := time.Parse("15:04", body.Hour)
	if err != nil {
		return 500, err
	}

	finalHour.Hour()

	finalHour = finalHour.Add(1 * time.Hour)

	finalHourstr := finalHour.Format("15:04")

	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return 500, err
	}

	query := "INSERT INTO agendamentos (cnpj,horario_inicial, horario_final) VALUES (?,?,?)"

	_, err = tx.ExecContext(ctx, query, body.Company.CNPJ, body.Hour, finalHourstr)
	if err != nil {
		fmt.Println("err", err.Error())
		return 500, err
	}

	err = m.updateAvailability(ctx, body, tx)
	if err != nil {
		fmt.Println("err", err.Error())
		return 500, err
	}

	tx.Commit()

	return 200, nil
}

func (m *mysqlRepo) updateAvailability(ctx context.Context, body entities.InputSchedule, tx *sql.Tx) error {
	query := "UPDATE horarios SET disponibilidade = 0 WHERE horarios_iniciais = ?"

	_, err := tx.ExecContext(ctx, query, body.Hour)
	if err != nil {
		fmt.Println("err", err.Error())
		return err
	}

	return nil

}

func (m *mysqlRepo) GetAllHours(ctx context.Context) ([]entities.ResponseAvailability, int, error) {
	query := "SELECT horarios_iniciais, horarios_finais, disponibilidade FROM horarios "

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println("aqui2", err.Error())
		return nil, 500, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		fmt.Println("aqui", err.Error())
		return nil, 500, err
	}

	schedules := []entities.ResponseAvailability{}

	for rows.Next() {
		var schedule entities.ResponseAvailability

		err = rows.Scan(&schedule.StartingHour, &schedule.FinalHour, &schedule.AvailableFromSQL)
		if err != nil {
			fmt.Println("Erro aqui", err.Error())
			return nil, 500, err
		}

		schedules = append(schedules, schedule)
	}
	if len(schedules) == 0 {
		return schedules, 204, nil
	}

	return schedules, 200, nil

}

func (m *mysqlRepo) IsHourAvailabe(ctx context.Context, body entities.InputSchedule) (bool, int, error) {
	query := "SELECT horarios_iniciais, horarios_finais, disponibilidade FROM horarios WHERE horarios_iniciais = ? AND disponibilidade = 1"

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println("aqui2", err.Error())
		return false, 500, err
	}

	rows, err := stmt.QueryContext(ctx, body.Hour)
	if err != nil {
		fmt.Println("aqui", err.Error())
		return false, 500, err
	}

	if rows.Next() {
		var schedule entities.ResponseAvailability

		err = rows.Scan(&schedule.StartingHour, &schedule.FinalHour, &schedule.AvailableFromSQL)
		if err != nil {
			fmt.Println("Erro aqui", err.Error())
			return false, 500, err
		}

		return true, 200, nil

	}

	return false, 400, nil
}
