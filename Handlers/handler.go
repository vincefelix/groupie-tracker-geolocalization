package handlers

import (
	Rt "Func/Routes"
	"fmt"
	"log"
	"net/http"
)

// handlers regroups all the routes supported by our servor.
// handlers launches it too.
func Handlers() {
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", Rt.Home)
	http.HandleFunc("/artists", Rt.Artists)
	http.HandleFunc("/info/", Rt.Info)
	http.HandleFunc("/search", Rt.Search)
	http.HandleFunc("/filter", Rt.Filter)
	http.HandleFunc("/localization", Rt.Localization)
	fmt.Println("server has started at : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
