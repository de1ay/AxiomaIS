/*
	Author: Nikita "De1aY" Ivanov
	Email: de1ay@nullteam.info
 */
package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"axioma/conf"
	"axioma/src"
	"crypto/tls"
	"time"
	"axioma/src/handlers/api/authorization"
	"axioma/src/handlers"
	"axioma/src/handlers/api/management/clinics"
	"axioma/src/handlers/api/management/franchises"
)

func main() {

	router := mux.NewRouter()

	mainSite := router.Host(conf.DOMAIN).Subrouter()
	wwwSite := router.Host("www." + conf.DOMAIN).Subrouter()
	apiSite := router.Host("api." + conf.DOMAIN).Subrouter()

	// Authorization
	authorization.HandleAuthorization(apiSite)
	authorization.HandleTokenValidation(apiSite)

	// Management
	clinics.HandleGetClinics(apiSite)
	clinics.HandleAddClinic(apiSite)
	clinics.HandleUpdateClinic(apiSite)
	clinics.HandleDeleteClinic(apiSite)
	franchises.HandleGetFranchises(apiSite)
	franchises.HandleAddFranchise(apiSite)
	franchises.HandleUpdateFranchise(apiSite)
	franchises.HandleDeleteFranchise(apiSite)

	// Main site
	handlers.HandleMainSite(mainSite)
	handlers.HandleMainSite(wwwSite)

	// Database
	src.Connection.Connect()
	src.Connection.Sync()

	src.Connection.Connection.Create(&src.Unit{Name: "Аксиома", Type: 0})
	src.Connection.Connection.Create(&src.User{Name: "никита", Surname: "иванов", Patronymic: "сергеевич", Password: "OTUxMjM2NTdObqafc8yiOprFyLVn3BhadW6XyYIWT-JYWeDR3MFHXICmFbISOvH1-UwR4-lALDrFWPUAGZ2VttPjAXWFhigdzSY=", Login: "De1ay", Role: 0, UnitID: 1})

	go HandleTLS(router)
	http.ListenAndServe(conf.HTTP_PORT, router)
}

func getTlsConfig() *tls.Config {
	TlsConfig := &tls.Config{}
	TlsConfig.Certificates = make([]tls.Certificate, 3)
	TlsConfig.Certificates[0], _ = tls.LoadX509KeyPair(conf.API_SITE_TLS, conf.API_SITE_TLS_KEY)
	TlsConfig.Certificates[1], _ = tls.LoadX509KeyPair(conf.MAIN_SITE_TLS, conf.MAIN_SITE_TLS_KEY)
	TlsConfig.Certificates[2], _ = tls.LoadX509KeyPair(conf.WWW_SITE_TLS, conf.WWW_SITE_TLS_KEY)
	TlsConfig.BuildNameToCertificate()
	return TlsConfig
}

func HandleTLS(router *mux.Router) {
	TlsConfig := getTlsConfig()
	Server := http.Server{
		TLSConfig: TlsConfig,
		Handler: router,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	TLSListener, _ := tls.Listen("tcp", conf.HTTPS_PORT, TlsConfig)
	Server.Serve(TLSListener)
}


