package main

import (
	"encoding/base64"
	"fmt"
)

//另外还有aes，des加解密
func base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src) //标准base64编码
	// return base64.URLEncoding.EncodeToString(src) //兼容base64编码
	// return base64.RawURLEncoding.EncodeToString(src) //兼容base64编码
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString((string(src)))
}
func main() {

	//加密
	// str := base64Encode([]byte("admin"))
	// fmt.Println(str)
	// fmt.Println(string(str))
	//解密
	str, err := base64Decode([]byte("YWRtaW4xMjM="))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(str))

}
