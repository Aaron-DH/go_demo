package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func testSimpleHttp() {

	fmt.Println("Begin http get")
	/**
	curl -X GET http://127.0.0.1:8080/check/health
	{
	     "message": "Healthy",
	     "status": 200
	}
	*/
	resp1, err1 := http.Get("http://127.0.0.1:8080/check/health")
	if resp1 != nil {
		defer resp1.Body.Close()
	}
	if err1 != nil {
		fmt.Println("-> result error: ", err1)
	} else {
		fmt.Println("->	result response code: ", resp1.StatusCode)
		body2, _ := ioutil.ReadAll(resp1.Body)
		fmt.Println("->	result body:", string(body2))
	}

	fmt.Println("Begin http post")
	/**
	curl -i -g -X POST http://127.0.0.1:8080/login -d '{"username":"dinghh","password":"dinghh"}'
	HTTP/1.1 200 OK
	Cache-Control: no-cache, no-store, max-age=0, must-revalidate, value
	Content-Type: application/json; charset=utf-8
	Expires: Thu, 01 Jan 1970 00:00:00 GMT
	Last-Modified: Thu, 18 Jul 2019 03:52:52 GMT
	X-Auth-Token: lahBKGNZR
	Date: Thu, 18 Jul 2019 03:52:52 GMT
	Content-Length: 257
	{
	    "code": 0,
	    "data": {
	        "create_id": 1,
	        "create_time": "2019-07-09T16:49:04+08:00",
	        "is_del": 0,
	        "update_id": 0,
	        "update_time": "0001-01-01T00:00:00Z",
	        "user_id": 1,
	        "user_mail": "dinghh@awcloud.com",
	        "user_name": "dinghh",
	        "user_passwd": "dinghh",
	        "user_status": 0
	    },
	    "message": "OK"
	}
	*/
}

func main() {
	testSimpleHttp()
}
