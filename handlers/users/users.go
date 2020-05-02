package users

import (
	"fmt"
	"github.com/coaraujo/users-go-api/infrastructure/config"
	"github.com/coaraujo/users-go-api/infrastructure/queue"
	userService "github.com/coaraujo/users-go-api/services/user"
	"github.com/coaraujo/users-go-api/domains"
	"github.com/labstack/echo"
	"net/http"
	"bytes"
	"encoding/json"
)

func GetUser(c echo.Context) error {
	id := c.Param("id")

	user, err := userService.GetInstance().Get(id)
	if err != nil {
		_ = c.NoContent(http.StatusNotFound)
		return err
	}

	_ = c.JSON(http.StatusOK, &user)
	return nil
}

func CreateUser(c echo.Context) error {
	user := new(domains.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(user)
	message := reqBodyBytes.Bytes()

	err := queue.GetInstance().SendMessage(config.UserCreateTopic, message)
	fmt.Println(err)

	_ = c.JSON(http.StatusOK, user)
	return nil
}