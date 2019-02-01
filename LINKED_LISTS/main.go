package main

import (
	//Remember that importing custom modules is relative to $GOPATH/src
	"CTCI/LINKED_LISTS/linked"
	"fmt"
)

//Game plan:
//Use a map to track the dups
//Keep the pointers to those duplicate nodes in a slice
//Use list.RemoveNode(node) to get rid of the dups
//Now the question is, besides refactoring everything, how can I painlessly reveal the
//node structure
func track_dups(list *linked.LinkedList) {
	cur_node := list.head
	for cur_node != nil {
		cur_node = cur_node.next
	}
	fmt.Println("Hello")

}

func main() {
	//Remember that structs are "types", like int, or string
	var my_list linked.LinkedList
	var nums = []int{1, 2, 3, 4, 1, 2, 5, 6}
	for _, n := range nums {
		my_list.Add(n)
	}
	fmt.Println(my_list.GetList())
	track_dups()
}
