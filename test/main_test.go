package clemens

import (
	"testing"

	"github.com/cureseven/clemens"
)

func TestGet(t *testing.T) {
	testIdea := clemens.Idea{
		id: 1,
		title: "テストタイトル",
		body: "テストボディ",
	}
	if testIdea.Get != testIdea {
		t.Fatal("違うよ")
	}
}
