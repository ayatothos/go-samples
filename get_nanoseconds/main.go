package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Date(2020, 2, 3, 14, 26, 31, 123456780, time.Local)

	fmt.Println(t.Format("2006/01/02 15:04:05.000"))
	fmt.Println(t.Format("2006/01/02 15:04:05.000000"))
	fmt.Println(t.Format("2006/01/02 15:04:05.000000000"))

	fmt.Println(t.Format("2006/01/02 15:04:05.999999999")) // 0は切り捨てされる。

	n := time.Now() // time.Now()ではナノ秒が取得できない。

	fmt.Println(n.Format("2006/01/02 15:04:05.000"))
	fmt.Println(n.Format("2006/01/02 15:04:05.000000"))
	fmt.Println(n.Format("2006/01/02 15:04:05.000000000"))
}
