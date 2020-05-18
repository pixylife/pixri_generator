package generator

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pixri_generator/pkg/model"
)


func GenerateApp(c echo.Context) error {

	generateRequest := model.GenRequest{}
	if error := c.Bind(&generateRequest); error != nil {
		return error
	}

	go GenerateApplication(generateRequest)

	return c.JSON(http.StatusOK, "Request Submitted")
}

func GenerateController(g *echo.Group, contextRoot string) {
	g.POST(contextRoot+"/generate", GenerateApp)
}
