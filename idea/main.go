package main

import (
	"fmt"
)

func main() {
	i := idea{
		id:    1,
		title: "テストタイトル",
		body:  "テストボディ",
	}
	tmp := *i.Get(1)
	fmt.Println(tmp.title)
	fmt.Println(tmp.body)
}
