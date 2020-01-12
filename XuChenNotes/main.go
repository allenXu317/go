package main

//import路径
import Io "./notes/notesIo"

//定义了myWriter的自定义类型
type myWriter struct{}

//实现Writer接口
func (w myWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

//实现了StringWriter接口
//WriteString参数列表中定义了w Writer接口型的值
//所以k := myWriter 要能被WriteString传入，必须全部实现Writer接口中的方法
func (w myWriter) WriteString(s string) (n int, err error) {
	return 1, nil
}

func main() {
	k := new(myWriter)
	//进行接口调用
	p := make([]byte, 10)
	//调用myWriter实现的接口
	k.Write(p)
	//调用Io中的函数
	//WriteString函数中还有myWriter实现的WriteString方法
	Io.WriteString(k, "sadds")
}
