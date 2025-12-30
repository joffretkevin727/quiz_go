package router

import (
	"net/http"
	"proj3/controller"
)

// CETTE FONCTION INITIALISE UN SERVEUR MUX, CONFIGURE LES ROUTES ET LES FICHIERS STATIQUES ET LE RETOURNE
func New() *http.ServeMux {
	mux := http.NewServeMux()

	//------------------- ROUTES -----------------------
	mux.HandleFunc("/home", controller.Home)
	mux.HandleFunc("/quiz", controller.Quiz)
	//--------------------------------------------------
	//-------------------- CSS -------------------------
	fileServer := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	//--------------------------------------------------
	return mux
}
