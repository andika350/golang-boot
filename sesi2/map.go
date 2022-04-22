package main

import "fmt"

func main() {

	// var person map[string]string //declare

	// person = map[string]string{} //initiate

	// person["name"] = "Dika"
	// person["alamat"] = "Jl. Layar Raya No. 74"
	// person["zodiac"] = "Cancer"

	// fmt.Printf("Name: %s\nAddress: %s\nZodiac: %s\n", person["name"], person["alamat"], person["zodiac"])

	// var car1 = map[string]string{
	// 	"merk": "Toyota",
	// 	"tipe": "Agya",
	// 	"cc":   "1200cc",
	// }

	// for key, value := range car1 {
	// 	fmt.Println(key, ":", value)
	// }

	// fmt.Println(person)
	// delete(person, "zodiac")
	// fmt.Println(person)

// 	value, exist := car1["merk"]

// 	if exist {
// 		fmt.Println(value)
// 	} else {
// 		fmt.Println("Value didnt exist")
// 	}

// 	delete(car1, "merk")

// 	value, exist = car1["merk"]

// 	if exist {
// 		fmt.Println(value)
// 	} else {
// 		fmt.Println("Value has been deleted")
// 	}

// 	fmt.Println(car1)

	// var people = []map[string]string{
	// 	{"nama":"Dika","age":"25"},
	// 	{"nama":"Satrio","age":"24"},
	// 	{"nama":"Pangestu","age":"26"},
	// }

	// for k, v := range people {
	// 	fmt.Printf("index: %d\tnama: %s\tage: %v\n",k,v["nama"],v["age"])
	// }

	//Aliase

	var a uint8 = 10
	var b byte

	b = a
	_ = b

	type second = uint
	var hour second = 3600

	fmt.Printf("hour type: %T\n", hour)

}
