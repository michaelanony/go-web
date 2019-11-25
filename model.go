package main

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
)
var DB *sqlx.DB
type MainConfig struct {
	User string
	Pwd string
}
type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

func init(){
	var err error
	JsonParse :=NewJsonStruct()
	v := MainConfig{}
	JsonParse.Load("./config.json",&v)
	DB,err = sqlx.Open(`mysql`,
		v.User+`:`+v.Pwd+`@tcp(192.168.11.32:30006)/homedb?charset=utf8&parseTime=true`)
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

func UserAll()([]User,error){
	mods:=make([]User,0)
	err:=DB.Select(&mods,"SELECT * FROM `goUser`")
	return mods,err
}

