package main

type Node struct {
	Value interface{}
	next  *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedList() *LinkedList {
	ll := &LinkedList{
		length: 0,
	}
	return ll
}

// Head returns the head of the linked list.
func (ll *LinkedList) Head() *Node {
	return ll.head
}

// Tail returns the tail of the linked list.
func (ll *LinkedList) Tail() *Node {
	return ll.tail
}

type SinglyLinkedList interface {
	Append(n Node)
	Prepend(n Node)
	Delete(n Node)
}

// Next returns the following node in the list.
func (n *Node) Next() *Node {
	return n.next
}

// Append adds a new node to the end of the linked list.
func (ll *LinkedList) Append(n *Node) {
	if ll.head == nil {
		// The list is empty, append the node as head
		ll.head = n
		return
	}

	currentNode := ll.head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = n
	ll.length++
	return
}

// Prepend adds a new node to the front of the linked list.
func (ll *LinkedList) Prepend(n *Node) {
	if ll.head == nil {
		ll.head = n
		return
	}

	n.next = ll.head
	ll.head = n
	ll.length++
	return
}

// Remove deletes the node from the linked list.
func (ll *LinkedList) Remove(n *Node) {
	// If the head is the node to be removed ...
	if ll.head == n {
		// Make the next node the head
		ll.head = n.next
		ll.length--
		return
	}

	currentNode := ll.head
	for currentNode.next != nil {
		// If the next node is the node to be removed ...
		if currentNode.next == n {
			// Set the next node to be the node after the next node
			currentNode.next = currentNode.next.next
			ll.length--
		}
		currentNode = currentNode.next
	}
	return
}
