package main

import (
	"log"
	"os"
	"text/template"
)


func main(){
	tpl,err  := template.ParseFiles("two.gmao")
	if err != nil {
		 log.Fatalln(err)		
	}

	err =  tpl.Execute(os.Stdout,nil)
	if  err !=nil {
		log.Fatalln(err)
	}

	tpl,err = tpl.ParseFiles("two.gmao","vespa.gmao")
	if  err !=nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout,"vespa.gmao",nil)
	if err !=nil{
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout,"vespa.gmao",nil)
	if err !=nil{
		log.Fatalln(err)
	}
err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}



}
