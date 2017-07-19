package study

/**
 * xml处理类库（类库）
 * ============================================================================
 * 版权所有 2017 文搏，并保留所有权利。
 * 网站地址: http://www.widerwill.com；
 * ----------------------------------------------------------------------------
 * ============================================================================
 * $Author: liuwenbohhh $
 * $Id: Lwb_xml 17155 2017-02-06 06:29:05Z $
 */
import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

type Result struct {
	Person []Person `xml:"person"`
}
type Person struct {
	Name      string    `xml:"name,attr"`
	Age       int       `xml:"age,attr"`
	Career    string    `xml:"career"`
	Interests Interests `xml:"interests"`
}
type Interests struct {
	Interest []string `xml:"interest"`
}

func (person *Person) Chkis18() (flag bool) {
	if person.Age > 18 {
		flag = true
	}
	return flag
}

type Checker interface {
	Chkis18() (flag bool)
}

func Get() {
	content, err := ioutil.ReadFile("lwb_xml/study/study.xml")
	if err != nil {
		log.Fatal(err)
	}
	var result Result
	err = xml.Unmarshal(content, &result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result.Person)
}
