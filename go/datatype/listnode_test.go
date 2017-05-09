package datatype

import (
	"testing"
)

func TestListToListNode(t *testing.T) {
	a1 := []int{1, 2, 3, 4, 5}
	l1 := ListToListNode(a1)
	for i := 0; i < len(a1); i++ {
		if a1[i] != l1.Val {
			t.FailNow()
		}
		l1 = l1.Next
	}
	if l1 != nil {
		t.FailNow()
	}
}
