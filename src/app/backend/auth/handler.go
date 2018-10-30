package auth

import (
	restful "github.com/emicklei/go-restful"
	authApi "github.com/wzt3309/k8sconsole/src/app/backend/auth/api"
	"github.com/wzt3309/k8sconsole/src/app/backend/errors"
	"net/http"
)

type FrontendAuthHandler struct {
	manager authApi.FrontendAuthManager
}

func (self FrontendAuthHandler) Install(ws *restful.WebService) {
	ws.Route(
		ws.POST("/login").
			To(self.handleLogin).
			Reads(authApi.FrontendAuthPayload{}).
			Writes(authApi.FrontendAuthResponse{}))
}

func (self FrontendAuthHandler) handleLogin(r *restful.Request, w *restful.Response) {
	payload := new(authApi.FrontendAuthPayload)
	if err := r.ReadEntity(payload); err != nil {
		w.AddHeader("Content-Type", "text/plain")
		w.WriteErrorString(errors.HandleHTTPError(err), err.Error()+"\n")
		return
	}
	resp, err := self.manager.Login(payload)

	if err != nil {
		w.AddHeader("Content-Type", "text/plain")
		w.WriteErrorString(errors.HandleHTTPError(err), err.Error()+"\n")
		return
	}

	w.WriteHeaderAndEntity(http.StatusOK, resp)
}

func NewFrontendAuthHandler(manager authApi.FrontendAuthManager) FrontendAuthHandler {
	return FrontendAuthHandler{manager: manager}
}
