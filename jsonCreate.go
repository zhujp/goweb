package main

import (
	"encoding/json"
	"fmt"
)

//go 需要json字段大写，如果需要转换小写，使用struct的tag
type User struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

type Data struct {
	Code int    `json:"code"`
	Data []User `json:"data"`
}

func main() {

	var data Data

	data.Code = 100
	data.Data = append(data.Data, User{"vilay", 23})
	data.Data = append(data.Data, User{"cicy", 23})

	response, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(response))
}
