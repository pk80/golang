package main

import (
	"fmt"
	"os"
	"log"
	"io"
	"strings"
)

func main()  {
	name:="Praveen Kumar"
	str:=fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en"
	<head>
	<meta charset="UTF-8">
	<title>firstWebApp</title>
	</head>
	<body>
	<h1>`+name+` welcomes you to his first web app.</h1>
	</body>
	</html>
	`)
	//fmt.Println(str)
	wp,err:=os.Create("hellowebapp.html")
	if err!=nil{
		log.Fatal(err)
	}
	defer wp.Close()
	io.Copy(wp,strings.NewReader(str))
}
