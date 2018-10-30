package user

import (
	"github.com/emicklei/go-restful"
	"github.com/wzt3309/k8sconsole/src/app/backend/api"
	"github.com/wzt3309/k8sconsole/src/app/backend/errors"
	userApi "github.com/wzt3309/k8sconsole/src/app/backend/user/api"
	"net/http"
)

type UserHandler struct {
	manager userApi.UserManager
}

func (self UserHandler) Install(ws *restful.WebService) {
	ws.Route(
		ws.POST("/users").
			To(self.userCreate).
			Reads(userApi.UserCreatePayload{}).
			Writes(api.User{}))
	ws.Route(
		ws.GET("/users").
			To(self.userList).
			Writes([]api.User{}))
}

func (self UserHandler) userCreate(r *restful.Request, w *restful.Response) {
	payload := new(userApi.UserCreatePayload)
	if err := r.ReadEntity(payload); err != nil {
		w.AddHeader("Content-Type", "text/plain")
		w.WriteErrorString(errors.HandleHTTPError(err), err.Error()+"\n")
		return
	}

	user, err := self.manager.UserCreate(payload)
	if err != nil {
		w.AddHeader("Content-Type", "text/plain")
		w.WriteErrorString(errors.HandleHTTPError(err), err.Error()+"\n")
		return
	}

	w.WriteHeaderAndEntity(http.StatusOK, user)
}

func (self UserHandler) userList(r *restful.Request, w *restful.Response) {
	users, err := self.manager.UserList()
	if err != nil {
		w.AddHeader("Content-Type", "text/plain")
		w.WriteErrorString(errors.HandleHTTPError(err), err.Error()+"\n")
		return
	}

	w.WriteHeaderAndEntity(http.StatusOK, users)
}

func NewUserHandler(manager userApi.UserManager) UserHandler {
	return UserHandler{manager: manager}
}
