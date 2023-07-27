package Func

import (
	fetch "Func/API"
	rep "Func/funcs"
	"fmt"
	"html/template"
	"net/http"
	"sort"
)

type Sendres struct {
	Memb       []string
	Top_artist []fetch.Artists
	Research   []fetch.Artists
	Res        []fetch.Artists
	Resloc     []string
	Nil        bool
	Q          string
}

// artists serves the route "/artists"
func Artists(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-----------------------------------")
	fmt.Println("🎶 Artist display")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "405")
		return
	}
	//parsing the artist page
	file, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		//sending metadata about the error to the servor
		w.WriteHeader(http.StatusInternalServerError)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}

	//storing the fetched datas
	res, error := fetch.Api_artists(w, r)
	resindex, error2 := fetch.Api_locations(w, r)

	//if an error occured while fetching datas
	if !error || !error2 {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}
	fmt.Println("☑ fetching datas📊 from the api ...")

	fmt.Println("☑ processing👨‍💻 datas ...  ")
	//retrieving the cities from the fectched location datas
	newlocs := []string{}
	for _, b := range resindex.Index {
		newlocs = append(newlocs, b.Location...)
	}

	//sorting the locations
	sort.Strings(newlocs)

	//retrieving the members from the fectched artists datas
	new_memb := []string{}
	for _, b := range res {
		new_memb = append(new_memb, b.Member...)
	}
	//removing the repeated members
	new_memb = rep.Norepeat(new_memb)

	//removing the repeated cities
	newlocs = rep.Norepeat(newlocs)
	final := Sendres{
		Memb:     new_memb,
		Res:      res,
		Resloc:   newlocs,
		Research: nil,
		Nil:      false,
		Q:        "",
	}
	err = file.Execute(w, final)

	if err != nil {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
	}
	fmt.Println("✅ Datas sent")
	fmt.Println("-----------------------------------")

}
