package lwb_rand

/**
 *rand处理类库（类库）
 * ============================================================================
 * 版权所有 2017 文搏，并保留所有权利。
 * 网站地址: http://www.widerwill.com；
 * ----------------------------------------------------------------------------
 * ============================================================================
 * $Author: liuwenbohhh $
 * $Id: Lwb_rand 17155 2017-02-06 06:29:05Z $
例子：	bs0 := RandomCreateBytes(16)
	bs1 := RandomCreateBytes(16)

	t.Log(string(bs0), string(bs1))
	if string(bs0) == string(bs1) {
		t.FailNow()
	}

	bs0 = RandomCreateBytes(4, []byte(`a`)...)

	if string(bs0) != "aaaa" {
		t.FailNow()
	}
*/
import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNum = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)

/**
    *得到随机字符串
 	* @access  n    字符串长度
	* @access  alphabets  用来加密的字符串
	* @return  []byte    随机字符串
*/
// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNum
	}
	var bytes = make([]byte, n)
	var randBy bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randBy = true
	}
	for i, b := range bytes {
		if randBy {
			bytes[i] = alphabets[r.Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return bytes
}

/**
    *得到随机数字
 	* @access  n    随机数字最大值
	* @return       随机数字
*/
func Getrandom(maxnum int) int {
	r.Seed(time.Now().UnixNano())
	return r.Intn(maxnum)
}
