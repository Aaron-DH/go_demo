package main

import (
	"fmt"
	"strconv"
)

func parseStr() {
	// Str转Int(正常)
	data1, err1 := strconv.Atoi("12345")
	if err1 != nil {
		fmt.Println("data1: error ", err1)
	}
	data1 += 3
	fmt.Println("data1: ", data1)

	// Str转Int(无效、出错)
	data2, err2 := strconv.Atoi("abc123")
	if err2 != nil {
		fmt.Println("data2: error ", err2)
	}
	fmt.Println("data2: ", data2)

	// Int转Str(正常)
	data3 := strconv.Itoa(12345)
	data3 += "3"
	fmt.Println("data3: ", data3)
}

func parseBool() {
	// Str转Bool(正常)
	bool1, err1 := strconv.ParseBool("true")
	if err1 != nil {
		fmt.Println("bool1: error ", err1)
	}
	fmt.Println("bool1: ", bool1)

	// Str转Bool(正常)
	bool2, err2 := strconv.ParseBool("True")
	if err2 != nil {
		fmt.Println("bool2: error ", err2)
	}
	fmt.Println("bool2: ", bool2)

	// Str转Bool(正常)
	bool3, err3 := strconv.ParseBool("TRUE")
	if err3 != nil {
		fmt.Println("bool3: error ", err3)
	}
	fmt.Println("bool3: ", bool3)

	// Str转Bool(错误)
	bool4, err4 := strconv.ParseBool("tRue")
	if err4 != nil {
		fmt.Println("bool4: error ", err4)
	}
	fmt.Println("bool4: ", bool4)
}

func parseFloar() {
	// Str转Float
	float1, err1 := strconv.ParseFloat("3.1415", 64)
	if err1 != nil {
		fmt.Println("float1: error ", err1)
	}
	fmt.Println("float1: ", float1)

	float2, err2 := strconv.ParseFloat("3", 64)
	if err2 != nil {
		fmt.Println("float1: error ", err2)
	}
	fmt.Println("float2: ", float2)

	float3, err3 := strconv.ParseFloat("3.1415", 32)
	if err3 != nil {
		fmt.Println("float3: error ", err3)
	}
	fmt.Println("float3: ", float3)
}

func parseInt() {
	// ParseInt()和ParseUint()有3个参数：
	// func ParseInt(s string, base int, bitSize int) (i int64, err error)
	// func ParseUint(s string, base int, bitSize int) (uint64, error)
	// bitSize参数表示转换为什么位的int/uint，有效值为0、8、16、32、64。当bitSize=0的时候，表示转换为int或uint类型。例如bitSize=8表示转换后的值的类型为int8或uint8。
	// base参数表示以什么进制的方式去解析给定的字符串，有效值为0、2-36。当base=0的时候，表示根据string的前缀来判断以什么进制去解析：0x开头的以16进制的方式去解析，0开头的以8进制方式去解析，其它的以10进制方式解析。

	// 以10进制方式解析"-42"，保存为int64类型：
	i1, _ := strconv.ParseInt("-42", 10, 64)
	fmt.Println("int1: ", i1)

	// 以5进制方式解析"23"，保存为int64类型：
	// 因为5进制的时候，23表示进位了2次，再加3，所以对应的十进制数为5*2+3=13
	i2, _ := strconv.ParseInt("23", 5, 64)
	fmt.Println("int2: ", i2)

	i3, _ := strconv.ParseInt("23", 16, 64)
	fmt.Println("int3: ", i3)

	i4, _ := strconv.ParseInt("23", 15, 64)
	fmt.Println("int4: ", i4)
}
func main() {
	parseStr()

	parseBool()

	parseFloar()

	parseInt()
}
