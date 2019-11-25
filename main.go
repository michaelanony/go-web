package main

import (
	"encoding/json"
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
func userAll(w http.ResponseWriter,r *http.Request)  {
	mods,_:=UserAll()
	buf,_:=json.Marshal(mods)
	w.Header().Set("Content-Type","application/json")
	w.Write(buf)
}
func main() {
	http.HandleFunc(`/`,listView)
	http.HandleFunc(`/user`,userAll)
	http.ListenAndServe(":80",nil)
	fmt.Println("run")
}
