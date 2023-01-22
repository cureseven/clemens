package main

import (
	"testing"
)

func TestGet(t *testing.T) {
	testIdea := Idea{
		id:    1,
		title: "テストタイトル",
		body:  "テストボディ",
	}
	if *testIdea.Get(1) != testIdea {
		t.Fatal("違うよ")
	}
}
