package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.Create("plan.txt")
	if err != nil {
		fmt.Println(err)
	}
	p := "I’m not afraid of difficulties and insist on learning programming"
	fp.Write([]byte(p)) //将内容写入文件
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fp.Close()
	file, err := os.Open("plan.txt") //打开文件
	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, 1024) //创建byte的slice用于接受文件读取数据
	fmt.Println("以下是文件内容:")
	for {
		len, _ := file.Read(buf)
		//读取字节为0的时候跳出循环
		if len == 0 {
			break
		}
		fmt.Println(string(buf))
	}
	file.Close()
}
