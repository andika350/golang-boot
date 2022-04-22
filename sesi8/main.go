package main

import (
	"encoding/json"
	"fmt"
)


type Employee struct {
	FullName string `json:"full_name"`
	Email	 string `json:"email"`
	Age		 int	`json:"age"`
}

func main() {
	var jsonString = `
	 [
		{
			"full_name": "Andika Satrio",
			"email"	: "andikapangesu79@gmail.com",
			"age"		: 25
		},
		{
			"full_name": "Cindy Morg",
			"email"	: "cindymorg9@gmail.com",
			"age"		: 26
		}
	 ]
	`


	 //JSON TO STRUCT
	// var result Employee

	// var err = json.Unmarshal([]byte(jsonString), &result)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("full_name: ", result.FullName)
	// fmt.Println("email: ", result.Email)
	// fmt.Println("age: ", result.Age)

	//JSON TO MAP
	// var result map[string]interface{}

	// var err = json.Unmarshal([]byte(jsonString), &result)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("full_name: ", result["full_name"])
	// fmt.Println("email ", result["email"])
	// fmt.Println("age: ", result["age"])

	//JSON TO EMPTY INTERFACE
	// var temp interface{}

	// var err = json.Unmarshal([]byte(jsonString), &temp)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	
	// var result = temp.(map[string]interface{}) //assertion karena empty interface menjadi map string

	// fmt.Println("full_name: ", result["full_name"])
	// fmt.Println("email ", result["email"])
	// fmt.Println("age: ", result["age"])

	//JSON TO SLICE OF STRUCT
	var result []Employee

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("Index %d: %+v\n", i+1, v)
	}
}