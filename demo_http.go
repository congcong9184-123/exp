package main

import (
	"fmt"
	"net/http"
)


func main1()  {
	resp , err := http.Get("http://www.baidu.com/")
	if err != nil{
		fmt.Println("http.get err: ",err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("status:",resp.Status)

	buf := make([]byte,4*1024)
	var tmp string
	for {
		n ,err := resp.Body.Read(buf)
		if n == 0{
			fmt.Println("error : ",err)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println(" body : ",tmp)

}