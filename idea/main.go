package main

import (
	"fmt"
)

func main() {
	idea := Idea{
		id:    1,
		title: "初めてのアイディア",
		body:  "初々しいよ.",
	}
	tmp := idea.Get(1)
	fmt.Println(tmp.id)
}
