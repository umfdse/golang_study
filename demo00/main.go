package main

import "fmt"

type ps1 struct {
	Name string
	Age  int8
}
type ps2 struct {
	Name string
	Age  int8
}
type ps3 struct {
	Name string
	Age  int8
	Job  string
}

func main() {
	// File, err := os.Open("./main.go") //以只读方式打开
	// File1, err := os.OpenFile("file/log.log", os.O_RDWR, 0666)
	// if err != nil {
	// 	fmt.Println("文件打开失败")
	// 	return
	// }
	// defer File.Close()
	// defer File1.Close()
	// var buf []byte = make([]byte, 1024)
	// for {
	// 	buf = make([]byte, 1024)
	// 	_, err := File.Read(buf)
	// 	fmt.Println(string(buf))
	// 	File1.Write(buf)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			fmt.Println("文件读到末尾了。。。")
	// 			break
	// 		} else {
	// 			fmt.Println("文件读取失败。。。。。。。。。。。。。")
	// 			break
	// 		}
	// 	}
	// }
	var p1 struct{} = struct{}{}
	var p2 struct{} = struct{}{}
	fmt.Println(p1 == p2) //任意空结构体类型相等
	fmt.Println(fmt.Sprintf("%T", p1) == fmt.Sprintf("%T", p2))
	var o1 ps1 = ps1{Name: "张三", Age: 18}
	var o2 ps2 = ps2{Name: "张三", Age: 18}
	var o3 ps3 = ps3{Name: "张三", Age: 18}
	if (fmt.Sprintf("%T", o1) == fmt.Sprintf("%T", o2)) && (o1 == ps1(o2)) {
		fmt.Println("o1和o2两数据相等。。。。")
	} else {
		fmt.Println("o1和o2两数据不相等。。。。")
	}

	if fmt.Sprintf("%T", o1) == fmt.Sprintf("%T", o3) {
		fmt.Println("两数据类型相等。。。。")
	} else {
		fmt.Println("两数据类型不相等。。。。")
	}
	var str string = fmt.Sprintf("%+v", o3)
	fmt.Println(str)
}
