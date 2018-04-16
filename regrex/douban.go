package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"regexp"
)

func main() {
	url:="https://movie.douban.com/subject/26580232/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	sHtml, _ := ioutil.ReadAll(resp.Body)
	rex:=regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*)</span>`)
	result:=rex.FindAllStringSubmatch(string(sHtml),-1)
	for _,v :=range result {
		fmt.Println(v[1])
	}

	rexScore := regexp.MustCompile(`<strong\s*class="ll\s*rating_num"\s*property="v:average">(.*)</strong>`)
	scoreArry := rexScore.FindAllStringSubmatch(string(sHtml), -1)
	//for key, value := range scoreArry {
	//	fmt.Println(key,value)
	//}
	//fmt.Println(len(scoreArry[0]))
	fmt.Println(scoreArry[0][1])
	//fmt.Println(string(sHtml))
}