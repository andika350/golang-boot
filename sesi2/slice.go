package main

import "fmt"

func main()  {
	
	// var fruits = []string{"apel","pisang","mangga"}

	// _=fruits
	// fmt.Printf("%#v\n",fruits)

	// //slice with make

	// var buah = make([]string,3)

	
	// fmt.Printf("%#v\n",buah)

	// buah = append(buah, "Naga","Jambu","Melon")
	// fmt.Printf("%#v\n",buah)

	// //append allipsis
	// var fruits1 = []string{"apple","banana","mango"}
	// var fruits2 = []string{"durian","pineapple","starfruit"}

	// fruits1 = append(fruits1, fruits2...)
	// fmt.Printf("%#v\n",fruits1)

	// kendaraan := []string{"Motor","Mobil"}
	// kendaraan1 := []string{"Motor Listrik","Mobil Listrik"}

	// nn := copy(kendaraan, kendaraan1)

	// fmt.Println("Kendaraan =>",kendaraan)
	// fmt.Println("Kendaraan1 =>",kendaraan1)
	// fmt.Println("Copied Elements =>",nn)

	// var nama = []string{"nana","nini","nene","nono","nunu"}

	// var nama1 = nama[0:1]
	// fmt.Println(nama1)

	// var nama2 = nama[1:3]
	// fmt.Println(nama2)

	// var nama3 = nama[:]
	// fmt.Println(nama3)

	// var nama4 = nama[2:3]
	// fmt.Println(nama4)

	// var fruits1 = []string{"apple","banana","durian","jambu","starfruit"}

	// fmt.Println(fruits1)
	// fruits1 = append(fruits1[1:], "rambutan")
	// fmt.Println(fruits1)


	cars := []string{"Ford","Honda","Audi","Toyota"}
	newCars := []string{}
	
	newCars = append(newCars, cars[0:3]...)
	fmt.Println(newCars)
}