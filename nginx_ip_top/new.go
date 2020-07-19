package main

import (
	"fmt"
	"strings"
	"github.com/hpcloud/tail"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os/exec"
	"bufio"
	"log"
)

type Visit struct {
	Id	int64
	Ip 	string
	Area 	string
	Counts	int64
}

func(this *Visit) TableName() string {
	return "visit"
}

var orm *xorm.Engine

//创建orm引擎
func init() {
	var err error
	orm, err = xorm.NewEngine("mysql", "root:thisislifeZy007~@tcp(127.0.0.1:3306)/blog?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
		fmt.Println("数据库连接失败:", err)
	}
	fmt.Println("[Connecting MySQL Success!]")

}

func HandleMysql(data *Visit) (error) {
	res, err := orm.Exist(&Visit{Ip:data.Ip})
	switch res{
	case true:
		var old Visit
		_, err := orm.Where("ip=?", data.Ip).Get(&old)

		old.Counts += 1
		_, err = orm.Where("ip=?", old.Ip).Update(&old)
		return err
	case false:
		data.Counts = 1
		_, err := orm.Insert(data)
		return err
	}
	return err
}


func HandleLog(data string) (*Visit) {
	arr := strings.Split(data, " ")

	area, err := IpArea(arr[0])
	if err != nil {
		fmt.Println("HandleLog's IpArea failed")
	}

	vi := &Visit{
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



	output, _, err := outputBuf.ReadLine()
	if err != nil {

		if err.Error() != "EOF" {
			fmt.Printf("Error :%s\n", err)
		}
		return "", nil
	}
	res = string(output)


	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return "", nil
	}

	return res, nil
}



func main() {

	t, err := tail.TailFile("/var/log/nginx/access.log", tail.Config{Follow: true})
	if err != nil {
		fmt.Println(err)
	}

	for line := range t.Lines {
		fmt.Println(line.Text)
		data := HandleLog(line.Text)
		fmt.Println(data)
		err := HandleMysql(data)
		if err != nil {
			fmt.Println("handle mysql failed", err.Error())
		}
	}
}
