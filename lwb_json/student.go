package lwb_json

//json处理类库（类库）
//版权所有 2017 文搏，并保留所有权利。
//网站地址: http://www.widerwill.com；
//$Author: liuwenbohhh $
//$Id: Lwb_json 17155 2017-02-06 06:29:05Z $
//import (
//	"encoding/json"
//	"fmt"
//	"test/lwb_json"
//)
//con := lwb_test.Init()
//con.Print()
//st := &lwb_json.Student{
//	"Xiao Ming",
//	16,
//	true,
//	[]string{"Math", "English", "Chinese"},
//	9.99,
//}
//fmt.Println("before JSON encoding :")
//st.ShowStu()

//b, err := json.Marshal(st)
//if err != nil {
//	fmt.Println("encoding faild")
//} else {
//	fmt.Println("encoded data : ")
//	fmt.Println(b)
//	fmt.Println(string(b))
//}
//ch := make(chan string, 1)
//go func(c chan string, str string) {
//	c <- str
//}(ch, string(b))
//strData := <-ch
//fmt.Println("--------------------------------")
//stb := &lwb_json.Student{}
//stb.ShowStu()
//err = json.Unmarshal([]byte(strData), &stb)
//if err != nil {
//	fmt.Println("Unmarshal faild")
//} else {
//	fmt.Println("Unmarshal success")
//	stb.ShowStu()
//}
import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
}

/**
 * 初始化student对象
 * @return  student对象
 */
func Initstudent() *Student {
	return new(Student)
}

/**
 * 转成json对象
 * @access  student对象
 * @return  转好的json参数
 * @return  错误
 */
func Encode(s *Student) (string, error) {
	b, err := json.Marshal(s)
	return string(b), err

}

/**
 * 转成json对象
 * @access  student对象
 * @return  转好的json参数
 * @return  错误
 */
func Decode(s string) (*Student, error) {
	student := Initstudent()
	err := json.Unmarshal([]byte(s), student)
	return student, err

}

/**
 * print student
 */
func (s *Student) ShowStu() {
	fmt.Println("show Student :")
	fmt.Println("\tName\t:", s.Name)
	fmt.Println("\tAge\t:", s.Age)
	fmt.Println("\tGuake\t:", s.Guake)
	fmt.Println("\tPrice\t:", s.Price)
	fmt.Printf("\tClasses\t: ")
	for _, a := range s.Classes {
		fmt.Printf("%s ", a)
	}
	fmt.Println("")
}
