package main

import "fmt"

type Person struct {
	nama string
	alamat string
	pekerjaan string
	alasan string
}

func main() {

	var input int
	fmt.Print("User index: ")
	fmt.Scanf("%d", &input)

	input = input - 1

	var people = []Person{
		{nama:"Andika",alamat:"Layar Raya 74",pekerjaan:"Backend Dev",alasan:"Mandatory"},
		{nama:"Alexander",alamat:"Kemuning 4",pekerjaan:"Backend Dev",alasan:"Udah Jago"},
		{nama:"Leo",alamat:"Borobudur raya 1",pekerjaan:"Backend Dev",alasan:"Bosen aja"},
	}

	fmt.Printf("Nama: %v\nAlamat: %v\nPekerjaan: %v\nAlasan Memilih Kelas Golang: %v\n", people[input].nama,people[input].alamat,people[input].pekerjaan,people[input].alasan)

}


