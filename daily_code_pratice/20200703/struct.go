package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// 文件对象
//type File struct {
//	fd int
//}
//
//// 关闭文件
//func (f *File) Close() error {
//	// ...
//}
//
//// 读文件数据
//func (f *File) Read(int64 offset, data []byte) int {
//	// ...
//}
//
//type Point struct{ X, Y float64 }
//
//type ColoredPoint struct {
//	Point
//	Color color.RGBA
//}
//
//type Cache struct {
//	m map[string]string
//	sync.Mutex
//}
//
//func (p *Cache) Lookup(key string) string {
//	p.Lock()
//	defer p.Unlock()
//
//	return p.m[key]
//}

type UpperWriter struct {
	io.Writer
}

func (p *UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}

func main() {
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")
}
