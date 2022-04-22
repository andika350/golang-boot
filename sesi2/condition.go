package main

import (
	"fmt"

	
)

func main()  {
	fmt.Println("================ If Else ================\n")
	var currentYear = 2022

	if age := currentYear - 2006; age < 17 {
		fmt.Println("Kamu belum punya KTP")
	} else {
		fmt.Println("Kamu sudah punya KTP")
	} 
	fmt.Println("\n================ Switch ================\n")

	score := 7

	switch score {
	case 8:
		fmt.Println("Sempurna")
	case 7:
		fmt.Println("Lumayan")
	default:
		fmt.Println("Kureng")
	}
	fmt.Println("\n================ Switch Operational ================\n")

	year := 1998

	switch {
	case year < 1945:
		fmt.Println("Indonesia belum merdeka")
	case (year > 1944) && (year < 1965):
		fmt.Println("Orde lama")
	case (year > 1964) && (year < 1996):
		fmt.Println("Orde baru")
	default: {
		fmt.Println("Era reformasi")
		}
	}


	test := 5

	switch {
	case test == 8:
		fmt.Println("Sempurna")
	case (test > 4) && (test < 8):
		fmt.Println("Oke")
		fallthrough
	case test < 5:
		fmt.Println("Study harder")
	}
	fmt.Println("\n================ Nested If ================\n")

	isLegal := false
	goJail := false

	if (isLegal == true) {
		goJail = false
	} else if (isLegal == false) {
		if(goJail == false) {
			goJail = false
		} else {
			goJail = true
		}
	}

	fmt.Println(goJail)
}