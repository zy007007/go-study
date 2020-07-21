package main

//今天有道题，是对字符串的操作
//我本意是要用字符数组，结果好像[3]string这样子直接还是字符串数组
//下面对go的字符串和字符操作，进行一个梳理和总结

import (
	"fmt"
)

func main() {
	// 字符串转字符数组
	s := "yiyiwen"

	sarr := []byte(s)

	fmt.Println("字符串:", s)
	fmt.Printf("转字符数组:")
	fmt.Println(sarr) //字符数组:[121 105 121 105 119 101 110]， 发现转换为了对应字符的ASCII码

	fmt.Println("获得单个字符数组arr[1]:", sarr[1], string(sarr[1])) // 得单个字符数组arr[1]: 105 i

	// 直接遍历字符串
	fmt.Println("直接遍历字符串")
	for i, v := range s {
		fmt.Println(i, v)
		//0 121
		//1 105
		//2 121
		//3 105
		//4 119
		//5 101
		//6 110
		// 可以看见同样是ASCII码，加string转单个字符
	}

	// 字符数组转字符串
	again := string(sarr)
	fmt.Println("字符数组转字符串，[]byte直接加string:", again)

	fmt.Println()
	// 字符串中含有中文，用rune处理，附带byte对比
	z := "疑问hhhh"
	zarr := []rune(z)
	fmt.Println("含有中文字符串:", z)
	fmt.Printf("[]rune()转:")
	fmt.Println(zarr, zarr[1], string(zarr[1]), string(zarr[5])) // [30097 38382 104 104 104 104] 38382 问 h

	fmt.Println("byte对比:")
	zbarr := []byte(z)
	fmt.Println("含有中文字符串:", z)
	fmt.Printf("[]byte()转:")
	fmt.Println(zbarr, zbarr[1], string(zbarr[1]), string(zbarr[5]), string(zbarr[7])) // [231 150 145 233 151 174 104 104 104 104] 150  ® h

	// rune byte 对中文转换后
	// byte 一个中文占了3个字节

	// 疑问 在unicode中的编码为30097,38382 因此使用rune可以直接获得对应的unicode，也是因为rune实际类型为int32所赐

	// byte的实际类型为uint8
	// string "疑问hhh" 这个字符串的底层数组是通过 byte数组实现的，中文字符在utf-8编码下占3个字节，而go默认编码为uft-8，因此在用byte转字符时，可以看到3位对应一个汉字

	// 所以目前需要知道的是
	// 有中文存在的情况下需要字符串转字符，使用[]rune()

}
