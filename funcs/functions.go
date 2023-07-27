package Func

import (
	fetch "Func/API"
	"strings"
)

// Norepeat range an array and remove the repeated slices
func Norepeat(tab []string) []string {
	var newtab []string
	var exist bool
	for _, v := range tab {
		for _, y := range newtab {
			if y == v {
				exist = true
				break
			} else {
				exist = false
			}
		}
		if !exist {
			newtab = append(newtab, v)
		}
	}
	return newtab
}

// Norepeatart range an array and remove the repeated slices
func Norepeatart(tab []fetch.Artists) []fetch.Artists {
	var newtab []fetch.Artists
	var exist bool
	for _, v := range tab {
		for _, y := range newtab {
			if y.Id == v.Id {
				exist = true
				break
			} else {
				exist = false
			}
		}
		if !exist {
			newtab = append(newtab, v)
		}
	}
	return newtab
}

// Reverese returns a reversed string
func Reverse(string_to_reverse string) string {
	tab := strings.Split(string_to_reverse, "-")
	for i, j := 0, len(tab)-1; i < j; i, j = i+1, j-1 {
		tab[i], tab[j] = tab[j], tab[i]
	}
	res := strings.Join(tab, "-")
	return res
}

// Validtab returns then trimmed version of the array.
// it does an intersection of the previous and the current array by returning the common Ids.
func Validtab(Id []int, result []int) []int {
	var res []int
	if len(Id) == 0 {
		res = result
	} else {
		for _, v := range Id {
			for _, y := range result {
				if v == y {
					res = append(res, v)
				}
			}
		}
	}
	return res

}
