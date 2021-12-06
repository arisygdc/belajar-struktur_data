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
		current = current.next
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
		current = current.next
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

func TestInsertDlink(t *testing.T) {
	dLink := NewD()
	for _, v := range first {
		dLink.Append(v)
	}

	node := dLink.Head()
	i := 0
	for {
		t.Logf("Link: %v, Arr: %v\n", node.Val(), first[i])
		if node.Val() != first[i] {
			t.FailNow()
		}
		if node.Next() == nil {
			break
		}
		i++
		node = node.Next()
	}

	i = len(first) - 1
	for {
		t.Logf("Link: %v, Arr: %v\n", node.Val(), first[i])
		if node.Val() != first[i] {
			t.FailNow()
		}
		if node.Prev() == nil {
			break
		}
		i--
		node = node.Prev()
	}

	node = dLink.Tail()
	i = len(first) - 1
	for {
		t.Logf("Link: %v, Arr: %v\n", node.Val(), first[i])
		if node.Val() != first[i] {
			t.FailNow()
		}
		if node.Prev() == nil {
			break
		}
		i--
		node = node.Prev()
	}

	i = 0
	for {
		t.Logf("Link: %v, Arr: %v\n", node.Val(), first[i])
		if node.Val() != first[i] {
			t.FailNow()
		}
		if node.Next() == nil {
			break
		}
		i++
		node = node.Next()
	}
}

func TestDeleteDlinkError(t *testing.T) {
	dLink := NewD()
	for _, v := range first {
		dLink.Append(v)
	}
	testTable := []struct {
		num    int
		expect error
	}{
		{num: 14, expect: fmt.Errorf("tidak ditemukan")},
		{num: 10, expect: nil},
		{num: 30, expect: nil},
		{num: 60, expect: nil},
		{num: 44, expect: fmt.Errorf("tidak ditemukan")},
	}

	link := NewD()

	for _, n := range testTable {
		t.Logf("Delete number %v", n.num)
		cur := link.Delete(n.num)
		if n.expect != nil {
			if cur != nil {
				t.Log("fail")
				continue
			} else {
				t.Log("Not expect")
				t.FailNow()
			}
		}
		if cur != n.expect {
			t.Logf("%v", cur)
		}
	}
}

func TestDeleteDlinkValue(t *testing.T) {
	second := []int{20, 40, 50, 70}
	dLink := NewD()
	for _, v := range first {
		dLink.Append(v)
	}
	dLink.Delete(10)
	dLink.Delete(30)
	dLink.Delete(60)
	dLink.Delete(80)

	node := dLink.Head()

	for _, n := range second {
		t.Logf("current: %v, expect: %v", node.Val(), n)
		if node.Val() != n {
			t.FailNow()
		}
		node = node.Next()
	}

	node = dLink.Tail()
	num := len(second) - 1
	for {
		t.Logf("current: %v, expect: %v", node.Val(), second[num])
		if node.Val() != second[num] {
			t.FailNow()
		}
		if node.Prev() == nil {
			break
		}
		node = node.Prev()
		num--
	}
}
