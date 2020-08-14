package main

import (
	"fmt"
	"time"
	"github.com/volatiletech/null"
)

func main() {
	nullTime := null.Time{}
	fmt.Println(nullTime) // {0001-01-01 00:00:00 +0000 UTC false}
	fmt.Println(nullTime.MarshalText()) // [110 117 108 108] <nil>
	byte, err := nullTime.MarshalText()
	if err != nil {
		fmt.Println("err", err)
	}
	s := string(byte)
	fmt.Printf("%v type:%T\n", s, s) // null type:string,


	time := null.Time{
		Time: time.Now(),
		Valid: true,
	}
	fmt.Println(time) // {2020-08-14 14:09:59.624455 +0900 JST m=+0.000755212 true}
	fmt.Println(time.MarshalText()) // [50 48 50 48 45 48 56 45 49 52 84 49 52 58 48 57 58 53 57 46 54 50 52 52 53 53 43 48 57 58 48 48] <nil>
	byte2, err := time.MarshalText()
	if err != nil {
		fmt.Println("err", err)
	}
	s2 := string(byte2)
	fmt.Printf("%v type:%T\n", s2, s2) // 2020-08-14T14:09:59.624455+09:00 type:string
}