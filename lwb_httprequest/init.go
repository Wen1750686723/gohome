package lwb_httprequest

//http请求处理类库（类库）
//版权所有 2017 文搏，并保留所有权利。
//网站地址: http://www.widerwill.com；
//$Author: liuwenbohhh $
//$Id: Lwb_file 17155 2017-02-06 06:29:05Z $
//例子：  $img = $_POST['img'];
//import (
//	"test/lwb_httprequest"
//)
//con := lwb_test.Init()
//con.Print()
import (
	"io/ioutil"
	"net/http"
	"strings"
)

func Httpget(url string, data string) (string, error) {
	resp, err := http.Get(url + data)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
func Httppost(url string, data string) (string, error) {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(data))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

//func httpPostForm(url string, data string) (string, error) {
//	resp, err := http.PostForm(url,
//		url.Values{"key": {"Value"}, "id": {"123"}})

//	if err != nil {
//		return "", err
//	}

//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return "", err
//	}

//	return string(body), nil

//}
func Httpdo(url string, contenytype string, header string, cookie string, post string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(post))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", header)
	req.Header.Set("Cookie", cookie)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
