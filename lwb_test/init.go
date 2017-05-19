package lwb_test

//文件处理类库（类库）
//版权所有 2017 文搏，并保留所有权利。
//网站地址: http://www.widerwill.com；
//$Author: liuwenbohhh $
//$Id: Lwb_file 17155 2017-02-06 06:29:05Z $
//例子：  $img = $_POST['img'];
//import (
//	"test/lwb_test"
//)
//con := lwb_test.Init()
//con.Print()

import (
	"fmt"
)

type conn struct {
}

/**
    *初始化结构体
	* @return  *conn     返回结构体
*/
func Init() *conn {
	return new(conn)
}

/**
    *输出
	* @main  origData  原始数据
	* @return  *conn     返回结构体
*/
func (d *conn) Print() error {
	fmt.Println("成功")
	return nil

}
func Testtype(data interface{}) {
	switch t := data.(type) {
	case string:
		fmt.Println(t)
	case []byte:
		fmt.Println(t)
	}
}
