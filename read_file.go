package main

import (
	"fmt"
	"io/ioutil"
)

func test() {
	data, err := ioutil.ReadFile("./access.log")

	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("result", string(data))
}
