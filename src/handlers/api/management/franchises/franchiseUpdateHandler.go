package franchises

import (
	"net/http"
	"github.com/gorilla/mux"
	"axioma/conf"
	"axioma/src/api/management/franchises"
	"axioma/src"
	"encoding/json"
	"io/ioutil"
	"axioma/src/handlers"
)

type franchiseRequestData struct {
	Token string `json:"token"`
	src.FranchiseRaw
}

func getFranchise(r *http.Request) (string, src.Franchise, *conf.ApiResponse) {
	var requestData franchiseRequestData
	var rawFranchise src.FranchiseRaw
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return requestData.Token, src.Franchise{}, conf.ERROR_DATA_FORMAT_INVALID_500
	}
	if len(body) > 0 {
		err = json.Unmarshal(body, &requestData)
		if err != nil {
			return requestData.Token, src.Franchise{}, conf.ERROR_DATA_FORMAT_INVALID_500
		}
	}
	rawFranchise = requestData.FranchiseRaw
	franchise, apiErr := rawFranchise.ToFranchise(true)
	return requestData.Token, franchise, apiErr
}

func updateFranchiseHandler(responseWriter http.ResponseWriter, r *http.Request){
	handlers.SetHeaders_API(responseWriter)
	responseWriter.WriteHeader(http.StatusOK)
	token, franchise, err := getFranchise(r); if err != nil {
		err.Execute(responseWriter)
		return
	}
	response := franchises.UpdateFranchise(token, franchise)
	response.Execute(responseWriter)
}

func HandleUpdateFranchise(router *mux.Router)  {
	router.HandleFunc("/franchises/{id}", updateFranchiseHandler).Methods("PUT", "OPTIONS")
}
