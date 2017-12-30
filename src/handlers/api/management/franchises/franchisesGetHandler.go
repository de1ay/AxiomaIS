package franchises

import (
	"net/http"
	"github.com/gorilla/mux"
	"axioma/conf"
	"strings"
	"net/url"
	"axioma/src/handlers"
	"axioma/src/api/management/franchises"
)

func getToken(r *http.Request) string{
	token, _ := url.QueryUnescape(strings.TrimSpace(strings.ToLower(r.FormValue("token"))))
	return token
}

func getFranchisesHandler(responseWriter http.ResponseWriter, r *http.Request){
	handlers.SetHeaders_API_GET(responseWriter)
	if r.Method == http.MethodGet {
		responseWriter.WriteHeader(http.StatusOK)
		response := franchises.GetFranchises(getToken(r))
		response.Execute(responseWriter)
	} else {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		conf.ERROR_METHOD_NOT_ALLOWED_402.Execute(responseWriter)
	}
}

func HandleGetFranchises(router *mux.Router)  {
	router.HandleFunc("/franchises.get", getFranchisesHandler)
}
