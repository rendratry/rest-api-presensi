package middleware

import (
	"golang/rest-api-presensi/entity/domain"
	"golang/rest-api-presensi/entity/web"
	"golang/rest-api-presensi/helper"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	//"12345678" == request.Header.Get("X-API-Key") &&
	if request.Header.Get("Authorization") != "" {
		jwt := request.Header.Get("Authorization")
		iss := domain.JwtClaims{}
		iss = helper.ValidateJWT(jwt, "secret-key")
		status := helper.GetIssuer(iss.Subject)
		if status {
			// ok
			middleware.Handler.ServeHTTP(writer, request)
		} else {
			// error
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}

			helper.WriteToResponseBody(writer, webResponse)
		}
	} else {
		if request.URL.Path == "/api/login" || request.URL.Path == "/api/register" ||
			strings.Contains(request.URL.Path, "/api/emailcheck/") ||
			strings.Contains(request.URL.Path, "/api/karyawancheck/") ||
			strings.Contains(request.URL.Path, "/api/sendotp") ||
			strings.Contains(request.URL.Path, "/api/otpvalidation") {
			if request.Header.Get("X-Api-Key") == "12345678" {
				// ok
				middleware.Handler.ServeHTTP(writer, request)
			} else {
				// error
				writer.Header().Set("Content-Type", "application/json")
				writer.WriteHeader(http.StatusUnauthorized)

				webResponse := web.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "UNAUTHORIZED",
				}

				helper.WriteToResponseBody(writer, webResponse)
			}
		} else {
			// error
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}

			helper.WriteToResponseBody(writer, webResponse)
		}
	}

}
