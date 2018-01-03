package franchises

import (
	"net/http"
	"github.com/gorilla/mux"
	"axioma/src/handlers"
	"axioma/src/api/management/franchises"
)

func addFranchiseHandler(responseWriter http.ResponseWriter, r *http.Request){
	handlers.SetHeaders_API(responseWriter)
	responseWriter.WriteHeader(http.StatusOK)
	token, franchise, err := getFranchise(r); if err != nil {
		err.Execute(responseWriter)
		return
	}
	response := franchises.AddFranchise(token, franchise)
	response.Execute(responseWriter)
}

func HandleAddFranchise(router *mux.Router)  {
	router.HandleFunc("/franchises", addFranchiseHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/franchises/", addFranchiseHandler).Methods("POST", "OPTIONS")
}
