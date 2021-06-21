package main

import "fmt"

// 在解析中使用自定义错误
func main() {
	var e error
	// 创建一个错误实例, 包含文件名和行号
	e = newParseError("test.go", 99)
	// 通过error接口查看错误实例
	fmt.Println(e.Error())

	// 根据错误接口具体的类型，获取详细错误信息
	switch detail := e.(type) {
	case *ParseError:
		fmt.Printf("解析错误，文件名: %s, 行号: %d\n", detail.Filename, detail.Line)
	default:
		fmt.Println("其他类型错误")
	}

}

// ParseError 声明一个解析错误
type ParseError struct {
	Filename string // 文件号
	Line     int    // 行号
}

// 实现error接口，返回具体错误实例
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s: %d", e.Filename, e.Line)
}

// 创建一些解析错误
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}
