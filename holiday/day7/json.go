package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	Name  string `json,string`
	Age   int
	Rmb   float64
	Sex   bool
	Hobby []string
}

//结构体转json字符串
func main071() {
	p := person{Name: "荣", Age: 18, Rmb: 1000.05, Sex: false, Hobby: []string{"土炮安排", "打球"}}
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
		return
	} else {
		jsonstr := string(bytes)
		fmt.Printf("type=%T,value=%v\n", jsonstr, jsonstr)
	}
}

//map转json字符串
func main072() {
	maps := make(map[string]interface{})
	maps["name"] = "阿荣"
	maps["age"] = 30
	maps["rmb"] = 500.55
	maps["hobby"] = []string{"喝酒", "烫头"}

	fmt.Printf("type:=%T,value=%v", maps, maps)
	bytes, err := json.Marshal(maps)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
		return
	} else {
		jsonStr := string(bytes)
		fmt.Printf("type=%T,value=%v\n", jsonStr, jsonStr)
	}

}

// slice转json
func main073() {
	slice := make([]map[string]interface{}, 0)
	slice = append(slice, map[string]interface{}{"name": "常建荣", "age": 18, "rmb": 550.68, "hobby": "我我我"})
	slice = append(slice, map[string]interface{}{"name": "铁汁", "age": 18, "rmb": 550.68, "hobby": "我我我"})
	slice = append(slice, map[string]interface{}{"name": "老铁", "age": 18, "rmb": 550.68, "hobby": "我我我"})
	bytes, _ := json.Marshal(slice)
	fmt.Println(string(bytes))

}

/*
json的反序列化：json转go对象
*/
// json转map
func main074() {
	jsonStr := `{,"Hobby":[]string{"我我我",},"Name":"常建荣","Age":18,"Rmb":550.68,"Sex":false}`
	maps := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &maps)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(maps)
	}

}

// json转slice
func main075() {
	jsonStr := `[{"age":18,"hobby":"我我我","name":"常建荣","rmb":550.68},
				{"age":18,"hobby":"我我我","name":"铁汁","rmb":550.68},
				{"age":18,"hobby":"我我我","name":"老铁","rmb":550.68}]`
	slice := make([]map[string]interface{}, 0)
	err := json.Unmarshal([]byte(jsonStr), &slice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(slice)
	}
}

// json转结构体
func main076() {
	jsonStr := `{"Name":"荣","Age":18,"Rmb":1000.05,"Sex":false,"Hobby":["土炮安排","打球"]}`
	h := person{}
	err := json.Unmarshal([]byte(jsonStr), &h)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(h)
	}

}

/*{}
得到基于指定文件的编码器
encoder := json.NewEncoder(file)
编码go数据为json（写出到文件）
encoder.Encode(goData)
创建基于指定文件的解码器
decoder := json.NewDecoder(file)
从文件中读取并解码json到go数据指针中
decoder.Decoder(goDataPtr)
*/
// 编码
func main077() {
	maps := make(map[string]interface{})
	maps["name"] = "啊荣"
	maps["age"] = 250
	maps["sex"] = false
	maps["hobby"] = []string{"啊荣", "ss", "aa"}

	file, _ := os.OpenFile("test.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	encoder := json.NewEncoder(file)
	err := encoder.Encode(maps)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success!")
	}
}

// 解码
func main() {
	file, _ := os.OpenFile("test.json", os.O_RDONLY, 0666)
	defer func() {
		file.Close()
		fmt.Println("文件已关闭")
	}()
	decoder := json.NewDecoder(file)
	maps := make(map[string]interface{})
	err := decoder.Decode(&maps)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success!", maps)
	}
}
