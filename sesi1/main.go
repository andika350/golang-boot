package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "nama saya"
	strArr := strings.Split(str, " ")

	fmt.Println("Sesi 1 Start")

	fmt.Print(strArr[0])
	fmt.Print(strArr[1])

}
