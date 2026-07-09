package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(data int) *ListNode {
	return &ListNode{Val: data, Next: nil}
}

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
}

func main() {

}
