package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func indexView(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("hello"))
}

func listView(w http.ResponseWriter,r *http.Request)  {
	buf,_:=ioutil.ReadFile("templates/list.html")
	w.Write(buf)
}
func main() {
	http.HandleFunc(`/`,listView)
	http.ListenAndServe(":80",nil)
	fmt.Println("run")
}
