package clinics

import (
	"net/http"
	"github.com/gorilla/mux"
	"axioma/conf"
	"axioma/src/handlers"
	"axioma/src"
	"axioma/src/api/management/clinics"
	"io/ioutil"
	"encoding/json"
)

type clinicRequestData struct {
	Token string `json:"token"`
	src.ClinicRaw
}

func getClinic(r *http.Request) (string, src.Clinic, *conf.ApiResponse) {
	var requestData clinicRequestData
	var rawClinic src.ClinicRaw
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return requestData.Token, src.Clinic{}, conf.ERROR_DATA_FORMAT_INVALID_500
	}
	if len(body) > 0 {
		err = json.Unmarshal(body, &requestData)
		if err != nil {
			return requestData.Token, src.Clinic{}, conf.ERROR_DATA_FORMAT_INVALID_500
		}
	}
	rawClinic = requestData.ClinicRaw
	clinic, apiErr := rawClinic.ToClinic(true)
	return requestData.Token, clinic, apiErr
}

func addClinicHandler(responseWriter http.ResponseWriter, r *http.Request){
	handlers.SetHeaders_API(responseWriter)
	responseWriter.WriteHeader(http.StatusOK)
	token, clinic, err := getClinic(r); if err != nil {
		err.Execute(responseWriter)
		return
	}
	response := clinics.AddClinic(token, clinic)
	response.Execute(responseWriter)
}

func HandleAddClinic(router *mux.Router)  {
	router.HandleFunc("/clinics", addClinicHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/clinics/", addClinicHandler).Methods("POST", "OPTIONS")
}
