package main

import (
	"encoding/json"
	"fmt"
	"crypto/md5"
)

func main(){
	//对数组编码
	x:=[5]int{1,2,3,4,5}
	fmt.Println(x)
	s,error:=json.Marshal(x)
	if(error!=nil){
		panic(error)
	}
	fmt.Println(s)
	fmt.Println(string(s))

	//对map编码
	m:=make(map[string]float64)
	m["zhangsan"]=200.4
	s2,error2:=json.Marshal(x)
	if(error2!=nil){
		panic(error2)
	}
	fmt.Println(s2)
	fmt.Println(string(s2))

	//对对象进行编码
	student:=Student{"zhangsan",22}
	s3,error3:=json.Marshal(student)
	if(error3!=nil){
		panic(error3)
	}
	fmt.Println(s3)
	fmt.Println(string(s3))

	//对s3进行解码
	var s4 interface{}
	json.Unmarshal([]byte(s3),&s4)
	fmt.Println(s4)

	//md5
	h := md5.New()
	h.Write([]byte("aaaaa"))
	reult:=h.Sum([]byte(""))

	fmt.Println(reult)
	fmt.Printf("%x\n",string(reult))
	fmt.Printf("%x\n",reult)
}
// 对象模型
type Student struct {
	//首字母小写不想让外部访问 ``设置json转换的名字
	Name string `json:"first_name"`
	Age int
}