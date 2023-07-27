package Func

import (
	fetch "Func/API"
	rep "Func/funcs"
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// Search returns the results of the user's search request
func Search(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "405")
		return
	}
	fmt.Println("-----------------------------------")
	fmt.Println("🔍 searching ⬛ ◼ ◾")

	//receiving the user search request
	element := r.FormValue("q")
	initial := element
	//removing the spaces before the user's entry
	if len(element) > 1 {
		for i := range element {
			if element[i] != ' ' {
				element = string(element[i:])
				break
			}

		}
	}
	splitelement := strings.Split(element, "  -  ")
	for i := range splitelement {
		splitelement[i] = (strings.ToLower(splitelement[i]))
	}
	fmt.Println("users request received ☑")
	//storing the fetched datas
	fetched_res, error := fetch.Api_artists(w, r)
	resindex, error2 := fetch.Api_locations(w, r)

	//retrieving the cities from the fetched location data
	locs := []string{}
	for _, b := range resindex.Index {
		locs = append(locs, b.Location...)
	}
	//Removing the repeated cities and capitalizing the first letters
	locs = rep.Norepeat(locs)

	//sorting the locations
	sort.Strings(locs)

	//retrieving the members from the fectched artists datas
	new_memb := []string{}
	for _, b := range fetched_res {
		new_memb = append(new_memb, b.Member...)
	}
	new_memb = rep.Norepeat(new_memb)
	//if an error occured while fetching datas
	if !error || !error2 {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}
	fmt.Println("☑ fetching data📊 from the Api ...")

	//creating an array of locations with the corresponding id
	var newlocs []fetch.Locations
	newlocs = append(newlocs, resindex.Index...)
	for i := range newlocs {
		for h := range newlocs[i].Location {
			newlocs[i].Location[h] = (strings.ToLower(newlocs[i].Location[h]))
		}
	}

	//filtering the id that corresponds to the search request
	var result []int

	fmt.Println("☑ Ranging database ...")
	if len(splitelement) == 2 {
		switch {
		case splitelement[1] == "location":
			for _, v := range newlocs {
				for _, y := range v.Location {
					if strings.Contains(y, splitelement[0]) {
						result = append(result, v.Id)
					}
				}
			}

		case splitelement[1] == "member":
			for l, m := range fetched_res {
				for _, memb := range m.Member {
					if strings.Contains(strings.ToLower(memb), splitelement[0]) {
						result = append(result, l+1)
					}
				}
			}

		default:
			creationD, errconv := strconv.Atoi(splitelement[0])
			if errconv == nil {
				for j, k := range fetched_res {
					if creationD == k.Creation_date {
						result = append(result, j+1)
					}
				}

			} else {

				for j, k := range fetched_res {
					if strings.Contains(strings.ToLower(k.First_album), splitelement[0]) || strings.Contains(strings.ToLower(k.Name), splitelement[0]) {
						result = append(result, j+1)
					}
				}
			}
		}
	} else if len(splitelement) == 1 {
		tabtemp := strings.Fields(splitelement[0])
		splitelement[0] = strings.Join(tabtemp, " ")
		//user searching for creation date
		onel, errconv2 := strconv.Atoi(splitelement[0])
		if errconv2 == nil {
			for o, p := range fetched_res {
				if onel == p.Creation_date {
					result = append(result, o+1)
				}
			}
		}
		//searching first album or group name
		for x, w := range fetched_res {

			if strings.Contains(strings.ToLower(w.First_album), splitelement[0]) || strings.Contains(strings.ToLower(w.Name), splitelement[0]) {
				result = append(result, x+1)
			}
		}

		//searching members
		for l, m := range fetched_res {
			for _, memb := range m.Member {
				if strings.Contains(strings.ToLower(memb), splitelement[0]) {
					result = append(result, l+1)
				}
			}
		}

		//searching locations
		for _, v := range newlocs {
			for _, y := range v.Location {
				if strings.Contains(y, splitelement[0]) {
					result = append(result, v.Id)
				}
			}
		}

	}
	fmt.Println("☑ database checked")

	//retrieving the informations corresponding to the id's request
	var resart []fetch.Artists
	for i := 0; i < len(result); {
		for _, founded := range fetched_res {
			if founded.Id == result[i] {
				resart = append(resart, founded)

			}
		}
		i++
	}

	resart = rep.Norepeatart(resart)

	// sending the results to the "artists" page
	file, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}

	final := Sendres{
		Memb:     new_memb,
		Research: resart,
		Res:      fetched_res,
		Resloc:   locs,
		Nil:      true,
		Q:        initial,
	}

	errexc := file.Execute(w, final)

	if errexc != nil {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
	}
	fmt.Println("✅ datas sent")
	fmt.Println("-----------------------------------")

}
