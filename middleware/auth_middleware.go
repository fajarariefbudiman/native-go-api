package middleware

import (
	"belajar-golang-api/helper"
	"belajar-golang-api/model/web"
	"net/http"
)

type Authmiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *Authmiddleware {
	return &Authmiddleware{Handler: handler}

}

func (middleware *Authmiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "RAHASIA" == r.Header.Get("X-API-Key") {
		//ok
		middleware.Handler.ServeHTTP(w, r)

	} else {
		//err
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
