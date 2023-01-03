package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Addr     string `ini:"addr"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	IsTest   bool   `ini:"isTest"`
}

type RedisConfig struct {
	Host     string  `ini:"host"`
	Port     int     `ini:"port"`
	Password string  `ini:"password"`
	Database int     `ini:"database"`
	Timeout  float64 `ini:"timeout"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
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

	strArr, err := readFile()

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
				err = fmt.Errorf("第%d行配置文件语法错误", idx+1)
				return
			}

			//section: mysql redis
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
				err = fmt.Errorf("第%d行配置文件语法错误", idx+1)
				return
			}

			sp := strings.Split(v, "=")
			if len(sp) < 1 { //切割成功
				err = fmt.Errorf("第%d行配置文件语法错误", idx+1)
			}
			key := strings.TrimSpace(v[0:strings.Index(v, "=")])
			value := strings.TrimSpace(v[strings.Index(v, "=")+1:])

			rv := reflect.ValueOf(s)
			v1 := rv.Elem().FieldByName(structName)
			if v1.Kind() != reflect.Struct {
				err = fmt.Errorf("%s不是一个结构体", structName)
				return
			}
			t1 := v1.Type()
			/**
			  field.tag -> addr
			  field.tag -> port
				...
			*/
			var tagName string
			for i := 0; i < v1.NumField(); i++ {
				fmt.Println(t1.Field(i).Tag.Get("ini"))
				if t1.Field(i).Tag.Get("ini") == key { // 这里比较的是 tag上 和配置文件上的key
					tagName = t1.Field(i).Name //这里那的是 结构体里的key
					break
				}
			}
			if tagName == "" {
				continue
			}
			vField := v1.FieldByName(tagName)
			switch vField.Type().Kind() {
			case reflect.String:
				vField.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				vInt, _ := strconv.ParseInt(value, 10, 64)
				vField.SetInt(vInt)
			case reflect.Bool:
				vBool, _ := strconv.ParseBool(value)
				vField.SetBool(vBool)
			case reflect.Float32, reflect.Float64:
				vFloat, _ := strconv.ParseFloat(value, 64)
				vField.SetFloat(vFloat)
			}

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

func readFile() ([]string, error) {
	//读取
	fd, err := os.Open("./src/github.com/slx618/day06/homework/conf.ini")
	if err != nil {
		return nil, err
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
	err = fd.Close()

	//用\n切割字符串
	return strings.Split(string(content), "\n"), nil
}

func main() {
	var config Config
	//var redisIni redisConfig
	err := loadIni(&config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", config)
}
