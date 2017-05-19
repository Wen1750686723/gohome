package main

import (
	"fmt"
	"test/lwb_json"
)

func main() {
	st := &lwb_json.Student{
		"Xiao Ming",
		16,
		true,
		[]string{"Math", "English", "Chinese"},
		9.99,
	}
	student, _ := lwb_json.Encode(st)
	fmt.Println(student)
	sts, _ := lwb_json.Decode(student)
	sts.ShowStu()
}
