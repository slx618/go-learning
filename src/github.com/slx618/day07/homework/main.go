package main

import "time"

func main() {
	log := NewLogger("test.log", "info")
	for i := 1; i < 10000; i++ {
		log.Info("序号:%d", i)
		log.Debug("ssss")

		m1 := make(map[string]interface{}, 20)
		m1["sdsa"] = []string{"232", "323232"}
		m1["sdf"] = []interface{}{1212, "12342", []int{1323, 3, 2, 32, 3, 2, 32, 222}}
		log.Info("%#v", m1)
		time.Sleep(time.Millisecond * 100)
	}
	log.Close()
}
