package gbsearch

import (
	"fmt"
	"testing"
)

func TestGeneralSearch(t *testing.T) {
	res, err := Search(ISBN, "9780321774637", nil)
	if err != nil {
		t.Error("Should not have gotten an error with this search")
	}
	if res.TotalItems != 1 {
		t.Error("ISBN search should have returned only one value")
	}
}

func TestAuthorSearch(t *testing.T) {
	res, err := Search(InAuthor, "Mark Summerfield", nil)
	if err != nil {
		t.Error("Should not have gotten an error with this search")
	}
	if res.TotalItems < 1 {
		t.Error("ISBN search should have returned at least one value")
	}

	for _, item := range res.Items {
		if item.VolumeInfo.Title == "Programming in Go" {
			return
		}
	}
	t.Error("Should have found a known authors book")
}

/*
func BenchmarkGeneralSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Search(ISBN, "9780321774637", nil)
	}
}
*/

func ExampleSearch() {

	res, err := Search(ISBN, "9780321774637", nil)
	if err != nil {
		return
	}

	fmt.Println(res.TotalItems)
	fmt.Println(res.Items[0].VolumeInfo.Title)

	// Output:
	// 1
	// Programming in Go

}
