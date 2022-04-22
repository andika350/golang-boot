package main

import (
	"database/sql"
	"fmt"
	
	_"github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "andika"
	password = "Bloodyruin83!@#"
	dbname = "andika"
)

var (
	db *sql.DB
	err error
)

type Employee struct {
	ID		int
	Full_name string
	Email	string
	Age		int
	Division string
}



func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	// CreateEmployee()
	// GetEmployees()
	// UpdateEmployee()
	DeleteEmployee()
}


func CreateEmployee() {

	var employee = Employee{}

	sqlStatement := `
	INSERT INTO employees (full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	Returning *
	`

	err := db.QueryRow(sqlStatement, "Andika Satrio", "andikapangestu79@gmail.com", 25, "BackEnd").
	Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age,&employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data %+v\n", employee)
}

func GetEmployees() {
	var results = []Employee{}

	sqlStatement := `SELECT * from employees`

	rows, err := db.Query(sqlStatement)

	if err!= nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age,&employee.Division)

		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}

	fmt.Println("Employee datas:", results)
}

func UpdateEmployee() {
	sqlStatement := `
	UPDATE employees
	SET full_name = $2, email = $3, division = $4, age = $5
	WHERE id = $1;`
	res, err := db.Exec(sqlStatement, 1, "Cindy Morgan", "cindymorg@gmail.com","Public Relation", 26)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated data employee:", count)
}

func DeleteEmployee() {
	sqlStatement :=`
	DELETE from employees
	WHERE id = $1;
	`
	
	res, err := db.Exec(sqlStatement, 1)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted data amount:", count)
}