package _0_hello

import (
	"fmt"
	"os"
	"log"
	"io"
	"strings"
)

type customer struct {
	name string
	product string
}

func main()  {
	c1:=customer{"Vasu","Samsung Tab"}
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
	<ul>
	<li>`+c1.name+` : Purchased `+c1.product+`</li>
	</ul>
	</body>
	</html>
	`)
	//fmt.Println(str)
	wp,err:=os.Create("hellowebapp.html")
	if err!=nil{
		log.Fatalln(err)
	}
	defer wp.Close()

	io.Copy(wp,strings.NewReader(str))
}
