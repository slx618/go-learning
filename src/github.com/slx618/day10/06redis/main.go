package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var rdb *redis.Client

func initRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()
	_, err = rdb.Ping(ctx).Result()
	return
}

func getSet() {
	ctx := context.Background()
	err := rdb.Set(ctx, "key", "value", time.Second*5).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println(err)
	}

	if err == redis.Nil {
		fmt.Println("empty key")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(val)
	}
}

func zSet() {
	ctx := context.Background()
	zsetKey := "lang_rank"
	lang := []*redis.Z{
		&redis.Z{Score: 100.1, Member: "Golang"},
		&redis.Z{Score: 200.1, Member: "Php"},
		&redis.Z{Score: 30.1, Member: "Python"},
		&redis.Z{Score: 40.1, Member: "Java"},
		&redis.Z{Score: 990.1, Member: "C/C++"},
		&redis.Z{Score: 10.1, Member: "C#"},
	}

	num, err := rdb.ZAdd(ctx, zsetKey, lang...).Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(num)

	score, err := rdb.ZIncrBy(ctx, zsetKey, 20, "Python").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(score)

	rst, err := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range rst {
		fmt.Println(v.Member, v.Score)
	}
}

// go get -u github.com/go-redis/redis
func main() {
	err := initRedis()
	if err != nil {
		fmt.Println(err)
	}

	zSet()
}
