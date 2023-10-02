package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
		if err := r.ParseForm(); err != nil{
			fmt.Fprintln(w, "ParseForm() error : %v",err)
			return
		}
		fmt.Fprintf(w, "Post request succesful \n")
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprint(w, "Name : %s\n",name)
		fmt.Fprint(w,"Address : %s\n",address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"heellooooo!")
}

func main(){

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting a server at port 8000")
	if err := http.ListenAndServe(":8000",nil); err !=nil{
		log.Fatal(err)
	}
}
