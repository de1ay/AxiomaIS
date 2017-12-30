package franchises

import (
	"net/http"
	"github.com/gorilla/mux"
	"axioma/conf"
	"strings"
	"net/url"
	"axioma/src/handlers"
	"axioma/src/api/management/franchises"
	"strconv"
)

func getFranchiseID(r *http.Request) (string, uint, error) {
	token, _ := url.QueryUnescape(strings.TrimSpace(strings.ToLower(r.PostFormValue("token"))))
	unitID, err := strconv.ParseInt(strings.TrimSpace(strings.ToLower(r.PostFormValue("id"))), 10, 64)
	if err != nil {
		return token, 0, err
	}
	return token, uint(unitID), nil
}

func deleteFranchiseHandler(responseWriter http.ResponseWriter, r *http.Request){
	handlers.SetHeaders_API_POST(responseWriter)
	if r.Method == http.MethodPost || r.Method == http.MethodOptions {
		responseWriter.WriteHeader(http.StatusOK)
		token, id, err := getFranchiseID(r); if err != nil {
			conf.ERROR_DATA_FORMAT_INVALID_102.Execute(responseWriter)
			return
		}
		response := franchises.DeleteFranchise(token, id)
		response.Execute(responseWriter)
	} else {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		conf.ERROR_METHOD_NOT_ALLOWED_402.Execute(responseWriter)
	}
}

func HandleDeleteFranchise(router *mux.Router)  {
	router.HandleFunc("/franchise.delete", deleteFranchiseHandler)
}
