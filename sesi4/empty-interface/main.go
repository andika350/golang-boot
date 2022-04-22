package main

import "fmt"



func main() {
	
	var randomValues interface{}

	_ = randomValues

	randomValues = "Jalan Sudirman"
	randomValues =	20
	// randomValues = true
	// randomValues = []string{"Dika","Cindy"}

	fmt.Println(randomValues)

    var v = randomValues.(int) * 2 //method assertion pada empty interface

	fmt.Println(v)

	if value, ok:= randomValues.(int); ok == true {
		randomValues = value + 5
		fmt.Println(randomValues)
	}
	
	rs := []interface{}{1, "Dika", true, 2, "Cindy", true}
	
	rm := map[string] interface{} {
		"Name" :"Dika",
		"Gender":"Male",
		"Age":25,
	}

	fmt.Println(rs)
	fmt.Println(rm)

	for i, v := range rs {
		fmt.Printf("Index:%d Value:%v\n",i,v)
	}

	for _, v := range rm {
		fmt.Printf("Value:%v\n", v)
	}

}