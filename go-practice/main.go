package main

import (
	"fmt"
	"strconv"
)

func main() {
	myValue, err := strconv.ParseInt("10", 0, 64)

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Println(myValue)
}
