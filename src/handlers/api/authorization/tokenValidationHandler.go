package authorization

import (
	"axioma/src/api/authorization"
	"net/http"
	"github.com/gorilla/mux"
	"axioma/conf"
	"strings"
	"net/url"
	"axioma/src/handlers"
)

func getToken(r *http.Request) string{
	token, _ := url.QueryUnescape(strings.TrimSpace(strings.ToLower(r.FormValue("token"))))
	return token
}

func tokenValidationHandler(responseWriter http.ResponseWriter, r *http.Request){
	handlers.SetHeaders_API_GET(responseWriter)
	if r.Method == http.MethodGet {
		responseWriter.WriteHeader(http.StatusOK)
		user, err := authorization.ValidateToken(getToken(r)); if err != nil {
			err.Execute(responseWriter)
		} else {
			//conf.REQUEST_SUCCESS_200.Execute(responseWriter)
			resp := &conf.ApiResponse{200, "success", &user}
			resp.Execute(responseWriter)
		}
	} else {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		conf.ERROR_METHOD_NOT_ALLOWED_402.Execute(responseWriter)
	}
}

func HandleTokenValidation(router *mux.Router)  {
	router.HandleFunc("/token.verify", tokenValidationHandler)
}
