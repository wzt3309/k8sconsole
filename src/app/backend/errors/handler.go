package errors

import "net/http"

func HandleHTTPError(err error) int {
	return http.StatusInternalServerError
}
