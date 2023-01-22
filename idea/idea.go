package main

type Idea struct {
	id    int64
	title string
	body  string
}

func (self *Idea) Get(id int64) *Idea {
	return self
}
