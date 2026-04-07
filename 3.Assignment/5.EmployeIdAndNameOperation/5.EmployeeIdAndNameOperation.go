package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var m = make(map[int]string)
	for {
		var employeeid int
		var employeeName string
		fmt.Print("please enter employee id : ")
		fmt.Scan(&employeeid)
		if employeeid == -999 {
			break
		}
		fmt.Print("please enter employee name : ")
		scanner.Scan()
		employeeName = scanner.Text()
		m[employeeid] = employeeName
	}
	fmt.Println(m)
}
