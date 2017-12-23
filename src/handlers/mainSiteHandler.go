package handlers

import (
	"axioma/conf"
	"net/http"
	"github.com/gorilla/mux"
)

func getStaticFilesFolder() http.Dir {
	ScriptsFolder := http.Dir(conf.PATH_STAIC_FILES)
	return ScriptsFolder
}

func staticFilesFolderHandler() http.Handler {
	staticFilesFolder := getStaticFilesFolder()
	staticFilesFolderServer := http.FileServer(staticFilesFolder)
	staticFilesFolderHandler := http.StripPrefix("/static", staticFilesFolderServer)
	return staticFilesFolderHandler
}

func indexHandler(responseWriter http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		SetHeaders_Main(responseWriter)
		http.ServeFile(responseWriter, r, conf.PATH_INDEX)
	} else {
		http.Redirect(responseWriter, r, "https://"+r.Host+r.URL.Path, http.StatusTemporaryRedirect)
	}
}

func HandleMainSite(router *mux.Router) {
	staticFolderHandler := staticFilesFolderHandler()
	router.PathPrefix("/static").Handler(staticFolderHandler)
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/panel", indexHandler)
}
