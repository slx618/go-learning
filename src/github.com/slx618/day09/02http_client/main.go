package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	u := url.Values{} //urlencode
	u.Set("name", "xxe==23?232/%x")
	u.Set("age", "18")

	urlObj, _ := url.Parse("http://127.0.0.1:9090/x")
	urlObj.RawQuery = u.Encode()

	req, _ := http.NewRequest(http.MethodGet, urlObj.String(), nil)
	//rsp, err := http.DefaultClient.Do(req)

	//添加头
	req.Header.Add("xxx", "xxxx")
	var client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false, //禁用长连接
		},
	}
	rsp, err := client.Do(req)
	//rsp, err := http.Get("http://127.0.0.1:9090/x")

	if err != nil {
		fmt.Println(err)

	}

	data, _ := ioutil.ReadAll(rsp.Body)
	rsp.Body.Close()
	fmt.Println(string(data))
}
