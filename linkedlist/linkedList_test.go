package linkedlist

import (
	"fmt"
	"testing"
)

var first = []int{10, 20, 30, 40, 50, 60, 70, 80}

func TestInsertLink(t *testing.T) {
	linkA := New()
	for _, v := range first {
		linkA.Append(v)
	}

	current := linkA.GetNode()
	for _, v := range first {
		t.Logf("Link: %v, array: %v\n", current.GetVal(), v)
		if current.GetVal() != v {
			t.FailNow()
		}
		current = current.Next
	}

	for _, v := range first {
		linkA.Prepend(v)
	}

	current = linkA.GetNode()
	for i := len(first) - 1; i >= 0; i-- {
		t.Logf("Link: %v, array: %v\n", current.GetVal(), first[i])
		if current.GetVal() != first[i] {
			t.FailNow()
		}
		current = current.Next
	}
}

func TestDeleteLink(t *testing.T) {
	testTable := []struct {
		num    int
		expect error
	}{
		{num: 14, expect: fmt.Errorf("tidak ditemukan")},
		{num: 10, expect: nil},
		{num: 30, expect: nil},
		{num: 80, expect: nil},
		{num: 44, expect: fmt.Errorf("tidak ditemukan")},
	}
	link := New()
	for _, v := range first {
		link.Prepend(v)
	}

	for _, n := range testTable {
		t.Logf("Delete number %v", n.num)
		cur := link.Delete(n.num)
		if n.expect != nil {
			if cur != nil {
				continue
			} else {
				t.FailNow()
			}
		}
		if cur != n.expect {
			t.FailNow()
		}
	}

}

func TestInsertAndDeleteDlink(t *testing.T) {
	dLink := NewD()
	if dLink.IsEmpty() != true {
		t.FailNow()
	}
	for _, v := range first {
		dLink.AddNode(v)
	}
	dLink.Delete(first[0])
	dLink.Delete(first[3])
	dLink.Delete(first[7])

	arr := dLink.Array()
	RArr := dLink.ReverseArray()
	if len(arr) != len(RArr) {
		t.Fail()
		t.Errorf("length array (%d) and reverse (%d) not match ", len(arr), len(RArr))
	}

	last := len(arr) - 1

	for _, v := range RArr {
		if arr[last] != v {
			t.Logf("got result, %d and reverse %d. expected match", arr[last], v)
		}
	}
}
