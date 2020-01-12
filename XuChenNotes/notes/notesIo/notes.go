package notesIo

import "fmt"

// WriteString writes the contents of the string s to w, which accepts a slice of bytes.
// If w implements StringWriter, its WriteString method is invoked directly.
// Otherwise, w.Write is called exactly once.

type Writer interface {
	Write(p []byte) (n int, err error)
}

type StringWriter interface {
	WriteString(s string) (n int, err error)
}

//WriteString 函数的传入参数列表中 w 是一个实现io.Writer接口的具体对象
//如果w 实现了StringWriter，那么则直接采用w对象中的WriteString函数
func WriteString(w Writer, s string) (n int, err error) {
	if sw, ok := w.(StringWriter); ok {
		fmt.Println("1")
		return sw.WriteString(s)
	}
	fmt.Println("2")
	return w.Write([]byte(s))
}

// func Test() {
// 	fmt.Println("sadsa")
// }
