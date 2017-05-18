package main

import (
	"fmt"
	"test/lwb_httprequest"
)

func main() {
	data, _ := lwb_httprequest.Httppost("http://www.baidu.com", "")
	fmt.Print(data)
}
