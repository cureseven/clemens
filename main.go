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
	idea.Get(1)
}

func (self *Idea) Get(id int64) () {
	fmt.Println("title: ", self.title)
}
