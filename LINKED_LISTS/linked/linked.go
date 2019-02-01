package linked

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
	prev  *node
}

type LinkedList struct {
	head *node
	tail *node
}

func (list *LinkedList) Make(values []int) {
	for _, n := range values {
		list.Add(n)
	}
}

func (list *LinkedList) Add(value int) {
	new_node := node{value: value}
	if list.head == nil {
		list.head = &new_node
		list.tail = &new_node
	} else {
		cur_node := list.tail
		cur_node.next = &new_node
		new_node.prev = cur_node
		list.tail = &new_node
	}

}

func (list *LinkedList) GetList() string {
	var sb strings.Builder
	cur_node := list.head
	for cur_node.next != nil {
		fmt.Fprintf(&sb, "%d <-> ", cur_node.value)
		cur_node = cur_node.next
	}
	fmt.Fprintf(&sb, "%d", cur_node.value)
	return sb.String()
}

func (list *LinkedList) RemoveNode(cur_node *node) {
	if cur_node == list.head {
		list.head = cur_node.next
		cur_node.next.prev = nil
		cur_node.next = nil
	} else if cur_node == list.tail {
		list.tail = cur_node.prev
		cur_node.prev.next = nil
		cur_node.prev = nil
	} else {
		cur_node.prev.next = cur_node.next
		cur_node.next.prev = cur_node.prev
		cur_node.next = nil
		cur_node.prev = nil
	}
}
