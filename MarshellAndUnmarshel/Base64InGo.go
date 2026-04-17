package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	setences := "this is go programing"

	toString := base64.StdEncoding.EncodeToString([]byte(setences))

	fmt.Println(toString) //dGhpcyBpcyBnbyBwcm9ncmFtaW5n

	endedString := "dGhpcyBpcyBnbyBwcm9ncmFtaW5n"

	decodeString, err := base64.StdEncoding.DecodeString(endedString)
	if err != nil {
		return
	}
	fmt.Println(string(decodeString))

}
