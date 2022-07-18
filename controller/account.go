package controller

import (
	"net/http"
	"strconv"

	dao "github.com/adsl99801/zoo/dao"
	model "github.com/adsl99801/zoo/model"
	"github.com/labstack/echo/v4"
)

var (
	users = map[int]*model.User{}
	seq   = 1
)

//----------
// Handlers
//----------

func CreateUser(c echo.Context) error {
	u := &model.User{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	account,_:= dao.GetAccount(c.Request().Context(),int32(id))
	return c.JSON(http.StatusOK, account)
	// return c.JSON(http.StatusOK, users[id])
}

func UpdateUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func GetAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Home")
}
func Login(c echo.Context) error {
	return c.String(http.StatusOK, "Login")
}
