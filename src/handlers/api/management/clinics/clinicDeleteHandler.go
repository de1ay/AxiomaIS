package clinics

import (
	"net/http"
	"github.com/gorilla/mux"
	"axioma/conf"
	"strings"
	"net/url"
	"axioma/src/handlers"
	"strconv"
	"axioma/src/api/management/clinics"
)

func getClinicID(r *http.Request) (string, uint, error) {
	token, _ := url.QueryUnescape(strings.TrimSpace(strings.ToLower(r.PostFormValue("token"))))
	unitID, err := strconv.ParseInt(strings.TrimSpace(strings.ToLower(r.PostFormValue("id"))), 10, 64)
	if err != nil {
		return token, 0, err
	}
	return token, uint(unitID), nil
}

func deleteClinicHandler(responseWriter http.ResponseWriter, r *http.Request){
	handlers.SetHeaders_API_POST(responseWriter)
	if r.Method == http.MethodPost || r.Method == http.MethodOptions {
		responseWriter.WriteHeader(http.StatusOK)
		token, id, err := getClinicID(r); if err != nil {
			conf.ERROR_DATA_FORMAT_INVALID_102.Execute(responseWriter)
			return
		}
		response := clinics.DeleteClinic(token, id)
		response.Execute(responseWriter)
	} else {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		conf.ERROR_METHOD_NOT_ALLOWED_402.Execute(responseWriter)
	}
}

func HandleDeleteClinic(router *mux.Router)  {
	router.HandleFunc("/clinic.delete", deleteClinicHandler)
}
