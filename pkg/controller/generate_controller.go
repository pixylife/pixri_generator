package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"pixri_generator/pkg/model"
)

type GenRequest struct {
	Application model.Application
	Entity []GenEntity
	Theme model.Theme
}


type GenEntity struct {
	Entity model.Entity
	Fields []*model.Field
}

func GenerateApplication(c echo.Context) error {

	fmt.Println(c.Request())

	generateRequest := GenRequest{}
	if error := c.Bind(&generateRequest); error != nil {
		return error
	}
	return c.JSON(http.StatusOK, generateRequest)
}

func GenerateController(g *echo.Group, contextRoot string) {
	g.POST(contextRoot+"/generate", GenerateApplication)
}
