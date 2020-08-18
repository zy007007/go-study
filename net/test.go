package main

import (
     "net"
     "fmt"
)

func main() {
   
 //   ipv4 := "12.56.30.1/32"
    ipv6 := "2401:1d40:0:3:0:0:0001::"
    // ParseIP 这个方法 可以用来检查 ip 地址是否正确，如果不正确，该方法返回 nil
    address := net.ParseIP(ipv6)  
    if address == nil {
         fmt.Println("ip地址格式不正确")
    }else {
         fmt.Println("正确的ip地址", address.String()) 
    }
     
}
