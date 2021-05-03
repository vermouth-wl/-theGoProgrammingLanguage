package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

// 请求结构体
type requestBody struct {
	Key    string `json:"key"`
	Info   string `json:"info"`
	UserId string `json:"userId"`
}

// 结果结构体
type responseBody struct {
	Text string   `json:"text"`
	List []string `json:"list"`
	Url  string   `json:"url"`
}

// 请求机器人
func process(inputChan <-chan string, userid string) {
	for {
		// 从通道中接收输入
		input := <-inputChan
		if input == "EOF" {
			break
		}
		// 构建请求体
		reqData := &requestBody{
			Key:    "9559869b4db24a73b1473c73f33c8494",
			Info:   input,
			UserId: userid,
		}
		// 转换为json
		byteData, _ := json.Marshal(&reqData)
		// 请求聊天机器人接口
		req, err := http.NewRequest("POST",
			"https://www.tuling123.com/openapi/api",
			bytes.NewReader(byteData))
		req.Header.Set("Content-Type", "application/json;charset=utf-8")
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("连接异常")
		} else {
			// 将结果从json中解析出来
			body, _ := ioutil.ReadAll(resp.Body)
			var respData responseBody
			json.Unmarshal(body, &respData)
			fmt.Println("AI：" + respData.Text)
		}
		resp.Body.Close()
	}
}

func main() {
	var input string
	fmt.Println("输入 'EOF' 结束对话: ")
	// 创建通道
	channel := make(chan string)
	// main结束时关闭通道
	defer close(channel)
	// 启动goroutine 运行机器人问答线程
	go process(channel, string(rand.Int63()))
	for {
		// 从命令行中读取输入
		fmt.Scanf("%s", &input)
		// 将输入放入通道中
		channel <- input
		// 结束程序
		if input == "EOF" {
			break
		}
	}
}
