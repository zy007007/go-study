package main

// 完成对文件的读取和处理，获得目标数据并写入文件

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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
// 过程中遇到了源文件空格在go语言这里" "分割时有点问题
// 在源文件基础上面使用shell产生v1解决
func cleanData(line []string) map[string]int {
	dict := make(map[string]int)

	for _, v := range line {

		arr := strings.Split(v, "-")
		value, _ := strconv.Atoi(arr[1])
		dict[arr[0]] = value
	}

	return dict
}

// 用并发wg改进
func cleanDataGoroutine(line []string) map[string]int {
	dict := make(map[string]int)

	var wg sync.WaitGroup

	var mutex sync.Mutex

	for _, v := range line {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()

			mutex.Lock()
			{
				arr := strings.Split(l, "-")
				value, _ := strconv.Atoi(arr[1])
				dict[arr[0]] = value
			}
			mutex.Unlock()

		}(v)
	}
	wg.Wait()

	return dict
}

// 用通道chan改进
func cleanDataChan(line []string) map[string]int {

	return nil
}

// 求每分钟的和
func sumData(data map[string]int) map[string]int {
	res := make(map[string]int)

	for v, k := range data {
		timeData := strings.Split(v, ".")[0]

		if _, ok := res[timeData]; ok {
			res[timeData] += k
		} else {
			res[timeData] = k
		}
	}

	return res
}

// 结果写入文件
func writeFile(data map[string]int, filename string) error {
	file, _ := os.Create(filename)

	w := bufio.NewWriter(file)

	for k, v := range data {
		str := fmt.Sprintf("%s %d", k, v)
		//n, _ := w.WriteString(str)
		//fmt.Println("写入", n, "一条成功")
		fmt.Fprintln(w, str)
	}
	return w.Flush()
}

func main() {
	data := readFile("exampleDataV1")

	//dict := cleanData(data)
	dict := cleanDataGoroutine(data)

	res := sumData(dict)

	writeFile(res, "result")
}
