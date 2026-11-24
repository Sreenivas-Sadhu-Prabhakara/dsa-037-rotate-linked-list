package main

import "fmt"

// Node is a singly linked list node.
type Node struct {
	Val  int
	Next *Node
}

// RotateRight rotates the list right by k places: the last k nodes move to
// the front. Close the list into a ring, then break it at the right spot.
func RotateRight(head *Node, k int) *Node {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	n := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	k %= n
	if k == 0 {
		return head
	}
	tail.Next = head // ring
	newTail := head
	for i := 1; i < n-k; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil
	return newHead
}

func fromSlice(vals []int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &Node{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func toSlice(head *Node) []int {
	out := []int{}
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

func main() {
	fmt.Println(toSlice(RotateRight(fromSlice([]int{1, 2, 3, 4, 5}), 2)))
}
