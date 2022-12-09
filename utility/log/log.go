package log

import "fmt"

func LogReq(req interface{}) {
	fmt.Println(req)
}

func LogResp(res interface{}) {
	fmt.Println(res)
}

func LogError(err error) {
	fmt.Println(err.Error())
}
