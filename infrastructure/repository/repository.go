package repository

import (
	"context"
	"database/sql"
	"log"
	"sync"

	"api/api-hotel/domain/entities"
	"api/api-hotel/interfaces"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlRepo struct {
	DB    *sql.DB
	mutex sync.Mutex
}

func NewMysqlHotelRepository(db *sql.DB) interfaces.HotelRepo {
	return &mysqlRepo{
		DB: db,
	}
}

func (m *mysqlRepo) ListUsers(ctx context.Context, userInfo entities.User) (users []entities.User, err error) {

	var param *string
	var query string

	if userInfo.Name != "" {
		param = &userInfo.Name
		query = "select * from public.user where name = $1"
	} else if userInfo.Document != "" {
		param = &userInfo.Document
		query = "select * from public.user where document = $1"
	} else if userInfo.Phone != "" {
		param = &userInfo.Phone
		query = "select * from public.user where phone = $1"
	} else {
		query = "select * from public.user"
	}

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Println("Err: ", err.Error())
		return
	}

	var rows *sql.Rows

	if param == nil {
		rows, err = stmt.QueryContext(ctx)
		if err != nil {
			log.Println("Err: ", err.Error())
			return
		}
	} else {
		rows, err = stmt.QueryContext(ctx, param)
		if err != nil {
			log.Println("Err: ", err.Error())
			return
		}
	}

	for rows.Next() {
		var user entities.User

		err = rows.Scan(&user.ID, &user.Name, &user.Document, &user.Phone)
		if err != nil {
			log.Println("Err: ", err.Error())
			return users, err
		}

		users = append(users, user)
	}

	return
}
func (m *mysqlRepo) DeleteUser(ctx context.Context, userID int) (err error) {
	query := "DELETE FROM public.user WHERE id = $1"

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Println("Err: ", err.Error())
		return
	}

	_, err = stmt.ExecContext(ctx, userID)
	if err != nil {
		log.Println("Err: ", err.Error())
		return
	}

	return nil
}

func (m *mysqlRepo) UpdateUser(ctx context.Context, user entities.User, userID int) (err error) {
	query := "UPDATE public.user SET name = $1, document = $2, phone = $3 WHERE id = $4"

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Println("Err: ", err.Error())
		return
	}

	_, err = stmt.ExecContext(ctx, user.Name, user.Document, user.Phone, userID)
	if err != nil {
		log.Println("Err: ", err.Error())
		return
	}

	return nil

}

func (m *mysqlRepo) RegisterUser(ctx context.Context, user entities.User) (err error) {
	query := `INSERT INTO public.user (name, document, phone) 
    VALUES ($1, $2, $3)`
	_, err = m.DB.Exec(query, user.Name, user.Document, user.Phone)
	if err != nil {
		return
	}

	return nil
}

func (m *mysqlRepo) IsRoomEmpty(ctx context.Context, roomID int) (isRoomEmpty bool, err error) {
	query := "SELECT * FROM public.room WHERE ID = $1 and AVAILABLE = true"

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Println("Err: ", err.Error())
		return
	}

	rows, err := stmt.QueryContext(ctx, roomID)
	if err != nil {
		log.Println("Err: ", err.Error())
		return
	}

	if rows.Next() {
		isRoomEmpty = true
		return
	}

	isRoomEmpty = false

	return
}

func (m *mysqlRepo) RegisterReservation(ctx context.Context, acommodation entities.Acommodation, price float64) (err error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	_, err = tx.ExecContext(ctx, `
	INSERT INTO public.acommodation (user_id, room_id, check_in, check_out, price) 
    VALUES ($1, $2, $3, $4, $5)
    `, acommodation.UserID, acommodation.RoomID, acommodation.CheckIn, acommodation.Checkout, price)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
        UPDATE public.room SET available = false WHERE id = $1
    `, acommodation.RoomID)
	if err != nil {
		return err
	}

	return
}

func (m *mysqlRepo) ListAcommodations(ctx context.Context) (acommodations []entities.Acommodation, err error) {

	query := "select * from public.acommodation"
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Println("Err: ", err.Error())
		return
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Println("Err: ", err.Error())
		return
	}

	for rows.Next() {
		var acommodation entities.Acommodation

		err = rows.Scan(&acommodation.ID, &acommodation.UserID, &acommodation.RoomID, &acommodation.CheckIn, &acommodation.Checkout, &acommodation.Price)
		if err != nil {
			log.Println("Err: ", err.Error())
			return
		}

		acommodations = append(acommodations, acommodation)
	}

	return
}
