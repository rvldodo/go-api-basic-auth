package router

import (
	"go-api-basic-auth/src/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StudentRouter(c *echo.Echo) {

	c.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${host}${path} ${status}\n",
	}))

	studentRouter := c.Group("/api/v1")

	studentRouter.GET("/student", controller.GetAllStudents)
	studentRouter.GET("/student/:id", controller.GetStudentById)
	studentRouter.POST("/student", controller.CreateStudent)
	studentRouter.PUT("/student/:id", controller.UpdateStudent)
	studentRouter.DELETE("/student/:id", controller.DeleteStudent)
}