package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("golang/01_template/firs_tpl.gohtml"))
}

type client struct {
	Uname string
	Pwd   string
}

func main() {
	c1 := client{"Praveen", "12345"}
	c2 := client{"Vasu", "67890"}
	clients := []client{c1, c2}
	err := tpl.Execute(os.Stdout, clients)
	if err != nil {
		log.Fatalln(err)
	}
}