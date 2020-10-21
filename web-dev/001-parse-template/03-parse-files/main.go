package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl2.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println()

	tpl, err = tpl.ParseFiles("tpl1.gohtml", "tpl3.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
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
