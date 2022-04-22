package main

import "fmt"

func main() {
	// age := 20

	// if age > 17 {
	// 	fmt.Println("Sudah punya KTP")
	// }

	// for i := 1; i <= 5; i++ {
	// 	fmt.Print(i, " ")
	// }

	// x := 0

	// for x < 5 {
	// 	fmt.Print("*")
	// 	x++
	// }

	// fmt.Println()

	//Latihan 1 max:= 50 bilangan prima

	//Latihan 2 fizzbuzz max:=50

	// for i:=1;i<=50;i++{

	// 	if i%15 == 0 {
	// 		fmt.Print("Fizz Buzz, ")
	// 	}else if i%3 == 0 {
	// 		fmt.Print("Fizz, ")
	// 	}else if i%5 == 0 {
	// 		fmt.Print("Buzz, ")
	// 	} else {
	// 		fmt.Print(i,", ")
	// 	}
	// }
	// fmt.Println()

	// var numbers [4]int = [4]int{1, 2, 3, 4}
	// var strings = [3]string{"Dika", "Satrio", "Pangestu"}

	// fmt.Printf("%v\t", numbers)
	// fmt.Printf("%v\t", strings)
	// fmt.Println()

	// nums := []int{5, 6, 7, 9}
	// fmt.Printf("Nilai: %d, panjangnya: %d\n", nums, len(nums))

	/*
	Latihan 
	2 produk tampilkan:
	-nama 
	-stok
	-harga

	total = (stok * harga) + ...*3 + (stok * harga)
	*/


	var barang []map[string]interface{}

	barang = []map[string]interface{}{
		{"nama":"Pensil","stok": 10,"harga": 3000},
		{"nama":"Pulpen","stok": 10,"harga": 4000},
		{"nama":"Buku","stok": 5,"harga": 6000},
	}

	total := 0

	for _,v := range barang {
		total+= v["harga"].(int) * v["stok"].(int)
	}

	fmt.Println(total)
}
