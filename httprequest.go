package main

import (
    "flag"
    "fmt"
    "net/http"
    "time"
)

const readMe = `输入第一个是网址 第二个是目标访问数量 举例:***.exe -u www.baidu.com -n 100`

func CreateHttpReq(url string, num int) {
    for i := 0; i < num; i++ {
        resp, _ := http.Get(url)
        resp.Body.Close()
        time.Sleep(time.Microsecond * 100)
        fmt.Printf("第%d次访问", i+1)
    }
}

var targetUrl string
var targetNum int

func Init() {
    flag.StringVar(&targetUrl, "u", "", "please paste the url")
    flag.IntVar(&targetNum, "n", 0, "please input the number you want to connection")
}

func main() {
    fmt.Println(readMe)
    Init()
    flag.Parse()
    fmt.Println(targetUrl, targetNum)
    CreateHttpReq(targetUrl, targetNum)
    fmt.Println("DONE")
}
