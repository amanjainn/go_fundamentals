/////////////////////////////////
// Writing Bytes to Files
// Go Playground: https://play.golang.org/p/Zc3KDG7kYvt
/////////////////////////////////

package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main(){
	 // opening the file in write-only mode if the file exists and then it truncates the file.
	// if the file doesn't exist it creates the file with 0644 permissions

	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY | os.O_TRUNC | os.O_CREATE,
		0644,
	)

	if err!= nil {
		log.Fatal(err)
	}
	 defer file.Close()
	byteSlice := []byte("Hello world") 	
	bytesWritten, err := file.Write(byteSlice)
	if err !=nil{
		log.Fatal(err)
	}
	log.Println("Bytes written",bytesWritten);
		// WRITING BYTES TO FILE USING ioutil.WriteFile()

	// ioutil.WriteFile() handles creating, opening, writing a slice of bytes and closing the file.
	// if the file doesn't exist WriteFile() creates it
	// and if it already exists the function will truncate it before writing to file.
	bs := []byte("Go programming is good")
	err = ioutil.WriteFile("c.txt",bs,0644)
	if err!=nil{
		log.Fatal(err)
	}
}