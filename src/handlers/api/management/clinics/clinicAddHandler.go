package clinics

import (
	"net/http"
	"github.com/gorilla/mux"
	"axioma/conf"
	"strings"
	"net/url"
	"axioma/src/handlers"
	"axioma/src"
	"axioma/src/api/management/clinics"
	"strconv"
)

func getClinicName(r *http.Request) (string, src.Unit, error) {
	var unit src.Unit
	token, _ := url.QueryUnescape(strings.TrimSpace(strings.ToLower(r.PostFormValue("token"))))
	unit.Name = strings.TrimSpace(r.PostFormValue("name"))
	unitParent, err := strconv.ParseInt(strings.TrimSpace(strings.ToLower(r.PostFormValue("parent"))), 10, 64)
	if err != nil {
		return token, unit, err
	}
	unit.Type = 2
	unit.Parent = int(unitParent)
	return token, unit, nil
}

func addClinicHandler(responseWriter http.ResponseWriter, r *http.Request){
	handlers.SetHeaders_API_POST(responseWriter)
	if r.Method == http.MethodPost || r.Method == http.MethodOptions {
		responseWriter.WriteHeader(http.StatusOK)
		token, clinic, err := getClinicName(r); if err != nil {
			conf.ERROR_DATA_FORMAT_INVALID_102.Execute(responseWriter)
			return
		}
		response := clinics.AddClinic(token, clinic)
		response.Execute(responseWriter)
	} else {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		conf.ERROR_METHOD_NOT_ALLOWED_402.Execute(responseWriter)
	}
}

func HandleAddClinic(router *mux.Router)  {
	router.HandleFunc("/clinic.add", addClinicHandler)
}
