package helpers

import (
	"go-api-basic-auth/src/config"
	"go-api-basic-auth/src/model"

	"github.com/labstack/echo/v4"
)

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	student := new(model.Student)
	if err := db.Where("email = ?", username).First(student).Error; err != nil {
		return false, nil
	}

	if student.Password != password {
		return false, nil
	}

	return true, nil
}