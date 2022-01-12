package integer

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(2, 2, 3)
	want := 7

	if got != want {
		t.Errorf("want %d got %d", want, got)
	}
}

func ExampleAdd() {
	sum := Add(36, 33)
	fmt.Println(sum)
	// Output: 69
}
