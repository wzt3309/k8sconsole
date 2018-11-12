package handler

import (
	"github.com/emicklei/go-restful"
	"github.com/wzt3309/k8sconsole/src/app/backend/api"
	"github.com/wzt3309/k8sconsole/src/app/backend/auth"
	authApi "github.com/wzt3309/k8sconsole/src/app/backend/auth/api"
	clientApi "github.com/wzt3309/k8sconsole/src/app/backend/client/api"
	"github.com/wzt3309/k8sconsole/src/app/backend/resource/common"
	"github.com/wzt3309/k8sconsole/src/app/backend/resource/dataselect"
	"golang.org/x/net/xsrftoken"
	"net/http"
	"strconv"
	"strings"
)

type APIHandler struct {
	cManager clientApi.ClientManager
}

// CreateHTTPAPIHandler creates a new HTTP handler that handles all requests to the API of the backend.
func CreateHTTPAPIHandler(cManager clientApi.ClientManager, authManager authApi.AuthManager) (http.Handler, error) {
	apiHandler := APIHandler{cManager: cManager}

	wsContainer := restful.NewContainer()
	wsContainer.EnableContentEncoding(true)

	apiV1Ws := new(restful.WebService)
	apiV1Ws.Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	wsContainer.Add(apiV1Ws)

	authHandler := auth.NewAuthHandler(authManager)
	authHandler.Install(apiV1Ws)

	apiV1Ws.Route(
		apiV1Ws.GET("csrftoken/{action}").
			To(apiHandler.handleGetCsrfToken).
			Writes(api.CsrfToken{}))

	return wsContainer, nil
}

func (apiHandler *APIHandler) handleGetCsrfToken(request *restful.Request, response *restful.Response) {
	action := request.PathParameter("action")
	token := xsrftoken.Generate(apiHandler.cManager.CSRFKey(), "none", action)
	response.WriteHeaderAndEntity(http.StatusOK, api.CsrfToken{Token: token})
}

// Get namespaces from path parameter
func parseNamespacePathParameter(request *restful.Request) *common.NamespaceQuery {
	namespace := request.PathParameter("namespace")
	namespaces := strings.Split(namespace, ",")
	var noBlankNamespaces []string
	for _, n := range namespaces {
		n = strings.Trim(n, " ")
		if len(n) > 0 {
			noBlankNamespaces = append(noBlankNamespaces, n)
		}
	}

	return common.NewNamespaceQuery(noBlankNamespaces)
}

func parsePaginationPathParameter(request *restful.Request) *dataselect.PaginationQuery {
	itemsPerPage, err := strconv.ParseInt(request.QueryParameter("itemsPerPage"), 10, 0)
	if err != nil {
		return dataselect.NoPagination
	}

	page, err := strconv.ParseInt(request.QueryParameter("page"), 10, 0)
	if err != nil {
		return dataselect.NoPagination
	}

	// Frontend page start from 1 and backend start from 0
	return dataselect.NewPaginationQuery(int(itemsPerPage), int(page - 1))
}