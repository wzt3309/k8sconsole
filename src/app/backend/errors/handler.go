package errors

import (
	"encoding/json"
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"net/http"
)

func (routeFunc LoggerHandler) handle(r *restful.Request, w *restful.Response) {
	err := routeFunc(r, w)
	if err != nil {
		writeErrorResponse(w, err)
	}
}

func writeErrorResponse(w *restful.Response, err *HandlerError) {
	glog.Errorf("http error: %s (err=%s) (code=%d)\n", err.Message, err.Err, err.StatusCode)
	w.AddHeader("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	json.NewEncoder(w).Encode(&errorResponse{Err: err.Message, Details: err.Err.Error()})
}

func HandleHTTPError(err error) int {
	return http.StatusInternalServerError
}
