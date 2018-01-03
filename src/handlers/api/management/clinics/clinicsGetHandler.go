package clinics

import (
	"net/http"
	"github.com/gorilla/mux"
	"strings"
	"net/url"
	"axioma/src/handlers"
	"axioma/src/api/management/clinics"
)

func getToken(r *http.Request) string{
	token, _ := url.QueryUnescape(strings.TrimSpace(strings.ToLower(r.FormValue("token"))))
	return token
}

func getClinicsHandler(responseWriter http.ResponseWriter, r *http.Request){
	handlers.SetHeaders_API(responseWriter)
	responseWriter.WriteHeader(http.StatusOK)
	response := clinics.GetClinics(getToken(r))
	response.Execute(responseWriter)
}

func HandleGetClinics(router *mux.Router)  {
	router.HandleFunc("/clinics", getClinicsHandler).Methods("GET")
	router.HandleFunc("/clinics/", getClinicsHandler).Methods("GET")
}
