package main

import (
	"fmt"
	"os"
)

func main() {
	name := "AMAN JAIN"
	tpl := `
	  <!DOCTYPE html>
	  <html> 
	<head>
	<title>
	Hello World
	
	</title>
     <body>
	<h1> Hello ` + name + `</h1>
	 </body>
	</head>

	  </html>
	  </html>
	 `

	 title := os.Args[1]
	 tpll := `
	  <!DOCTYPE html>
	  <html> 
	<head>
	<title>
	Hello World
	
	</title>
     <body>
	<h1> Hello ` + title + `</h1>
	 </body>
	</head>

	  </html>
	  </html>
	 `

	fmt.Println(tpl)
	fmt.Println(tpll)
}