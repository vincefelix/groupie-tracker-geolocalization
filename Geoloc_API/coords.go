package Func

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// MAP API used: OPENCAGE
//Website: https://opencagedata.com/
//KEY: 6d8cc52873f14d3e8ccd9fcf0acfddfe (have to log first in order to get the key)
//URL structure: "https://api.opencagedata.com/geocode/v1/json?q=URI-ENCODED-PLACENAME&key=YOUR-KEY"

// --------        structure of the API according to needed datas        -------------------------//
type Geo struct {
	Results Result `json:"results"`
}
type Result []struct {
	Geometry Coord `json:"geometry"`
}
type Coord struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// -------------------------------------------------------------------------------------------------//

// Geoloc fetches data from the OPENCAGE geocoding api which provides addresses conversion.
// it takes as argument a concert location then returns its coordinates
func Geoloc(w http.ResponseWriter, r *http.Request, concert_loc string) (float64, float64, bool) {
	// boolean that checks whether there is an error or not
	error := true
	// getting datas from the api
	geodatas, err := http.Get("https://api.opencagedata.com/geocode/v1/json?q=" + concert_loc + "&key=6d8cc52873f14d3e8ccd9fcf0acfddfe")
	if err != nil {
		fmt.Println("❌ error while geocoding")
		error = false
		return 0, 0, error
	}

	// reading the geodatas
	geocontent, erread := io.ReadAll(geodatas.Body)
	if erread != nil {
		fmt.Println("❌ error while reading geodatas")
		error = false
		return 0, 0, error
	}

	//converting JSON file and storing geocontent
	var geocoords Geo
	errjson := json.Unmarshal(geocontent, &geocoords)
	if errjson != nil {
		fmt.Println("❌ error while converting json file")
		error = false
		return 0, 0, error
	}

	// retrieving the coordinates
	var initial_tab struct {
		Geometry Coord `json:"geometry"`
	}

	if len(geocoords.Results) > 2 {
		initial_tab = geocoords.Results[0] // initial structure
	} else if len(geocoords.Results) == 1 {
		initial_tab = geocoords.Results[0] // initial structure
	} else {
		initial_tab = geocoords.Results[1] // initial structure
	}

	var lat, lng float64
	lat = initial_tab.Geometry.Lat // affecting lat data
	lng = initial_tab.Geometry.Lng // affecting lng data
	return lat, lng, error

}
