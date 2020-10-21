package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println()

	err = tpl.ExecuteTemplate(os.Stdout, "tpl1.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println()

	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println()

	err = tpl.ExecuteTemplate(os.Stdout, "tpl3.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println()
}
