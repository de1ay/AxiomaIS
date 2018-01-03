package franchises

import (
	"net/http"
	"github.com/gorilla/mux"
	"axioma/src/api/management/franchises"
	"axioma/src/handlers"
)

func deleteFranchiseHandler(responseWriter http.ResponseWriter, r *http.Request){
	handlers.SetHeaders_API(responseWriter)
	responseWriter.WriteHeader(http.StatusOK)
	token, franchise, err := getFranchise(r); if err != nil {
		err.Execute(responseWriter)
		return
	}
	response := franchises.DeleteFranchise(token, franchise.ID)
	response.Execute(responseWriter)
}

func HandleDeleteFranchise(router *mux.Router)  {
	router.HandleFunc("/franchises/{id}", deleteFranchiseHandler).Methods("DELETE", "OPTIONS")
}
