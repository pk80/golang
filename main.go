package main

import (
	"log"
	"os"
	"text/template"
	"fmt"
)

var tpl *template.Template

func init() {
	//tpl = template.Must(template.ParseFiles("golang/01_template/03_access/access_tpl.gohtml"))
	tpl = template.Must(template.ParseFiles("access_tpl.gohtml"))
}

type client struct {
	Uname string
	Pwd   string
}

func main() {
	c1u:=""
	c1p:=""
	fmt.Println("Provide username: ")
	fmt.Scanln(&c1u)

	fmt.Println("Enter you password :")
	fmt.Scanln(&c1p)

	c1 := client{"Praveen", "12345"}
	c2 := client{"Vasu", "abcdef"}
	clients := []client{c1,c2}
	if c1u==string(c1.Uname) && c1p==string(c1.Pwd) {
		fmt.Println("Login Success...")
		err:=tpl.Execute(os.Stdout,clients)
		if err!=nil{
			log.Fatalln(err)
		}
	}else {
		fmt.Println("Login failed...")
	}
}