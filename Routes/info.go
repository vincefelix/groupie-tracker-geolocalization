package Func

import (
	fetch "Func/API"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"
	"strings"
)

// info serve the route ("/info").
func Info(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "405")
		return
	}
	fmt.Println("-----------------------------------")
	fmt.Println("ℹ Info display")
	fmt.Println("--------------")
	//retrieving the id from the url
	recup_id := path.Base(r.URL.Path)
	if r.URL.Path != "/info/"+recup_id {
		w.WriteHeader(http.StatusNotFound)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "404")
		return
	}

	if recup_id == "" {
		w.WriteHeader(http.StatusNotFound)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "404")
		return
	}

	//storing the api artist datas
	artists_data, error := fetch.Api_artists(w, r)
	//if an error occured while fetching datas
	if !error {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}
	fmt.Println("☑ fetching artist datas from the api ...")

	//converting the id into int and setting a limit
	id, err := strconv.Atoi(recup_id)
	if err != nil || id <= 0 || id > len(artists_data) {
		w.WriteHeader(http.StatusNotFound)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "404")
		return
	}
	fmt.Println("✅ 🆔checkced")

	//retrieving the informations corresponding to the id
	var artists_checked fetch.Artists
	for _, art := range artists_data {
		if art.Id == id {
			artists_checked = art
			break
		}
	}

	//storing the api dates datas
	dates_data, error1 := fetch.Api_dates(w, r)
	//if an error occured while fetching datas
	if !error1 {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}
	fmt.Println("☑ fetching date datas from the api ...")
	//retrieving the informations corresponding to the id
	var dates_checked fetch.Dates
	for _, days := range dates_data.Index {
		if days.Id == id {
			dates_checked = days
			break
		}
	}

	//Removing the "*" in front of dates
	for i := range dates_checked.Date {
		dates_checked.Date[i] = strings.ReplaceAll(dates_checked.Date[i], "*", "")
	}

	//storing the api locations datas
	locations_data, error2 := fetch.Api_locations(w, r)
	//if an error occured while fetching datas
	if !error2 {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}
	fmt.Println("☑ fetching location datas from the api ...")
	//retrieving the informations corresponding to the id
	var locations_checked fetch.Locations
	for _, city := range locations_data.Index {
		if city.Id == id {
			locations_checked = city
			break
		}
	}

	//modifying the locations
	for i := range locations_checked.Location {
		locations_checked.Location[i] = strings.Title(locations_checked.Location[i])
		locations_checked.Location[i] = strings.ReplaceAll(locations_checked.Location[i], "_", " ")
		locations_checked.Location[i] = strings.ReplaceAll(locations_checked.Location[i], "-", "    - - >    ")
	}

	//storing the api relations datas
	relations_data, error3 := fetch.Api_relation(w, r)

	//if an error occured while fetching datas
	if !error3 {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}
	fmt.Println("☑ fetching relation datas from the api ...")
	//retrieving the informations corresponding to the id
	var relations_checked fetch.Relations
	for _, link := range relations_data.Index {
		if link.Id == id {
			relations_checked = link
			break
		}
	}

	// modifying the relation map
	newmap := map[string][]string{}
	for position, day := range relations_checked.Dates_location {
		splitted := strings.ReplaceAll(position, "-", "\n")
		splitted = strings.ReplaceAll(splitted, "_", " ")
		newmap[splitted] = day
	}

	previd := artists_checked.Id - 1
	nextid := artists_checked.Id + 1
	//struct to excecute
	todisplay := struct {
		The_arts fetch.Artists
		Days     fetch.Dates
		Cities   fetch.Locations
		Links    map[string][]string
		Prev     int
		Next     int
	}{
		The_arts: artists_checked,
		Days:     dates_checked,
		Cities:   locations_checked,
		Links:    newmap,
		Prev:     previd,
		Next:     nextid,
	}

	file, errp := template.ParseFiles("templates/info.html")
	if errp != nil {
		//sending metadata about the error to the servor
		w.WriteHeader(http.StatusInternalServerError)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}

	err = file.Execute(w, todisplay)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
	}
	fmt.Println("✅ Datas sent")
	fmt.Println("-----------------------------------")

}
