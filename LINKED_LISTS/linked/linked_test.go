package linked

import (
	//"fmt"
	"testing"
)

//TODO
//Find out how to load in a LinkedList data structure to each test
//A test fixture, since creating each structure over and over again has code smell

func TestAddFirst(t *testing.T) {
	var expected int = 4
	list := LinkedList{head: nil, tail: nil}
	list.Add(4)
	if list.head.value != expected {
		t.Errorf("Expected: %d, Actual: %d", expected, list.head.value)
	}
}

func TestAddSecond(t *testing.T) {
	var expected_nums = []int{1, 2}
	list := LinkedList{head: nil, tail: nil}
	for _, n := range expected_nums {
		list.Add(n)
	}
	if list.head.value != 1 {
		t.Errorf("HEAD. Expected: %d, Actual: %d", 1, list.head.value)
	} else if list.tail.value != 2 {
		t.Errorf("TAIL. Expected: %d, Actual: %d", 2, list.tail.value)
	}
}

func TestGetList(t *testing.T) {
	list := LinkedList{head: nil, tail: nil}
	var expected_nums = []int{1, 2, 3, 4}
	var expected_str = "1 <-> 2 <-> 3 <-> 4"
	for _, n := range expected_nums {
		list.Add(n)
	}
	var my_str = list.GetList()

	if my_str != expected_str {
		t.Errorf("Expected: %s, Actual: %s", expected_str, my_str)
	}
}

func TestRemoveHead(t *testing.T) {
	list := LinkedList{head: nil, tail: nil}
	var expected_nums = []int{1, 2, 3, 4}
	var expected_no_head = "2 <-> 3 <-> 4"

	for _, n := range expected_nums {
		list.Add(n)
	}
	list.RemoveNode(list.head)
	var my_str = list.GetList()

	if my_str != expected_no_head {
		t.Errorf("Expected: %s, Actual: %s", expected_no_head, my_str)
	}
}

func TestRemoveTail(t *testing.T) {
	list := LinkedList{head: nil, tail: nil}
	var expected_nums = []int{1, 2, 3, 4}
	var expected_no_tail = "1 <-> 2 <-> 3"

	list.Make(expected_nums)
	list.RemoveNode(list.tail)

	var my_str = list.GetList()

	if my_str != expected_no_tail {
		t.Errorf("Expected: %s, Actual: %s", expected_no_tail, my_str)
	}
}

func TestRemoveMiddle(t *testing.T) {
	list := LinkedList{head: nil, tail: nil}
	var expected_nums = []int{1, 2, 3, 4}
	var expected_no_mid = "1 <-> 3 <-> 4"

	list.Make(expected_nums)
	list.RemoveNode(list.head.next)
	var my_str = list.GetList()

	if my_str != expected_no_mid {
		t.Errorf("Expected: %s, Actual: %s", expected_no_mid, my_str)
	}
}
