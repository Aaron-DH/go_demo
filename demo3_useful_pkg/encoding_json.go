package main

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct {
	UserID    int
	UserName  string
	IsStudent bool
}

// 定义struct时加上tag会自动转成tag的名称
type UserInfoWithTag struct {
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	IsStudent bool   `json:"is_student"`
}

type Role struct {
	RoleID   int    `json: "role_id"`
	RoleName string `json: "role_name"`
}

// 嵌套struct
type UserRoles struct {
	UserInfoWithTag
	roles []Role
}

func testStruct2Json() {
	userinfo := UserInfo{
		UserID:    111,
		UserName:  "Aaron",
		IsStudent: true,
	}
	data1, _ := json.Marshal(userinfo)
	fmt.Println(string(data1))

	// 使用带tag的struct最终会转成user_id,user_name这些tag名称
	userinfotag := UserInfoWithTag{
		UserID:    111,
		UserName:  "Aaron",
		IsStudent: true,
	}
	data2, _ := json.Marshal(userinfotag)
	fmt.Println(string(data2))

	// 定义嵌套struct
	userroles := UserRoles{
		UserInfoWithTag: UserInfoWithTag{
			UserID:    111,
			UserName:  "Aaron",
			IsStudent: true,
		},
		roles: []Role{
			Role{RoleID: 1, RoleName: "admin"},
			Role{RoleID: 2, RoleName: "user"},
		},
	}

	// 对嵌套struct进行json序列化转换失败
	data3, _ := json.Marshal(userroles)
	fmt.Println(string(data3))
}

func testJson2Struct() {
	var data = `
	    {
	        "user_id":1,
	        "user_name":"aaron",
	        "is_student": false,
	        "roles": [
	            {
	                "role_id": 1,
		            "role_name": "admin"
	            },
		        {
		            "role_id": 2,
		            "role_name": "user"
				}
			]
		}
	`

	// 嵌套的roles没有解析出来
	p1 := &UserInfoWithTag{}
	json.Unmarshal([]byte(data), p1)
	fmt.Println(*p1)

	// 没有定义tag, 无法转换
	p2 := &UserInfo{}
	json.Unmarshal([]byte(data), p2)
	fmt.Println(*p2)
}

func main() {
	testStruct2Json()

	testJson2Struct()
}
