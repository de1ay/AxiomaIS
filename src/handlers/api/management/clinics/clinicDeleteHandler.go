package clinics

import (
	"net/http"
	"github.com/gorilla/mux"
	"axioma/src/handlers"
	"axioma/src/api/management/clinics"
)

func deleteClinicHandler(responseWriter http.ResponseWriter, r *http.Request){
	handlers.SetHeaders_API(responseWriter)
	responseWriter.WriteHeader(http.StatusOK)
	token, clinic, err := getClinic(r); if err != nil {
		err.Execute(responseWriter)
		return
	}
	response := clinics.DeleteClinic(token, clinic)
	response.Execute(responseWriter)
}

func HandleDeleteClinic(router *mux.Router)  {
	router.HandleFunc("/clinics/{id}", deleteClinicHandler).Methods("DELETE", "OPTIONS")
}
