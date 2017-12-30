package franchises

import (
	"net/http"
	"github.com/gorilla/mux"
	"axioma/conf"
	"strings"
	"net/url"
	"axioma/src/handlers"
	"axioma/src/api/management/franchises"
	"axioma/src"
)

func getFranchiseName(r *http.Request) (string, src.Unit) {
	var unit src.Unit
	token, _ := url.QueryUnescape(strings.TrimSpace(strings.ToLower(r.PostFormValue("token"))))
	unit.Name = strings.TrimSpace(r.PostFormValue("name"))
	unit.Type = 1
	unit.Parent = 0
	return token, unit
}

func addFranchiseHandler(responseWriter http.ResponseWriter, r *http.Request){
	handlers.SetHeaders_API_POST(responseWriter)
	if r.Method == http.MethodPost || r.Method == http.MethodOptions {
		responseWriter.WriteHeader(http.StatusOK)
		token, franchise := getFranchiseName(r)
		response := franchises.AddFranchise(token, franchise)
		response.Execute(responseWriter)
	} else {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		conf.ERROR_METHOD_NOT_ALLOWED_402.Execute(responseWriter)
	}
}

func HandleAddFranchise(router *mux.Router)  {
	router.HandleFunc("/franchise.add", addFranchiseHandler)
}
