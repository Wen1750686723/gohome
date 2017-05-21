package main

import (
	"fmt"
	"test/lwb_encode"
)

func main() {
	fmt.Println(lwb_encode.Base64Encode("ss"))
	fmt.Println(lwb_encode.Base64Decode(lwb_encode.Base64Encode("ss")))
}
