package auth

import (
	restful "github.com/emicklei/go-restful"
	authApi "github.com/wzt3309/k8sconsole/src/app/backend/auth/api"
	"github.com/wzt3309/k8sconsole/src/app/backend/errors"
)

type AuthHandler struct {
	manager authApi.FrontendAuthManager
}

func (self AuthHandler) Install(ws *restful.WebService) {
	ws.Route(
		ws.POST("/login").
			To(self.handleLogin).
			Reads(authApi.FrontendAuthPayload{}).
			Writes(authApi.FrontendAuthResponse{}))
}

func (self AuthHandler) handleLogin(request *restful.Request, response *restful.Response) {
	payload := new(authApi.FrontendAuthPayload)
	if err := request.ReadEntity(payload); err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(errors.HandleHTTPError(err), err.Error()+"\n")
		return
	}
}
