package lwb_config

import (
	"fmt"
	"testing"
)

func TestRsaEncrypt(t *testing.T) {
	myConfig := new(Config)
	myConfig.InitConfig("Test.txt")
	fmt.Println(myConfig.Read("default", "path"))
	fmt.Printf("%v", myConfig.Mymap)
}
