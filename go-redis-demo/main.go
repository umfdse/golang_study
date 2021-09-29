package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()

	// 写数据
	_, err := conn.Do("Set", "name", "admin123")
	if err != nil {
		fmt.Println("redis.write err=", err)
		return
	}

	// 读数据
	data, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("redis.read err=", err)
		return
	}
	fmt.Println(data)
}
