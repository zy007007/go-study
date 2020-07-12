package main
import (
    "fmt"
    "github.com/garyburd/redigo/redis"
)

var c redis.Conn

func AddRedis(){
    _, err := c.Do("SET", "mykey", "superWang666")
    if err != nil {
        fmt.Println("redis set failed:", err)
    }
}

func main() {
    c, err = redis.Dial("tcp", "127.0.0.1:6379")
    fmt.Println(c)
    if  err != nil {
        fmt.Println("Connect to redis error", err)
        return
    }
    defer c.Close()
    
    AddRedis()
    //_, err = c.Do("SET", "mykey", "superWang")
    //if err != nil {
    //    fmt.Println("redis set failed:", err)
    //}

    username, err := redis.String(c.Do("GET", "mykey"))
    if err != nil {
        fmt.Println("redis get failed:", err)
    } else {
        fmt.Printf("Get mykey: %v \n", username)
    }
}
