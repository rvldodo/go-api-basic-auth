package controller

import (
	"go-api-basic-auth/src/config"
	"go-api-basic-auth/src/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllStudents(c echo.Context) error {
	var students []model.Student

	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = db.Find(&students).Error
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, students)
}

func GetStudentById(c echo.Context) error {

	var student model.Student

	id := c.Param("id")

	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	if err = db.First(&student, id).Error; err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, student)
}

func CreateStudent(c echo.Context) error {
	s := new(model.Student)

	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}
	
	if err := c.Bind(s); err != nil {
		panic(err)
	}

	student := &model.Student{
		FirstName: s.FirstName,
		LastName: s.LastName,
		Age: s.Age,
		IsPass: s.IsPass,
	}

	if err := db.Create(&student).Error; err != nil {
		panic(err)
	}

	response := map[string]interface{}{
		"Student": s,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateStudent(c echo.Context) error {
	return c.JSON(http.StatusOK, "Success update student")
}

func DeleteStudent(c echo.Context) error {
	return c.JSON(http.StatusOK, "Success deleted student")
}