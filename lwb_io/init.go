package lwb_io

/**
 *io处理类库（类库）
 * ============================================================================
 * 版权所有 2017 文搏，并保留所有权利。
 * 网站地址: http://www.widerwill.com；
 * ----------------------------------------------------------------------------
 * ============================================================================
 * $Author: liuwenbohhh $
 * $Id: Lwb_io 17155 2017-02-06 06:29:05Z $
 */
import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Readfile() {
	ra, _ := ioutil.ReadFile("Test.txt")
	fmt.Printf("%s", ra)
}
func Writefile(mode int) {
	if mode == 1 {
		fn := "Test.txt"
		s := []byte("Hello World!")
		Writefileappend(fn, s, 0777)
		rf, _ := ioutil.ReadFile(fn)
		fmt.Printf("%s", rf)
	} else {
		fn := "Test.txt"
		s := []byte("Hello World!")
		Writefileclear(fn, s, 0777)
		rf, _ := ioutil.ReadFile(fn)
		fmt.Printf("%s", rf)
	}
	// Hello World!
}
func Writefileappend(filename string, data []byte, perm os.FileMode) error {

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)

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
func Writefileclear(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)

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
func Readdir() {
	rd, err := ioutil.ReadDir("")
	for _, fi := range rd {
		fmt.Println("")
		fmt.Println(fi.Name())
		fmt.Println(fi.IsDir())
		fmt.Println(fi.Size())
		fmt.Println(fi.ModTime())
		fmt.Println(fi.Mode())
	}
	fmt.Println("")
	fmt.Println(err)
}
