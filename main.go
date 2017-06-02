package main

import (
	"fmt"
	"test/lwb_config"
)

func main() {
	myConfig := new(lwb_config.Config)
	myConfig.InitConfig("Test.txt")
	fmt.Println(myConfig.Readone("path", "err", "ere"))
	fmt.Printf("%v", myConfig.Mymap)
}
