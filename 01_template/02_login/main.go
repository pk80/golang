package main

import (
	//"fmt"
	"os"
	"log"
	"io"
	"strings"
)

func main()  {
	hpt:=`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>LoginForm</title>
	</head>
	<body>
	<p><b>Welcome to Web App login form</b></p>
	<p>Username :<input id="uname" type="text" /></p>
	<p>Password :<input id="pwd" type="password" /></p>
	<p><input id="login" type="button" value="Login" /></p>
	</body>
	</html>
	`
	//fmt.Println(hpt)
	tpl,err:=os.Create("loginform.html")
	if err!=nil{
		log.Fatalln(err)
	}
	defer tpl.Close()
	io.Copy(tpl,strings.NewReader(hpt))
}
