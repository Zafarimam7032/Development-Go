package loop

import "fmt"

func Loop() {
	fmt.Println("This is a loop function")
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("This is a for range loop")

	for i, v := range []string{"a", "b", "c"} {
		fmt.Println(i, v)
	}
	fmt.Println("This is a for range loop with string")
	for j, v := range `hello world` {
		fmt.Println(j, string(v))
	}
	i := 0
	for i < 5 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 5; j = j + 1 {
		fmt.Println(j)
	}
}
