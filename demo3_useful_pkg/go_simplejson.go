package main

import (
	"fmt"
	json "github.com/bitly/go-simplejson"
)

type UserInfo struct {
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
	UserInfo
	roles []Role
}

func testString2Json() {
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
	jsondata, _ := json.NewJson([]byte(data))
	fmt.Println(jsondata)
	roles, _ := jsondata.Get("roles").Array()
	fmt.Println(roles)
	fmt.Println(len(roles))
}

func main() {
	testString2Json()
}
