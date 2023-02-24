package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"io/ioutil"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main(){
	http.HandleFunc("/", server)
	http.ListenAndServe(":3000", nil)
}

func openFile(file string)(string,error){
	ret, err := ioutil.ReadFile(file) 
	if err != nil {
		return "", err
	}
	return string(ret),err
}

func server(w http.ResponseWriter, r *http.Request){
	var html, err = openFile("index.html")
	if err != nil {
		w.WriteHeader(404)
	}
	fmt.Fprintf(w,html)
}