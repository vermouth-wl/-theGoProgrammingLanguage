package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 调用http接口并返回数据
func main() {

	// 构建请求post数据
	// json序列化
	requestBody := `{
        "regNo": "12588",
        "patientId": "P4406392",
        "name": "申京仕",
        "sex": "1",
        "dateOfBirth": "2004-02-08",
        "insuranceNo": "",
        "visitType": "1",
        "fee": "17.00",
        "deptCode": "1010101",
        "visitDept": "神经内一门诊",
        "doctorId": "01582",
        "visitDoctor": "刘树胜",
        "regDate": "20210706144000",
        "feeStatus": "1",
        "tradeNo": "16670162143N",
        "regType": ""
    }`

	var jsonStr = []byte(requestBody)

	url := "http://132.147.1.235:9981/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	// 构建调用客户端
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
