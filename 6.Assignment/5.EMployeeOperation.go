package main

import "fmt"

type Employee struct {
	id        int
	firstName string
	lastName  string
}

func main() {
	var empId int = 100
	emplies := make([]Employee, 0)
	emplies = append(emplies, Employee{
		id:        empId,
		firstName: "John",
		lastName:  "Smith",
	})
	empId++
	emplies = append(emplies, Employee{
		id:        empId,
		firstName: "Zafar",
		lastName:  "Imam",
	})
	empId++
	emplies = append(emplies, Employee{
		id:        empId,
		firstName: "Jhon",
		lastName:  "stlaier",
	})
	empId++
	emplies = append(emplies, Employee{
		id:        empId,
		firstName: "Rahul",
		lastName:  "Jhon",
	})
	empId++
	emplies = append(emplies, Employee{
		id:        empId,
		firstName: "Rohan",
		lastName:  "Khan",
	})
	fmt.Println(emplies)
	changeLastName(emplies)
}

func changeLastName(empSl []Employee) {
	for i := range empSl {
		fmt.Printf("Enter Last Name: ")
		fmt.Scan(&empSl[i].lastName)
	}
	for _, e := range empSl {
		fmt.Println(" Name : " + e.firstName + " " + e.lastName)
	}
}
