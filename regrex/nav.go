package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"regexp"
)

func main() {
	url:="http://www.jikedaohang.com/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	sHtml, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(sHtml))
	//rex:=regexp.MustCompile(`<li class="layui-tab-li">(.*)</li>`)
	//result:=rex.FindAllStringSubmatch(string(sHtml),-1)
	//fmt.Println(len(result))
	//for key, value := range result {
	//	fmt.Println(key,value[1])
	//}
	//fmt.Println(string(sHtml))
	rex2:=regexp.MustCompile(`<div class="layui-tab layui-tab-card" id="bar\d*">([\s\S]*?)</div>`)
	result2:=rex2.FindAllStringSubmatch(string(sHtml),-1)
	fmt.Println(len(result2))
	//fmt.Println(result2[0][1])
	titleRex:=regexp.MustCompile(`<li class="layui-tab-li">(.*)</li>`)
	itemRex:=regexp.MustCompile(`<a class="box-item" href="(.*)" target="_blank">(.*)</a>`)
	for _, value := range result2 {
		contentHtml:=value[1]
		//fmt.Println(key,contentHtml)
		title:=titleRex.FindAllStringSubmatch(contentHtml,-1)
		item:=itemRex.FindAllStringSubmatch(contentHtml,-1)
		fmt.Println(title[0][1])
		for _, value := range item {
			fmt.Print(value[1],"\t")
			fmt.Println(value[2])
		}

		fmt.Println("----------")
	}


}