package lwb_time

//时间处理类库（类库）
//版权所有 2017 文搏，并保留所有权利。
//网站地址: http://www.widerwill.com；
//$Author: liuwenbohhh $
//$Id: Lwb_rsa 17155 2017-02-06 06:29:05Z $
import (
	//	"fmt"
	"time"
)

/**
 	* @access  stsr 要格式化的时间格式  "1930-01-01"
	* @access  layout  原始数据   "2006-01-02"
	* @return  int16    时间戳
	* @return  error     是否有错误
*/
func Getunix(tstr, layout string) int64 {
	toBeCharge := tstr //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	//	timeLayout := "2006-01-02" //转化所需模板
	timeLayout := ""
	if layout == "" {
		timeLayout = "2006-01-02"
	} else {
		timeLayout = layout
	}
	//	fmt.Println(timeLayout)
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	/*fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println(sr)     */ //打印输出时间戳 1420041600
	return sr
	//时间戳转日期
	//	dataTimeStr := time.Unix(sr, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	//	fmt.Println(dataTimeStr)
}

/**
 	* @access  intunix 要格式化的时间戳
	* @access  layout  原始数据   "2006-01-02"
	* @return  int16    格式化好的时间
	* @return  error     是否有错误
*/
func Unixtostr(intunix int64, layout string) string {
	sr := intunix
	timeLayout := ""
	if layout == "" {
		timeLayout = "2006-01-02"
	} else {
		timeLayout = layout
	}
	dataTimeStr := time.Unix(sr, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	return dataTimeStr
}
