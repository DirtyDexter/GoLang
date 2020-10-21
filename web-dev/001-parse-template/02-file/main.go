package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err2 := os.Create("index.html")
	if err2 != nil {
		log.Fatalln(err2)
	}
	defer nf.Close()

	err3 := tpl.Execute(nf, nil)
	if err3 != nil {
		log.Fatalln(err3)
	}
}
