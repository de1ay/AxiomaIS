package franchises

import (
	"net/http"
	"github.com/gorilla/mux"
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
	handlers.SetHeaders_API(responseWriter)
	responseWriter.WriteHeader(http.StatusOK)
	response := franchises.GetFranchises(getToken(r))
	response.Execute(responseWriter)
}

func HandleGetFranchises(router *mux.Router)  {
	router.HandleFunc("/franchises/", getFranchisesHandler).Methods("GET")
	router.HandleFunc("/franchises", getFranchisesHandler).Methods("GET")
}
