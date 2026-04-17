package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//err := os.Mkdir("./FileOperation/basic", 0777)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//err := os.MkdirAll("./basic/test", 0777)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fil, err := os.Create("./basic/test/test.txt")
	//fil, err = os.Create("./basic/test/test1.txt")
	//fil, err = os.Create("./basic/test/test2.txt")
	//if err != nil {
	//	return
	//}

	dir, err3 := os.ReadDir("./basic/test")
	if err3 != nil {
		return
	}
	for _, d := range dir {
		fmt.Println("fileName :", d.Name())
		data, err5 := os.ReadFile(filepath.Join("./basic/test/", d.Name()))
		if err5 != nil {
			fmt.Println(err5)
		}
		fmt.Println(string(data))
	}

	//writeString, err1 := fil.WriteString("this is go file writer")
	//if err1 != nil {
	//	return
	//}
	//fmt.Println(writeString)
	//
	//at, err := os.ReadFile("./basic/test/test.txt")
	//if err != nil {
	//	return
	////}
	//fmt.Println(string(at))
	//defer func(fil *os.File) {
	//	err := fil.Close()
	//	if err != nil {
	//
	//	}
	//}(fil)

}
