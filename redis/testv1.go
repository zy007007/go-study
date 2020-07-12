package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var r redis.Conn

func StartRedis() {
	r, _ = redis.Dial("tcp", "127.0.0.1:6379")
	defer r.Close()
}

func AddtoRedis(){
	_, err := r.Do("SET", "mykey", "superWang321")
    	if err != nil {
        	fmt.Println("redis set failed:", err)
    	}
}

func main() {
    r, _ = redis.Dial("tcp", "127.0.0.1:6379")
    defer r.Close()
   // StartRedis()
    AddtoRedis()
    username, err := redis.String(r.Do("GET", "mykey"))
    if err != nil {
        fmt.Println("redis get failed:", err)
    } else {
        fmt.Printf("Get mykey: %v \n", username)
    }

	
}
