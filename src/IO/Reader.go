package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// 操作Reader对象

func main() {

	/*
		Read方法
	*/
	data1 := []byte("凌之轩家的小贝尔")
	rd1 := bytes.NewReader(data1)
	r1 := bufio.NewReader(rd1)
	var buf1 [128]byte

	// 读取数据并存放到字节切片buf中去
	n1, err := r1.Read(buf1[:])
	fmt.Println(string(buf1[:n1]), n1, err)

	/*
		ReadByte()方法
	*/
	data2 := []byte("_凌之轩家的小贝尔")
	rd2 := bytes.NewReader(data2)
	r2 := bufio.NewReader(rd2)

	// 读取并返回一个字节
	c, err := r2.ReadByte()
	fmt.Println(string(c), err)

	/*
		ReadBytes()方法
	*/
	data3 := []byte("凌之轩家的小贝尔, 其实他以前叫神乐")
	rd3 := bytes.NewReader(data3)
	r3 := bufio.NewReader(rd3)

	// 设置分隔符
	var delim byte = ','
	// 读取数据直到遇到第一个分隔符“delim”，返回读取的字节序列（包括“delim”）
	line, err := r3.ReadBytes(delim)
	fmt.Println(string(line), err)

	/*
		ReadLine()方法
	*/
	data4 := []byte("凌之轩家的小贝尔\r\n其实他以前叫神乐")
	rd4 := bytes.NewReader(data4)
	r4 := bufio.NewReader(rd4)

	// 读取一行数据
	line, prefix, err := r4.ReadLine()
	fmt.Println(string(line), prefix, err)

	// 等价于如下方式
	var delimNew byte = '\n'
	lineNew, err := r4.ReadBytes(delimNew)
	fmt.Println(string(lineNew), err)

	/*
		ReadRune()方法
	*/
	data5 := []byte("凌之轩家的小贝尔")
	rd5 := bytes.NewReader(data5)
	r5 := bufio.NewReader(rd5)

	// 读取一个 UTF-8 编码的字符，并返回其 Unicode 编码和字节数
	ch, size, err := r5.ReadRune()
	fmt.Println(string(ch), size, err)

	/*
		ReadSlice()方法
	*/
	data6 := []byte("凌之轩家的小贝尔,其实他以前叫神乐")
	rd6 := bytes.NewReader(data6)
	r6 := bufio.NewReader(rd6)
	var delim6 byte = ','

	// 读取数据直到分隔符“delim”处，并返回读取数据的字节切片
	line61, err := r6.ReadSlice(delim6)
	fmt.Println(string(line61), err)
	line62, err := r6.ReadSlice(delim6)
	fmt.Println(string(line62), err)
	line63, err := r6.ReadSlice(delim6)
	fmt.Println(string(line63), err)

	/*
		ReadString()方法
	*/
	data7 := []byte("凌之轩家的小贝尔,其实他以前叫神乐")
	rd7 := bytes.NewReader(data7)
	r7 := bufio.NewReader(rd7)
	var delim7 byte = ','

	// 读取数据直到分隔符“delim”第一次出现，并返回一个包含“delim”的字符串
	line7, err := r7.ReadString(delim7)
	fmt.Println(string(line7), err)

	/*
		Buffered()方法
	*/
	data8 := []byte("凌之轩家的小贝尔")
	rd8 := bytes.NewReader(data8)
	r8 := bufio.NewReader(rd8)
	var buf8 [15]byte
	n8, err := r8.Read(buf8[:])
	fmt.Println(string(buf8[:n8]), n8, err)

	// 返回可从缓冲区读出数据的字节数
	rn8 := r8.Buffered()
	fmt.Println(rn8)

	n8, err = rd8.Read(buf8[:])
	fmt.Println(string(buf8[:n8]), n8, err)
	rn8 = r8.Buffered()
	fmt.Println(rn8)

	/*
		peek()方法
	*/
	data9 := []byte("凌之轩家的小贝尔")
	rn9 := bytes.NewReader(data9)
	r9 := bufio.NewReader(rn9)
	bl, err := r9.Peek(9)
	fmt.Println(string(bl), err)
	bl, err = r9.Peek(12)
	fmt.Println(string(bl), err)
	bl, err = r9.Peek(24)
	fmt.Println(string(bl), err)
	bl, err = r9.Peek(25)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bl), err)
}
