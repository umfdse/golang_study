package main

import (
	"fmt"
	"strconv"
)

func main() {
	var f1 float64 = 3.1415926
	var str string = fmt.Sprintf("%.6f", f1) //四舍五入保留6位小数
	fmt.Printf("str的数据类型是%T 值是%v \n", str, str)
	f2, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("数据转化出错啦！！！")
	}
	fmt.Println(f2)
	f3, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", f1), 64) //四舍五入保留2位小数
	fmt.Println(f3)
}
