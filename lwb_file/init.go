package lwb_file

/**
 * file处理类库（类库）
 * ========================================================================
 * 版权所有 2017 文搏，并保留所有权利。
 * 网站地址: http://www.widerwill.com；
 * ------------------------------------------------------------------------
 * ========================================================================
 * $Author: liuwenbohhh $
 * $Id: Lwb_encode 17155 2017-02-06 06:29:05Z $
 */
import (
	"io"
	"io/ioutil"
	"os"
)

/**
    * 从文件中取数据
 	* @access  filename  文件地址
	* @access  data      数据
	* @access  perm      文件权限
	* @return  error     是否有错误
*/
func File_get_content(filepath string) string {
	ra, error := ioutil.ReadFile(filepath)
	if error != nil {
		return ""
	}
	return string(ra)
}

/**
    *向文件中加数据没有则新建
 	* @access  filename  文件地址
	* @access  data      数据
	* @access  perm      文件权限
	* @return  error     是否有错误
*/
func File_put_content(filename string, data []byte) error {

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
	// Hello World!
}
