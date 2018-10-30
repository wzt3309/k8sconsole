package handler

import (
	"github.com/emicklei/go-restful"
	"github.com/wzt3309/k8sconsole/src/app/backend/auth"
	authApi "github.com/wzt3309/k8sconsole/src/app/backend/auth/api"
	"github.com/wzt3309/k8sconsole/src/app/backend/user"
	userApi "github.com/wzt3309/k8sconsole/src/app/backend/user/api"
	"net/http"
)

func CreateHTTPAPIHandler(fAuthManager authApi.FrontendAuthManager,
	userManager userApi.UserManager) (http.Handler, error) {

	wsContainer := restful.NewContainer()
	wsContainer.EnableContentEncoding(true)

	apiV1Ws := new(restful.WebService)

	apiV1Ws.Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	wsContainer.Add(apiV1Ws)

	fAuthHandler := auth.NewFrontendAuthHandler(fAuthManager)
	fAuthHandler.Install(apiV1Ws)

	userHandler := user.NewUserHandler(userManager)
	userHandler.Install(apiV1Ws)

	return wsContainer, nil
}
