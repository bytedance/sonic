package issue_test

import (
	"testing"

	"github.com/bytedance/sonic"
)

type issue1000Level struct {
	Price string `json:"price"`
	Qty   string `json:"qty"`
	ID    int64  `json:"id"`
}

type issue1000Book struct {
	Levels []issue1000Level `json:"levels"`
}

func TestIssue1000EmptyArrayPreservesSliceCapacity(t *testing.T) {
	const empty = `{"levels":[]}`
	const full = `{"levels":[{"price":"1","qty":"1","id":1}]}`

	book := issue1000Book{Levels: make([]issue1000Level, 0, 64)}
	backing := &book.Levels[:cap(book.Levels)][0]

	if err := sonic.UnmarshalString(full, &book); err != nil {
		t.Fatal(err)
	}
	if cap(book.Levels) != 64 || &book.Levels[:cap(book.Levels)][0] != backing {
		t.Fatalf("after non-empty array: cap=%d, backing preserved=%v; want 64, true", cap(book.Levels), &book.Levels[:cap(book.Levels)][0] == backing)
	}
	if err := sonic.UnmarshalString(empty, &book); err != nil {
		t.Fatal(err)
	}
	if len(book.Levels) != 0 || cap(book.Levels) != 64 {
		t.Fatalf("after empty array: len=%d cap=%d; want 0, 64", len(book.Levels), cap(book.Levels))
	}
	if &book.Levels[:cap(book.Levels)][0] != backing {
		t.Fatal("after empty array: backing array changed")
	}
	if err := sonic.UnmarshalString(full, &book); err != nil {
		t.Fatal(err)
	}
	if len(book.Levels) != 1 || cap(book.Levels) != 64 || &book.Levels[:cap(book.Levels)][0] != backing || book.Levels[0].ID != 1 {
		t.Fatalf("after reused decode: len=%d cap=%d backing preserved=%v value=%+v; want 1, 64, true, ID=1", len(book.Levels), cap(book.Levels), &book.Levels[:cap(book.Levels)][0] == backing, book.Levels)
	}

	var emptyBook issue1000Book
	if err := sonic.UnmarshalString(empty, &emptyBook); err != nil {
		t.Fatal(err)
	}
	if emptyBook.Levels == nil || len(emptyBook.Levels) != 0 {
		t.Fatalf("empty array decoded to %#v; want a non-nil empty slice", emptyBook.Levels)
	}
}
