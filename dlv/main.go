package main

// 简单的试试了dlv调试工具，后续这个可以搞起来哈
// 后面开始阅读其他源码或者项目时，可以操作一下

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const port = "8000"

func main() {
	http.HandleFunc("/hi", hi)

	fmt.Println("runing on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func hi(w http.ResponseWriter, r *http.Request) {
	hostName, _ := os.Hostname()
	fmt.Fprintf(w, "HostName: %s", hostName)
}

//dlv debug ./main.go
//Type 'help' for list of commands.
//(dlv) b main.main
//Breakpoint 1 set at 0x7590d8 for main.main() ./main.go:12
//(dlv) c
//> main.main() ./main.go:12 (hits goroutine(1):1 total:1) (PC: 0x7590d8)
//     7:		"os"
//     8:	)
//     9:
//    10:	const port = "8000"
//    11:
//=>  12:	func main() {
//    13:		http.HandleFunc("/hi", hi)
//    14:
//    15:		fmt.Println("runing on port: " + port)
//    16:		log.Fatal(http.ListenAndServe(":"+port, nil))
//    17:	}
//(dlv) b main.hi
//Breakpoint 2 set at 0x759278 for main.hi() ./main.go:19
//(dlv) c
//runing on port: 8000
//> main.hi() ./main.go:19 (hits goroutine(6):1 total:1) (PC: 0x759278)
//    14:
//    15:		fmt.Println("runing on port: " + port)
//    16:		log.Fatal(http.ListenAndServe(":"+port, nil))
//    17:	}
//    18:
//=>  19:	func hi(w http.ResponseWriter, r *http.Request) {
//    20:		hostName, _ := os.Hostname()
//    21:		fmt.Fprintf(w, "HostName: %s", hostName)
//    22:	}
//(dlv) n
//> main.hi() ./main.go:20 (PC: 0x75928f)
//    15:		fmt.Println("runing on port: " + port)
//    16:		log.Fatal(http.ListenAndServe(":"+port, nil))
//    17:	}
//    18:
//    19:	func hi(w http.ResponseWriter, r *http.Request) {
//=>  20:		hostName, _ := os.Hostname()
//    21:		fmt.Fprintf(w, "HostName: %s", hostName)
//    22:	}
//(dlv) p hostName
//""
//(dlv) n
//> main.hi() ./main.go:21 (PC: 0x7592c2)
//    16:		log.Fatal(http.ListenAndServe(":"+port, nil))
//    17:	}
//    18:
//    19:	func hi(w http.ResponseWriter, r *http.Request) {
//    20:		hostName, _ := os.Hostname()
//=>  21:		fmt.Fprintf(w, "HostName: %s", hostName)
//    22:	}
//(dlv) p hostName
//"yiyiwen"
//(dlv) locals
//hostName = "yiyiwen"
//(dlv) args
//w = net/http.ResponseWriter(*net/http.response) 0xc000099850
//r = ("*net/http.Request")(0xc0000f6000)
