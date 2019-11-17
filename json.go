package main

import (
	"encoding/json"
	"fmt"
)

//Json 结构体字段必须都是可导出的也就是首字母大写，公有
type User struct {
	Username string
	Age      int
}

//type Userslice struct {
//	Users []User
//}

type Data struct {
	Code int
	Data []User
}

func main() {

	var s Data
	//	str := `{"users":[{"username":"vilay","age":30},{"username":"cicy","age":30}]}`
	str := `{"code":100,"data":[{"username":"vilay","age":30},{"username":"cicy","age":30}]}`
	json.Unmarshal([]byte(str), &s)
	//	fmt.Println(s)
	//	fmt.Println(s.Code)
	if s.Code == 100 {
		for _, v := range s.Data {
			fmt.Println(v.Username)
		}
	}
	//	for k, v := range s {
	//		fmt.Println(v.Username)
	//	}
}
