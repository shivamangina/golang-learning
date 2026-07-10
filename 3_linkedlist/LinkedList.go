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
}

func (ll *LinkedList) Append(n *ListNode) {
	if ll.Head == nil {
		ll.Head = n
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = n

}

func (ll *LinkedList) Preappend(n *ListNode) {
	second := ll.Head
	n.Next = second
	ll.Head = n

}

func (ll *LinkedList) Reverse() {
	var prev *ListNode
	curr := ll.Head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	ll.Head = prev
}

func (ll *LinkedList) Reorder() {
	var prev *ListNode
	curr := ll.Head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	ll.Head = prev
}

func (ll *LinkedList) Print() {
	cur := ll.Head
	for cur != nil {
		println(cur.Val)
		cur = cur.Next
	}

}

func main() {
	// 1,2,3,4,5
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)
	node5 := NewNode(5)

	ll := LinkedList{
		Head: nil,
	}

	ll.Append(node1)
	ll.Append(node2)
	ll.Append(node3)
	ll.Append(node4)
	ll.Append(node5)

	ll.Print()

	println("------")
	// ll.Reverse()
	ll.Reorder()

	println("------")

	ll.Print()

}
