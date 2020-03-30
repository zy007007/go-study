package main

// 完成对文件的读取和处理，获得目标数据并写入文件

import (
	"bufio"
	"fmt"
	"os"
)

// 读取文件内容，放入切片
// 采取了按行读写入的方法，这样子切片长度和容量到最后岂不是超大
func readFile(filename string) []string {
	var res []string

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(filename + " is not exist")
		return nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res
}

// 清洗数据，只保留源数据的时间点和对应数值，返回映射
func cleanData(line []string) map[string]int {
	return nil
}

// 求每分钟的和
func sumData(data map[string]int) map[string]int {
	return nil
}

// 结果写入文件
func writeFile(data map[string]int, filename string) bool {
	return false
}

func main() {
	data := readFile("exampleData1")

}
