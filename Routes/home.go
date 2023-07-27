package Func

import (
	fetch "Func/API"
	rep "Func/funcs"
	"fmt"
	"html/template"
	"net/http"
)

// home serves the route "/"
func Home(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "405")
		return
	}

	//checking whether the route exists or not
	if r.URL.Path != "/" && r.URL.Path != "/artists" && r.URL.Path != "/info/" && r.URL.Path != "/search" && r.URL.Path != "/filter" && r.URL.Path != "/localization" {
		w.WriteHeader(http.StatusNotFound)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "404")
		return
	}

	file, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}

	tab, error := fetch.Api_artists(w, r)
	resindex, error2 := fetch.Api_locations(w, r)

	//if an error occured while fetching datas
	if !error || !error2 {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}
	//retrieving the cities from the fectched location datas
	newlocs := []string{}
	for _, b := range resindex.Index {
		newlocs = append(newlocs, b.Location...)
	}
	//removing the repeated cities
	newlocs = rep.Norepeat(newlocs)

	//retrieving the members from the fectched artists datas
	new_memb := []string{}
	for _, b := range tab {
		new_memb = append(new_memb, b.Member...)
	}
	new_memb = rep.Norepeat(new_memb)

	one := tab[0:6]
	final := Sendres{
		Memb:       new_memb,
		Top_artist: one,
		Res:        tab,
		Resloc:     newlocs,
		Research:   nil,
		Nil:        false,
		Q:          "",
	}

	err = file.Execute(w, final)
	if err != nil {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
	}
	fmt.Println("✅ homepage🏡 loaded successfully")
	fmt.Println("-----------------------------------")
}
