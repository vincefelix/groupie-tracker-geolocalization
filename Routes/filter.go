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
	"time"
)

// Search returns the results of the user's search request
func Filter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-----------------------------------")
	fmt.Println("🔎 filtering ⬛ ◼ ◾")
	fmt.Println("------------------")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "405")
		return
	}
	//receiving the user search request
	intervalle_min := r.FormValue("mincread")
	intervalle_max := r.FormValue("maxcread")
	first_al_min := r.FormValue("firstd_min")
	first_al_max := r.FormValue("firstd_max")
	membnum1 := r.FormValue("memb1")
	membnum2 := r.FormValue("memb2")
	membnum3 := r.FormValue("memb3")
	membnum4 := r.FormValue("memb4")
	membnum5 := r.FormValue("memb5")
	membnum6 := r.FormValue("memb6")
	membnum7 := r.FormValue("memb7")
	membnum8 := r.FormValue("memb8")

	city := r.FormValue("loca")
	fmt.Printf("this is the min inter: %v\nthis is the max inter:%v\nthis is the first album min date: %v\nthis is the first album max date: %v\nThis is the number of member (checkbox7): %v\nThis is the city: %v\n", intervalle_min, intervalle_max, first_al_min, first_al_max, membnum7, city)
	println("*****************************************************")

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
	//removing the repeated members
	new_memb = rep.Norepeat(new_memb)

	//if an error occured while fetching datas
	if !error || !error2 {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}

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

	//filtering the creation dates
	mincread, errconv1 := strconv.Atoi(intervalle_min)
	maxcread, errconv2 := strconv.Atoi(intervalle_max)
	if errconv2 == nil && errconv1 == nil {
		for o, p := range fetched_res {
			if mincread == maxcread {
				if p.Creation_date == mincread {
					result = append(result, o+1)
				}
			} else {
				if p.Creation_date >= mincread && p.Creation_date <= maxcread {
					result = append(result, o+1)
				}
			}
		}
	}
	fmt.Println("1)✅ creation date passed 🗓")
	fmt.Println(result)
	fmt.Println("------------------")

	//filtering first album
	// processing the first album dates
	//--retrieving the members from the fectched artists datas

	var resalbum []int
	minal, erral1 := time.Parse("2006-01-02", (first_al_min))
	maxal, erral2 := time.Parse("2006-01-02", (first_al_max))
	//converting dates type string into time
	fmt.Println("💱 converting the dates into time⏱type...")
	fmt.Println("min first album:", minal)
	fmt.Println("max first album:", maxal)

	if erral1 == nil && erral2 == nil {
		for index, el := range fetched_res {
			eltime, errel := time.Parse("2006-01-02", rep.Reverse(el.First_album))

			if errel == nil {
				if first_al_min == "" && first_al_max == "" {
					fmt.Println("gere")
					break
				} else if first_al_min != "" && first_al_max == "" {
					if eltime.Equal(minal) {
						resalbum = append(resalbum, index+1)
					}
				} else if first_al_min == "" && first_al_max != "" {
					if eltime.Equal(maxal) {
						resalbum = append(resalbum, index+1)
					}
				} else if minal == maxal {
					if eltime.Equal(minal) {
						resalbum = append(resalbum, index+1)
					}
				} else {
					if eltime.Equal(minal) || eltime.After(minal) {
						if eltime.Equal(maxal) || eltime.Before(maxal) {
							resalbum = append(resalbum, index+1)
						}
					}
				}
			}
		}
	}

	result = rep.Validtab(resalbum, result)

	fmt.Println("2)✅ first album passed 🎼")
	fmt.Println(result)
	fmt.Println("------------------")

	//searching members
	//--storing datas in an array
	var res_memb []int
	membarray := []string{membnum1, membnum2, membnum3, membnum4, membnum5, membnum6, membnum7, membnum8}
	fmt.Println("mem array", membarray)
	if len(membarray) != 0 {
		for l, m := range fetched_res {
			for _, v := range membarray {
				size, errsize := strconv.Atoi(v)
				if errsize == nil {
					if len(m.Member) == size {
						res_memb = append(res_memb, l+1)
					}
				}
			}
		}
	}
	fmt.Println("res memb", res_memb)
	result = rep.Validtab(res_memb, result)
	fmt.Println("3)✅ member passed 🎻")
	fmt.Println(result)
	fmt.Println("------------------")

	//searching locations
	var res_city []int
	if city != "" {
		for _, v := range newlocs {
			for _, y := range v.Location {
				if strings.Contains(y, city) {
					res_city = append(res_city, v.Id)
				}
			}
		}
	}

	result = rep.Validtab(res_city, result)
	fmt.Println("4)✅  location passed 📍")
	fmt.Println(result)
	fmt.Println("------------------")

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
	fmt.Println("📜 final result 📜")
	fmt.Println(result)
	fmt.Println("final result size")
	fmt.Println(len(result))
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
		Q:        "",
	}

	errexc := file.Execute(w, final)

	if errexc != nil {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
	}
	fmt.Println("✅ datas sent")
	fmt.Println("-----------------------------------")

}
