package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "Dika", Age: 25, Division: "BackEnd"},
	{ID: 2, Name: "Cindy", Age: 25, Division: "Public Relation"},
	{ID: 3, Name: "Rehan", Age: 24, Division: "Sales"},
}

var PORT = ":8080"

//========================================================================================

func main() {

	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/c-employee", createEmployee)

	fmt.Println("Application is live and listening on port", PORT)
	http.ListenAndServe(PORT, nil)

}

//========================================================================================

func getEmployees(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	// if r.Method == "GET" {
	// 	json.NewEncoder(w).Encode(employees)
	// 	return
	// }

	// http.Error(w, "Invalid method", http.StatusBadRequest)

	if r.Method == "GET" {
		tpl, err := template.ParseFiles("template.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, employees)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)
		return
	}

	http.Error(w, "invalid method", http.StatusBadRequest)
}
