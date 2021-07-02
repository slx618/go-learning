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

	req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1:9090/x?"+u.Encode(), nil)

	rsp, err := http.Get(urlObj.String())
	if err != nil {
		fmt.Println(err)

	}
	data, _ := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(data))
}
