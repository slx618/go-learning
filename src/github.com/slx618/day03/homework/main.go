package main

import "fmt"

var (
	coins = 100
	users = []string{
		"Mettew", "Sarah", "Augustus", "Heidi", "Elizabeth",
		"Emilie", "Peter", "Giana", "Adriano", "Araon",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	//分金币

	var used int
	for _, v := range users {
		for _, vv := range v {
			switch string(vv) {
			case "e", "E":
				addCoin(distribution, v, 1, &used)
			case "i", "I":
				addCoin(distribution, v, 2, &used)
			case "o", "O":
				addCoin(distribution, v, 3, &used)
			case "u", "U":
				addCoin(distribution, v, 4, &used)

			}

		}
	}
	fmt.Println(distribution, coins-used)
}

func addCoin(dist map[string]int, name string, num int, used *int) {
	if _, ok := dist[name]; ok {
		dist[name] += num
	} else {
		dist[name] = num
	}
	*used += num
}
