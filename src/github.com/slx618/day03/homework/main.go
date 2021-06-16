package main

import "fmt"

var (
	coins = 100
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Elizabeth",
		"Emilie", "Peter", "Giana", "Adriano", "Araon",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	//分金币

	var used int
	for _, name := range users {
		for _, letter := range name {
			switch string(letter) {
			case "e", "E":
				addCoin(distribution, name, 1, &used)
			case "i", "I":
				addCoin(distribution, name, 2, &used)
			case "o", "O":
				addCoin(distribution, name, 3, &used)
			case "u", "U":
				addCoin(distribution, name, 4, &used)

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
