package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

 func init(){
   tpl = template.Must(template.ParseGlob("tpl.gohtml"))
 }


 func main(){
	  err := tpl.Execute(os.Stdout,32);
	  if err!=nil{
		  log.Fatalln(err)
	  }
 }