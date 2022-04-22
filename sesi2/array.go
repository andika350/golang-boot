package main 

import "fmt"

func main()  {
	// var numbers [4]int
	// numbers = [4]int{1,3,2,4}
	// var angka = [4]int{4,6,8,7}
	// var strings = [3]string{"Anton","Tini","Joko"}

	// fmt.Println(numbers, angka, strings)
	// fmt.Print(numbers, angka, strings,"\n")
	// fmt.Printf("%d,%d,%s\n", numbers, angka, strings)
	// fmt.Printf("%v,%v,%v\n", numbers, angka, strings)

	// var buah = [3]string{"Apel","Melon","Nanas"}

	// fmt.Print(buah,"\n")
	// fmt.Println(buah[0])
	// fmt.Printf("%v\n", buah[2])

	// buah[2] = "Buah Naga"	
	// fmt.Println(buah[2])
	// fmt.Printf("%#v\n", buah)


	// var nama = [3]string{"Andika","Cindy","Bayu"}

	// for i, v := range nama {
	// 	fmt.Printf("Index: %d Value: %v\n", i, v)
	// }

	// for i:=0; i<len(nama); i++ {
	// 	fmt.Printf("Index: %d Value: %v\n", i, nama[i])
	//}

	var arrAngka = [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(arrAngka)
	fmt.Println(arrAngka[0][2])

	

	// cars := []string{"Honda","Toyota","Ford","Porsche"}
	// newCars := []string{}

	// fmt.Println(cars)
	// fmt.Println(newCars)

	// cars[0] = "BMW"
	// newCars = append(newCars, "Tesla")

	// fmt.Println(cars)
	// fmt.Println(newCars)

	// cars = append(cars, newCars...)
	// fmt.Println(cars)

	// newCars = append(newCars, cars[2:3]...)
	// fmt.Println(newCars)
}