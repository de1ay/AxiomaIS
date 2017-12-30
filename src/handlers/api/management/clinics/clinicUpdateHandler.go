package clinics

import (
	"net/http"
	"github.com/gorilla/mux"
	"axioma/conf"
	"strings"
	"net/url"
	"axioma/src/handlers"
	"axioma/src"
	"strconv"
	"axioma/src/api/management/clinics"
)

func getClinic(r *http.Request) (string, src.Unit, error) {
	token, _ := url.QueryUnescape(strings.TrimSpace(strings.ToLower(r.PostFormValue("token"))))
	var unit src.Unit
	unitID, err := strconv.ParseInt(strings.TrimSpace(strings.ToLower(r.PostFormValue("id"))), 10, 64)
	if err != nil {
		return token, unit, err
	}
	unitType, err := strconv.ParseInt(strings.TrimSpace(strings.ToLower(r.PostFormValue("type"))), 10, 64)
	if err != nil {
		return token, unit, err
	}
	unitParent, err := strconv.ParseInt(strings.TrimSpace(strings.ToLower(r.PostFormValue("parent"))), 10, 64)
	if err != nil {
		return token, unit, err
	}
	name := strings.TrimSpace(r.PostFormValue("name"))
	unit.ID = uint(unitID)
	unit.Name = name
	unit.Type = int(unitType)
	unit.Parent = int(unitParent)
	return token, unit, nil
}

func updateClinicHandler(responseWriter http.ResponseWriter, r *http.Request){
	handlers.SetHeaders_API_POST(responseWriter)
	if r.Method == http.MethodPost || r.Method == http.MethodOptions {
		responseWriter.WriteHeader(http.StatusOK)
		token, clinic, err := getClinic(r); if err != nil {
			conf.ERROR_DATA_FORMAT_INVALID_102.Execute(responseWriter)
			return
		}
		response := clinics.UpdateClinic(token, clinic)
		response.Execute(responseWriter)
	} else {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		conf.ERROR_METHOD_NOT_ALLOWED_402.Execute(responseWriter)
	}
}

func HandleUpdateClinic(router *mux.Router)  {
	router.HandleFunc("/clinic.update", updateClinicHandler)
}