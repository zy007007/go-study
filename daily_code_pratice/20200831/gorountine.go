





func main() {
	tmp := make(chan int)
	ch2 := make(chan interface{})

	ch2 <- "0"
	ch2 <- "zhaoyi"


	go task1()

	go task2()
}

func task1() {
	for i:=1;i<10;i++{
		http.Get("www.baidu.com")
	}
}




func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("start gorountine")

		ch <- 0

		fmt.Println("end gorountine")
	}()

	fmt.Println("wait gorountine")

	<-ch

	fmt.Println("all down")
}



func main() {

	ch := make(chan int)

	for i:=0; i<3; i++ {
		fmt.Println("chan add data")
		ch <- i
		time.Sleep(time.Second)
	}

	for data := range ch {
		fmt.Println(data)

		if data==2 {
			break
		}
	}
}




func main() {
	go func() {

	}
}


package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//go的并发
	timestart := time.Now().UnixNano()
	isstatus1 := new(bool)
	isstatus2 := new(bool)
	*isstatus1 = false
	*isstatus2 = false

	go task1(isstatus1)
	go task2(isstatus2)

	for{
		if *isstatus1 && *isstatus2{
			fmt.Println("执行任务完毕用时：",float64(time.Now().UnixNano() - timestart)/1000000000)
			break
		}
		time.Sleep(time.Millisecond)
	}
}

func task1(isstatus1 *bool){
	for i:=0;i<100;i++{
		fmt.Println("。。。。正在执行任务一。。。。")
		client := &http.Client{}
		req,err := http.NewRequest("GET","https://www.baidu.com/",nil)
		if err != nil{
			fmt.Println("NewRequest err:",err)
			continue
		}
		reqponse,_ := client.Do(req)
		fmt.Println("任务一状态:",reqponse.StatusCode)
	}
	fmt.Println("。。。。END 任务一。。。。")
	*isstatus1 = true
}


client := &http.Client{}
req, err := http.NewRequest()
client.Do()


func task2(isstatus2 *bool){
	for i:=0;i<100;i++ {
		fmt.Println("。。。。正在执行任务二。。。。")
		client := &http.Client{}
		req,err := http.NewRequest("GET","https://www.so.com/",nil)
		if err != nil{
			fmt.Println("NewRequest err:",err)
			continue
		}
		reqponse,_ := client.Do(req)
		fmt.Println("任务二状态:",reqponse.StatusCode)
	}
	fmt.Println("。。。。END 任务二。。。。")
	*isstatus2 = true
}