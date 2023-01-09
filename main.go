package main

import (
	"fmt"
)

type Idea struct {
	id	int64
	title	string
	body	string
}

func main() {
	idea := Idea{
		id:	1,
		title:	"初めてのアイディア",
		body:	"初々しいよ.",
	}
	tmp := idea.Get(1)
	fmt.Println(tmp.id)
}

func (self *Idea) Get(id int64) (*Idea) {
	return self
}

