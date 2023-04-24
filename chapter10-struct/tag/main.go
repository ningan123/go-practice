package main

import (
	"encoding/json"
	"fmt"
)

// type Monster struct {
// 	Name  string
// 	Age   int
// 	Skill string
// }

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	//

	monster := Monster{"test", 10, "打拳"}

	// 将monster变量序列化为json格式字串
	jsonMonster, err := json.Marshal(monster) // 该函数使用了反射
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(jsonMonster) // func json.Marshal(v interface{}) ([]byte, error)
	fmt.Println(string(jsonMonster))

}

/*
// 结构体定义 不加tag 输出是大写的
[123 34 78 97 109 101 34 58 34 116 101 115 116 34 44 34 65 103 101 34 58 49 48 44 34 83 107 105 108 108 34 58 34 230 137 147 230 139 179 34 125]
{"Name":"test","Age":10,"Skill":"打拳"}

// 结构体定义 加上tag 输出是小写的
[123 34 110 97 109 101 34 58 34 116 101 115 116 34 44 34 97 103 101 34 58 49 48 44 34 115 107 105 108 108 34 58 34 230 137 147 230 139 179 34 125]
{"name":"test","age":10,"skill":"打拳"}
*/
