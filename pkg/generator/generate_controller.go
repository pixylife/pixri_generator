package generator

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"net/http"
	"pixri_generator/functions"
	"pixri_generator/pkg/controller"
	"pixri_generator/pkg/generator/app"
	"pixri_generator/pkg/model"
)


func GenerateApp(c echo.Context) error {

	generateRequest := model.GenRequest{}
	if error := c.Bind(&generateRequest); error != nil {
		return error
	}

	appName := functions.SpaceStringsBuilder(generateRequest.Application.Name)
	repo,_,_ :=controller.CreateRepository(appName)

	guid := xid.New()

	app.ProjectResponse.GithubURL = *repo.CloneURL
	app.ProjectResponse.WSURL = "/ws/"+guid.String()
	app.ProjectResponse.ClientRequestId = guid.String()


	go GenerateApplication(generateRequest,repo)



	return c.JSON(http.StatusOK, app.ProjectResponse)
}

func GenerateController(g *echo.Group, contextRoot string) {
	g.POST(contextRoot+"/generate", GenerateApp)
}
