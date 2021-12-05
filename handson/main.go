package main

import (
	"log"
	"os"
	"text/template"
)

type city struct{
	Name string 
	Address string 
	City string 
	Zip int 
	Region string
}

var tpl *template.Template;


func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}




func main(){

     data := []city{
		 {"Aman Jain","Nr colony","Bengaluru",5656423232,"Southern"},
		 {"Rahul","Eros line","Delhi",28976783232,"Central"},
		 {"Rohit","Chruscgh strom","Mumbai",287883232,"Northern"},
		 {"Dab","Nsdfsdfsdfy","Kolkata",211213232,"Southern"},
	}

	 err  := tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",data)
	 if err!=nil {
		 log.Fatalln(err);
	 }
}