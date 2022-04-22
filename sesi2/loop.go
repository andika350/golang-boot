package main

import "fmt"

func main()  {
	//normal  loop
	for i:= 1; i<4; i++ {
		fmt.Println("Loop ke-", i)
	}
	fmt.Println("----------------------------")

	//while loop
	var x = 1

	for x < 4 {
		fmt.Println("While ke-", x)
		x++
	}
	fmt.Println("----------------------------")

	//break loop
	var y = 1

	for {
		fmt.Println("Break ke-", y)

		y++
		if y == 4 {
			break
		}
	}
	fmt.Println("----------------------------")

	//loop dengan break dan continue
	for i:=1; i<=10; i++ {
		if i%2 == 1 {
			continue
		}
		if i>8 {
			break
		}

		fmt.Println("Angka", i)
	}
	fmt.Println("----------------------------")

	//nested loop
	for i:=0; i<5; i++ {
		for j:=1; j<=i; j++ {
			fmt.Print(j," ")
		}
		fmt.Println()
	}
	fmt.Println("----------------------------")

	//break to label
	outerloop:
	for i:=0; i<3; i++ {
		fmt.Println("Loop ke-", i)
		for j:=0; j<3; j++ {
			if i==2 {
				break outerloop
			}
			fmt.Print(j," ")
		}
		fmt.Println()
	}
}