package main

import (
	"net/http"
	"fmt"
	//"io"

	//"encoding/json"
	"io/ioutil"
)

func main() {
	//http get post
	resp,error:=http.Get("http://www.baidu.com")
	if error!=nil{
		panic(error)
	}
	defer resp.Body.Close()
	//fmt.Printf(json.Unmarshal([]byte()))
	fmt.Println(resp.StatusCode)
	body, error:=ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	//开启服务监听 ListenAndServe
	//http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Println("hello")
	//	io.WriteString(writer, "hello, world!\n")
	//})
	//
	//fmt.Println("启动成功。。。")
	//http.ListenAndServe(":8080",nil)
	//fmt.Println("启动成功。。。")
}
