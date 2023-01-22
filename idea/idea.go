package main

type idea struct {
	id    int64
	title string
	body  string
}

func (self *idea) Get(id int64) *idea {
	return self
}
