package main

import (
	//	"fmt"
	"test/lwb_directory"
)

func main() {
	files_ch := make(chan *lwb_directory.FileInfo, 100)
	go lwb_directory.WalkFiles("/Users/liuwenbo/Applications/Go/src/myproject", "", files_ch) //在一个独立的 goroutine 中遍历文件
	lwb_directory.WriteFiles("/Users/liuwenbo/Applications/Go/src/myproject1", files_ch)
	//获取本地location

}
