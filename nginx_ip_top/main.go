package main

import (
	"fmt"
	"strings"
	"github.com/hpcloud/tail"
	"github.com/gomodule/redigo/redis"
	"os/exec"
	"bufio"
)

type Visitip struct {
	Ip 	string
	Area 	string
}

var r redis.Conn

func AddtoRedis(data *Visitip) (bool, error) {
	key := data.Ip + ":" + data.Area
	fmt.Println(key)
	_, err := redis.Int64(r.Do("INCR", key))
	if err != nil {
		fmt.Println("incr %s failed", key)
		return false, err
	}
	return true, nil
}

func HandleLog(data string) (*Visitip) {
	arr := strings.Split(data, " ")

	area, err := IpArea(arr[0])
	if err != nil {
		fmt.Println("HandleLog's IpArea failed")
	}

	vi := &Visitip{
		Ip:arr[0],
		Area:area,
	}

	return vi
}

func IpArea(ip string) (string, error) {
	var res string

	cmd := exec.Command("/bin/bash", "-c", "curl -s --user-agent foobar https://ip.cn/index.php?ip="+ip+" | grep '所在地理位置'| awk -F '>' '{print $9}' | awk -F '<' '{print $1}'")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return "", nil
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return "", nil
	}

	outputBuf := bufio.NewReader(stdout)

	//for {

		output, _, err := outputBuf.ReadLine()
		if err != nil {

			if err.Error() != "EOF" {
				fmt.Printf("Error :%s\n", err)
			}
			return "", nil
		}
		res = string(output)
	//	fmt.Println(res)
	//}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return "", nil
	}
	//fmt.Println("res", res)
	return res, nil
}

func main() {

   	r, _ = redis.Dial("tcp", "127.0.0.1:6379")
    	defer r.Close()


	//t, err := tail.TailFile("/var/log/nginx/access.log", tail.Config{Follow: true})
	t, err := tail.TailFile("access.log", tail.Config{Follow: true})
	if err != nil {
		fmt.Println(err)
	}

	for line := range t.Lines {
		data := HandleLog(line.Text)
		res, err := AddtoRedis(data)
		if !res {
			fmt.Println("add to redis failed", err.Error())
		}
	}
}




// 然后在博客系统中，加入从redis读取排行数据并展示的功能

// 还要研究一下，hpcloud/tail如何从文件末尾开始读取
