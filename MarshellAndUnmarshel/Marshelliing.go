package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name    string
	Age     int
	Address string
	Email   string
}

func main() {

	employee := Employee{Name: "zafar imam",
		Age:     28,
		Address: "pune",
		Email:   "zafarimam@gmail.com",
	}

	fmt.Println(employee)
	jsonData, err := json.Marshal(employee)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonData))

	jsonPayload := "{\"Name\":\"zafar imam\",\"Age\":28,\"Address\":\"pune\",\"Email\":\"zafarimam@gmail.com\"}"
	var emp Employee
	err = json.Unmarshal([]byte(jsonPayload), &emp)
	if err != nil {
		return
	}
	fmt.Println(emp)
}
