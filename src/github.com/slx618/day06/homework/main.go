package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

type mysqlConfig struct {
	Addr     string `ini:"addr"`
	Port     string `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type redisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

type config struct {
	mysqlConfig `ini:"mysql"`
	redisConfig `ini:"redis"`
}

func loadIni(s interface{}) (err error) {

	//检查传进来的 s 是结构体 并且是指针类型
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Ptr {
		err = errors.New("不是指针类型")
		return
	}
	if rt.Elem().Kind() != reflect.Struct {
		err = errors.New("不是结构体")
		return
	}

	//读取
	fd, err := os.Open("./src/github.com/slx618/day06/homework/conf.ini")
	if err != nil {
		return
	}
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
	fd.Close()

	//用\n切割字符串
	strArr := strings.Split(string(content), "\n")

	var structName string
	for idx, v := range strArr {
		//去掉空格
		v = strings.TrimSpace(v)
		//注释就跳过
		if strings.HasPrefix(v, "#") || strings.HasPrefix(v, ";") {
			continue
		}
		//空行
		if len(v) == 0 {
			continue
		}

		// []
		if strings.HasPrefix(v, "[") {
			// [   ] 为空的情况
			if len(strings.TrimSpace(v[1:len(v)-1])) == 0 {
				err = fmt.Errorf("第%d行配置文件语法错误", idx + 1)
				return
			}
			sectionName := strings.TrimSpace(v[1 : len(v)-1])

			for i := 0; i < rt.Elem().NumField(); i++ {
				if sectionName == rt.Elem().Field(i).Tag.Get("ini") {
					structName = rt.Elem().Field(i).Name
					//fmt.Println(sectionName, structName)
					break

				}

			}

		} else {
			//切割等号的配置
			if strings.Index(v, "=") <= 0 {
				err = fmt.Errorf("第%d行配置文件语法错误", idx +1)
				return
			}

			sp := strings.Split(v, "=")
			if len(sp) < 1 { //切割成功
				err = fmt.Errorf("第%d行配置文件语法错误", idx + 1)
			}
			strings.TrimSpace(sp[0])
			strings.TrimSpace(sp[1])
			//eqIdx := strings.Index(v, "=")
			//key := strings.T

			rv := reflect.ValueOf(s)

			v1 := rv.Elem().FieldByName(structName)
			if v1.Kind() != reflect.Struct {
				err = fmt.Errorf("%s不是一个结构体", structName)
				return

			}
			t1, ok := rt.FieldByName(structName)
			if !ok {
				err = fmt.Errorf("%s不是一个结构体", structName)
				return
			}
			t1.

//strconv.ParseInt()
//			v2 := v1.Elem().FieldByName(toUpCase(sp[0]))
			//v2.Set(reflect.ValueOf(v))

			//fmt.Println(configMap)
			//for k, v := range configMap {
			//	//首字母大写
			//	fmt.Println(toUpCase(k))
			//	fmt.Println(v)
			//	field := rv.Elem().FieldByName(toUpCase(k))
			//	field.Set(reflect.ValueOf(v)) //这里有问题 其他类型 int怎么办
			//}

		}

	}

	return
}

func toUpCase(s string) string {
	strArr := strings.Split(s, "")
	sp := strArr[:]
	sp[0] = strings.ToUpper(sp[0])
	return strings.Join(sp[:], "")
}

func main() {
	var config config
	//var redisIni redisConfig
	err := loadIni(&config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", config)
}
