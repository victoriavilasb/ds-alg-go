package structures

import (
	"errors"
)

type Node struct {
	next *Node
	prev *Node
	elem int
}

type LinkedList struct {
	Head     *Node
	Tail     *Node
	Max_size int
	Size     int
}

func NewLinkedList(max_size int) *LinkedList {
	return &LinkedList{
		Head:     nil,
		Tail:     nil,
		Size:     0,
		Max_size: max_size,
	}
}

func (ll *LinkedList) InsertNodeInFrom(n *Node) error {
	if ll.Size == ll.Max_size {
		return errors.New("list is full")
	}

	ll.Size++

	// if elem is the first of the list
	if ll.Head == nil {
		ll.Head = n
		ll.Tail = n

		return nil
	}

	n.next = ll.Head
	ll.Head = n

	return nil
}

func (ll *LinkedList) InsertInFront(elem int) error {
	if ll.Size == ll.Max_size {
		return errors.New("cannot insert another element, list is full")
	}

	ll.Size++

	// if elem is the first of the list
	if ll.Head == nil {
		node := &Node{
			next: nil,
			prev: nil,
			elem: elem,
		}

		ll.Head = node
		ll.Tail = node

		return nil
	}

	// put elem in head
	nn := &Node{
		next: ll.Head,
		prev: nil,
		elem: elem,
	}

	ll.Head = nn

	return nil
}

func (ll *LinkedList) RemoveByNode(n *Node) {
	if ll.Head == n {
		ll.Head = nil
		ll.Tail = nil

		return
	}

	n.prev.next = n.next
	ll.Size -= 1
}

func (ll *LinkedList) FindNode(elem int) (*Node, error) {
	it := ll.Head
	for {
		if it == nil {
			return nil, errors.New("element not found")
		}

		if it.elem != elem {
			it = it.next
		} else {
			return it, nil
		}
	}
}

func (ll *LinkedList) Remove(elem int) error {
	it := ll.Head
	for {
		if it == nil {
			return errors.New("element not found in list")
		}

		if it.elem != elem {
			it = it.next
			continue
		} else if ll.Size == 1 {
			ll.Head = nil
			ll.Tail = nil
		} else if it.elem == ll.Head.elem {
			ll.Head = it.next
		} else {
			it.prev.next = it.next
		}

		ll.Size -= 1
		return nil

	}
}

func (ll *LinkedList) Find(elem int) bool {
	it := ll.Head
	for {
		if it == nil {
			return false
		}

		if it.elem != elem {
			it = it.next
		} else {
			return true
		}
	}
}

func (ll *LinkedList) GetElements() []int {
	it := ll.Head
	elements := []int{}
	for {
		if it == nil {
			return elements
		}

		elements = append(elements, it.elem)
		it = it.next
	}
}
