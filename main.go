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
	"axioma/src/handlers/api/management/franchises"
	"axioma/src/handlers/api/management/clinics"
)

func main() {

	router := mux.NewRouter()

	mainSite := router.Host(conf.DOMAIN).Subrouter()
	wwwSite := router.Host("www." + conf.DOMAIN).Subrouter()
	apiSite := router.Host("api." + conf.DOMAIN).Subrouter()

	// Authorization
	authorization.HandleAuthorization(apiSite)
	authorization.HandleTokenValidation(apiSite)

	// Franchises
	franchises.HandleGetFranchises(apiSite)
	franchises.HandleAddFranchise(apiSite)
	franchises.HandleUpdateFranchise(apiSite)
	franchises.HandleDeleteFranchise(apiSite)
	// Clinics
	clinics.HandleGetClinics(apiSite)
	clinics.HandleAddClinic(apiSite)
	clinics.HandleUpdateClinic(apiSite)
	clinics.HandleDeleteClinic(apiSite)

	// Main site
	handlers.HandleMainSite(mainSite)
	handlers.HandleMainSite(wwwSite)

	// Database
	src.Connection.Connect()
	src.Connection.Sync()

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


