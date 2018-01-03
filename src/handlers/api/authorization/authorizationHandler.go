package authorization

import (
	"net/http"
	"strings"
	"github.com/gorilla/mux"
	"axioma/src/api/authorization"
	"axioma/src/handlers"
)

func getAuthorizationData(r *http.Request) authorization.AuthData{
	userLogin := strings.TrimSpace(r.FormValue("login"))
	userPassword := strings.TrimSpace(r.FormValue("password"))
	return authorization.AuthData{userLogin, userPassword, -1}
}

func loginAndPasswordAuthHandler(responseWriter http.ResponseWriter, r *http.Request){
	handlers.SetHeaders_API(responseWriter)
	responseWriter.WriteHeader(http.StatusOK)
	authData := getAuthorizationData(r)
	authData.Authorize(responseWriter)
}

func HandleAuthorization(router *mux.Router)  {
	router.HandleFunc("/token.get", loginAndPasswordAuthHandler).Methods("GET")
}
