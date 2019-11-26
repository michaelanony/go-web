package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
var DB *sqlx.DB


func init(){
	var err error
	JsonParse :=NewJsonStruct()
	v := MainConfig{}
	JsonParse.Load("./config.json",&v)
	fmt.Println("sss",v.Host,v.Pwd)
	DB,err = sqlx.Open(`mysql`,
		v.User+`:`+v.Pwd+`@tcp(`+v.Host+`)/homedb?charset=utf8&parseTime=true`)
	if err!=nil{
		panic("Connection Error!")
	}
	if err = DB.Ping();err != nil{
		panic("Run Error")
	}
}

type User struct{
	Id int `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Num string `json:"num" form:"num"`
	Password string `json:"password" form:"password"`
	Logo string `json:"logo" form:"logo"`
	Age int `json:"age" form:"age"`
}

//获取所有用户信息
func UserAll()([]User,error){
	mods:=make([]User,0)
	err:=DB.Select(&mods,"SELECT * FROM `goUser`")
	return mods,err
}
//获取一条信息
//func Userone()([]User,error){
//
//}


