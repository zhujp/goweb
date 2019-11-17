package main

import (
	"fmt"
	"os"
)

func readFile(filename string) {
	fl, err := os.Open(filename)
	if err != nil {
		fmt.Println(fl, err)
		return
	}

	defer fl.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

func delFile(filename string) {
	os.Remove(filename)
}

func main() {
	userFile := "user.txt"
	//写文件
	//	fout, err := os.Create(userFile)

	//	if err != nil {
	//		fmt.Println(userFile, err)
	//		return
	//	}

	//	defer fout.Close()

	//	for i := 0; i < 10; i++ {
	//		fout.WriteString("name is vilay")
	//		fout.Write([]byte("Just a test!\r\n"))
	//	}
	//读文件
	//	readFile(userFile)
	//删除文件
	delFile(userFile)

}
