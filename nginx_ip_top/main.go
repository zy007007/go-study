package main

import (
	"fmt"
	"github.com/hpcloud/tail"
)

type Visitip struct {
	Ip 	string
	Area 	string
}

func AddtoRedis(data *Visitip) (bool, error) {

}


func HandleLog(data string) (*Visitip, error) {
	arr := string.Split(data, " ")

	area, err := IpArea(arr[0])

	vi := &Visitip{
		Ip:arr[0],
		Area:area
	}

	return vi, nil
}

func IpArea(ip string) (string, error) {
	var res string

	cmd := exec.Command("/bin/bash", "-c", "curl -s --user-agent foobar https://ip.cn/index.php?ip="+ip+" | grep '所在地理位置'| awk -F '>' '{print $9}' | awk -F '<' '{print $1}'")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return
	}

	outputBuf := bufio.NewReader(stdout)

	for {

		output, _, err := outputBuf.ReadLine()
		if err != nil {

			if err.Error() != "EOF" {
				fmt.Printf("Error :%s\n", err)
			}
			return
		}
		res = string(output)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return
	}

}

func main() {
	t, err := tail.TailFile("/var/log/nginx/access.log", tail.Config{Follow: true})
	if err != nil {
		fmt.Println(err)
	}

	for line := range t.Lines {
		fmt.Println(line.Text)
		data := HandleLog(line.Text)
		res, err := AddtoRedis(data)
		if err != nil {
			fmt.Println("add to redis failed", err.Error())
		}
	}
}




// 然后在博客系统中，加入从redis读取排行数据并展示的功能

// 还要研究一下，hpcloud/tail如何从文件末尾开始读取
