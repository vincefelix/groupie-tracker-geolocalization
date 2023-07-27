package Func

import (
	fetch "Func/API"
	Geo "Func/Geoloc_API"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// Localization return the concert locations coordinates of a given artist
func Localization(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-----------------------------------")
	fmt.Println("🌍 World map loaded ✅")
	fmt.Println("------------------")

	if r.Method != "POST" && r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "405")
		return
	}
	// receiving the user search request
	artist := r.FormValue("q")
	initial := artist
	fmt.Println("🛎 request:", "<<", initial, ">>")
	artist = strings.ToLower(artist)
	fmt.Println("🌐 users request received ☑")

	//removing the spaces before the user's entry
	if len(artist) > 1 {
		for i := range artist {
			if artist[i] != ' ' {
				artist = string(artist[i:])
				break
			}

		}
	}

	//removing the spaces after the user's entry
	temp := strings.Fields(artist)
	artist = strings.Join(temp, " ")

	fmt.Println("🔍 searching artist's locations⬛ ◼ ◾")
	// storing the fetched datas
	fetched_res, error := fetch.Api_artists(w, r)
	resindex, error2 := fetch.Api_locations(w, r)

	// if an error occured while fetching datas
	if !error || !error2 {
		fmt.Println("❌error while fetching datas from artist api")
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}
	fmt.Println("☑ fetching data📊 from the Api ...")

	// retrieving the artists/band names
	names := []string{}
	for _, b := range fetched_res {
		names = append(names, b.Name)
	}

	// retrieving the artist's id
	var id int
	for i, v := range fetched_res {
		if strings.ToLower(v.Name) == artist {
			id = i + 1
			break
		}
	}
	fmt.Println("👨‍🎤 id", id)
	// retrieving the locations
	var locs []string
	for _, y := range resindex.Index {
		if y.Id == id {
			locs = y.Location
			break
		}
	}

	fmt.Println("📍 loc found", locs, "size:", len(locs))

	// mapping each city to its coordinates
	geomap := make(map[string]Geo.Coord)
	var coordinates Geo.Coord

	for _, xy := range locs {
		lat, lng, errcord := Geo.Geoloc(w, r, xy)
		if !errcord {
			fmt.Println("❌error while fetching locations coordinates")
			error_file := template.Must(template.ParseFiles("templates/error.html"))
			error_file.Execute(w, "500")
			return
		}
		coordinates.Lat = lat
		coordinates.Lng = lng
		geomap[xy] = coordinates
	}
	//final result
	fmt.Println("✅ addresses converted successfully")
	fmt.Println("🗺 ", geomap, "size:", len(geomap))

	// sending the results to the "artists" page
	file, err := template.ParseFiles("templates/location.html")
	if err != nil {
		fmt.Println("❌error while parsing location.html")
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}

	errgeo := file.Execute(w, struct {
		Name   []string
		Locmap map[string]Geo.Coord
		Q      string
	}{
		Name:   names,
		Locmap: geomap,
		Q:      initial,
	})

	if errgeo != nil {
		fmt.Println("❌error while excecuting the localization struct")
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
	}
	fmt.Println("✅ datas sent")
	fmt.Println("-----------------------------------")

}
