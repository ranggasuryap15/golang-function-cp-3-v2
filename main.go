package main

import (
	"fmt"
)

func CariDanUbahTerpendek(name1, name2 string) string {
	nameShortest := ""
	if len(name1) < len(name2) {
		nameShortest = name1
	} else if len(name1) > len(name2) {
		nameShortest = name2
	} else {
		if name1 < name2 {
			nameShortest = name1
		} else {
			nameShortest = name2
		}
	}
	return nameShortest// reset current name
}

func FindShortestName(names string) string {
	curName := ""
	minName := "hsdgjhajksghafjkdghjkldfshgljkdshgjkldfhsgjklhdfjhgkdlhl" // agar nilai defaultnya yang bukan keluar

	for _, c := range names {
		if string(c) == ";" || string(c) == " " || string(c) == "," {
			minName = CariDanUbahTerpendek(curName, minName)
			curName = "" // reset current name
		} else {
			curName += string(c) 
		}
	}

	if curName != "" {
		minName = CariDanUbahTerpendek(curName, minName)
	}
	return minName // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}
