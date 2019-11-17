package main

import (
	"fmt"
	"os"
)

func main() {

	os.Mkdir("vilay", 0777)
	os.MkdirAll("ciciy/test", 0777)

	err := os.Remove("vilay")
	if err != nil {
		fmt.Println(err)
	}

	os.RemoveAll("ciciy")
}
