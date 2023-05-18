package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"api/api-hotel/domain/entities"
	"api/api-hotel/interfaces"
	errorUtils "api/api-hotel/utils/error"

	_ "api/api-hotel/docs"

	"github.com/gin-gonic/gin"
)

type HotelHandler struct {
	HotelUseCase interfaces.HotelUseCase
}

func NewProjectHandler(r *gin.Engine, us interfaces.HotelUseCase) *gin.Context {
	handler := &HotelHandler{
		HotelUseCase: us,
	}
	r.POST("/user", handler.registerUser)
	r.GET("/user", handler.listUsers)
	r.PUT("/user/:id", handler.updateUser)
	r.DELETE("/user/:id", handler.deleteUser)
	r.POST("/reservation", handler.registerReservation)
	r.GET("/reservation", handler.listReservation)

	return nil
}

// @Summary Register Reservation
// @Description Register a new reservation for a user in a room
// @Tags Reservation
// @Accept  json
// @Produce  json
// @Param   body     body    entities.Acommodation    true        "reservation info"
// @Success 201 {string} string	"Reservation created"
// @Failure 400 {object} entities.Error
// @Failure 500 {object} entities.Error
// @Router /reservation [post]
func (uh *HotelHandler) registerReservation(c *gin.Context) {
	ctx := context.Background()

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	acommodation := entities.Acommodation{}

	if err = json.Unmarshal(bytes, &acommodation); err != nil {
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	if acommodation.UserID == 0 || acommodation.RoomID == 0 {
		err = errors.New("missing userID or RoomID")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(400, errResp)
		return
	}

	err = uh.HotelUseCase.RegisterReservation(ctx, acommodation)
	if err != nil {
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	c.Status(http.StatusCreated)

}

// @Summary Delete User
// @Description Delete a user by ID
// @Tags Users
// @Param   id      path    int     true        "User ID"
// @Success 200 {string} string	"User deleted"
// @Failure 400 {object} entities.Error
// @Failure 500 {object} entities.Error
// @Router /users/{id} [delete]
func (uh *HotelHandler) deleteUser(c *gin.Context) {
	ctx := context.Background()

	userID := c.Param("id")
	if len(userID) == 0 {
		err := errors.New("invalid user id")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(400, errResp)
		return
	}

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	err = uh.HotelUseCase.DeleteUser(ctx, intUserID)
	if err != nil {
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	c.JSON(200, "User deleted")
}

// @Summary Update User
// @Description Update a user by ID
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   id      path    int     true        "User ID"
// @Param   body     body    entities.User    		true        "user info"
// @Success 200 {string} string	"User updated"
// @Failure 400 {object} entities.Error
// @Failure 500 {object} entities.Error
// @Router /users/{id} [put]
func (uh *HotelHandler) updateUser(c *gin.Context) {
	ctx := context.Background()

	userID := c.Param("id")
	if len(userID) == 0 {
		err := errors.New("invalid user id")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(400, errResp)
		return
	}

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	user := entities.User{}

	if err = json.Unmarshal(bytes, &user); err != nil {
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	if user.Name == "" || user.Document == "" || user.Phone == "" {
		err = errors.New("fields can not be empty")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(400, errResp)
		return
	}

	err = uh.HotelUseCase.UpdateUser(ctx, user, intUserID)
	if err != nil {
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	c.JSON(200, "User updated")
}

// @Summary List Users
// @Description List all users
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   name      query    string     false        "User Name"
// @Param   document  query    string     false        "User Document"
// @Param   phone     query    string     false        "User Phone"
// @Success 200 {array} entities.User
// @Failure 500 {object} entities.Error
// @Router /users [get]
func (uh *HotelHandler) listUsers(c *gin.Context) {
	ctx := context.Background()

	name := c.Query("name")
	document := c.Query("document")
	phone := c.Query("phone")

	user := entities.User{
		Name:     name,
		Document: document,
		Phone:    phone,
	}

	users, err := uh.HotelUseCase.ListUsers(ctx, user)
	if err != nil {
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	c.JSON(200, users)
}

// @Summary Register User
// @Description Register a new user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   body     body    entities.User    true        "user info"
// @Success 201 {string} string	"User created"
// @Failure 400 {object} entities.Error
// @Failure 500 {object} entities.Error
// @Router /users [post]
func (uh *HotelHandler) registerUser(c *gin.Context) {
	ctx := context.Background()

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	user := entities.User{}

	if err = json.Unmarshal(bytes, &user); err != nil {
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	if user.Name == "" || user.Document == "" || user.Phone == "" {
		err = errors.New("fields can not be empty")
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(400, errResp)
		return
	}

	err = uh.HotelUseCase.RegisterUser(ctx, user)
	if err != nil {
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary List all Reservation
// @Description Get a list of all Reservation
// @Tags Reservation
// @Accept json
// @Produce json
// @Success 200 {array} entities.Acommodation
// @Failure 500 {object} entities.Error
// @Router /reservation [get]
func (uh *HotelHandler) listReservation(c *gin.Context) {
	ctx := context.Background()

	acommodations, err := uh.HotelUseCase.ListAcommodations(ctx)
	if err != nil {
		errResp := errorUtils.CreateError(err.Error())
		c.JSON(500, errResp)
		return
	}

	c.JSON(200, acommodations)
}
