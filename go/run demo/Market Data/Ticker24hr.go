package main

import (
	"demo/spot"
	"fmt"
)

//说明：
//在params中输入json格式的参数,如:`{"symbol":"BTCUSDT",	"limit":"200"}`
//如果没有参数则输入:""

//现货参数
// var params string = ""
var params string = `{"symbol":"BTCUSDT"}`

func main() {

	//接口调用
	Ticker24hr := spotList.Ticker24hr(params)
	fmt.Println("返回信息:", Ticker24hr)

}
