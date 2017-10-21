package main


import (
	"log"
	"net/http"
)


func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/",fs)
	log.Println("Hello World!!!")
	http.ListenAndServe(":3525", nil)
}

