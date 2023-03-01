package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	_ "github.com/lib/pq"
	"github.com/tgrangeo/go_serv/database"
)

func main(){

	database.InitDb()
	database.InsertDB()
	database.DropTable()

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