package main

import "fmt"


func main() {

	c := make (chan string)

	// go introduce("Dika", c)

	// go introduce("Cindy", c)

	// go introduce("Rehan", c)

	// msg1 := <- c
	// fmt.Println(msg1)

	// msg2 := <- c
	// fmt.Println(msg2)

	// msg3 := <- c
	// fmt.Println(msg3)

	students := []string{"Dika","Cindy","Rehan"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("Student", student)
			result := fmt.Sprintf("Hai, my name is %s", student)
			c <- result
		} (v)
	}

	for i := 1; i<= len(students); i++ {
		print(c)
	}

	close(c)

}


// func introduce (student string, c chan string) {
// 	result := fmt.Sprintf("Hei, my name is %s", student)

// 	c <- result
// }

func print( c chan string) {
	fmt.Println(<- c)
}