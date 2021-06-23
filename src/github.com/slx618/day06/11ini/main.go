package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

type mcs struct {
	Addr     string `ini:"addr"`
	Port     string `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

func loadIni(s interface{}) {
	fd, err := os.Open("./src/github.com/slx618/day06/11ini/conf.ini")
	if err != nil {
		fmt.Println(err)
	}
	defer fd.Close()

	fdStat, _ := fd.Stat()
	var str [128]byte
	var content = make([]byte, 0, fdStat.Size())
	for {
		n, err := fd.Read(str[:])
		if err == io.EOF {
			//content = append(content, str[:n]...)
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		content = append(content, str[:n]...)
	}

	//用\n切割字符串
	strArr := strings.Split(string(content), "\r\n")
	var configMap = make(map[string]interface{}, len(strArr[1:]))
	for _, v := range strArr {
		sp := strings.Split(v, "=")
		if len(sp) > 1 { //切割成功
			configMap[sp[0]] = sp[1]
		}

	}
	rv := reflect.ValueOf(s)
	//rt := reflect.TypeOf(s)
	fmt.Println(configMap)
	for k, v := range configMap {
		//首字母大写
		fmt.Println(toUpCase(k))
		fmt.Println(v)
		field := rv.Elem().FieldByName(toUpCase(k))
		field.Set(reflect.ValueOf(v)) //这里有问题 其他类型 int怎么办
	}
}

func toUpCase(s string) string {
	strArr := strings.Split(s, "")
	sp := strArr[:]
	sp[0] = strings.ToUpper(sp[0])
	return strings.Join(sp[:], "")
}

func main() {
	var mysqlIni mcs
	loadIni(&mysqlIni)
	fmt.Printf("%#v\n", mysqlIni)
}
