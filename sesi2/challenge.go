package main

import "fmt"




func main() {
	//Latihan 1 max:= 50 bilangan prima
	
	
	for i:=1; i<50; i++ {
		var count int = 0
		
		for j:=2; j<i; j++ {
			if i%j == 0 {
				count = 1
				break
			}
		}

		if i>1 && count == 0 {
			fmt.Print(i," ")
		}
	} 

	
	fmt.Println()

	//Latihan 2 fizzbuzz max:=50

	for i:=1; i<=50; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz ")
		} else if i%3 == 0 {
			fmt.Print("Fizz ")
		} else if i%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Print(i," ")
		}
	}
	fmt.Println()
}