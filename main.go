/*
	Author: Nikita "De1aY" Ivanov
	Email: de1ay@nullteam.info
 */
package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"axioma/conf"
)

func main() {
	router := mux.NewRouter()
	/*
	main_site := router.Host(conf.DOMAIN).Subrouter()
	www_site := router.Host("www." + conf.DOMAIN).Subrouter()
	beta_site := router.Host("beta." + conf.DOMAIN).Subrouter()
	api_site := router.Host("api." + conf.DOMAIN).Subrouter()
	*/
	http.ListenAndServe(conf.HTTP_PORT, router)
}

