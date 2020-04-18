package main

import (
	"fmt"
	"net"
)

func main() {
	iprecords, _ := net.LookupIP("baidu.com")
	for _, ip := range iprecords {
		fmt.Println(ip)
	}

	cname, _ := net.LookupCNAME("baidu.com")
	fmt.Println(cname)

	ptr, _ := net.LookupAddr("8.8.8.8")
	for _, ptrvalue := range ptr {
		fmt.Println(ptrvalue)
	}

	nameserver, _ := net.LookupNS("baidu.com")
	for _, ns := range nameserver {
		fmt.Println(ns)
	}
}
