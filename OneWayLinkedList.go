/*
 * @Name: One Way Linked List Go
 * @Author: Max Base
 * @Date: 2022-11-14
 * @Repository: https://github.com/basemax/OneWayLinkedListGo
 */

package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type List struct {
	Head *Node
}

// Create, O(1)
func (list *List) Create(value interface{}) {
	list.Head = &Node{Value: value}
}

// InsertBegin, O(1)
func (list *List) InsertBegin(value interface{}) {
	node := &Node{Value: value}
	node.Next = list.Head
	list.Head = node
}

// InsertEnd, O(n)
func (list *List) InsertEnd(value interface{}) {
	node := list.Head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = &Node{Value: value}
}

// InsertAfter, O(n)
func (list *List) InsertAfter(value interface{}, after interface{}) {
	node := list.Head
	for node != nil {
		if node.Value == after {
			newNode := &Node{Value: value}
			newNode.Next = node.Next
			node.Next = newNode
			return
		}
		node = node.Next
	}
}

// InsertBefore, O(n)
func (list *List) InsertBefore(value interface{}, before interface{}) {
	node := list.Head
	for node != nil {
		if node.Next.Value == before {
			newNode := &Node{Value: value}
			newNode.Next = node.Next
			node.Next = newNode
			return
		}
		node = node.Next
	}
}

// Delete, O(1)
// func (list *List) Delete() {
// 	if list.Head != nil {
// 		list.Head = list.Head.Next
// 	}
// }

// DeleteBegin, O(1)
func (list *List) DeleteBegin() {
	if list.Head != nil {
		list.Head = list.Head.Next
	}
}

// DeleteEnd, O(n)
func (list *List) DeleteEnd() {
	node := list.Head
	for node.Next.Next != nil {
		node = node.Next
	}
	node.Next = nil
}

// DeleteValue, O(n)
func (list *List) DeleteValue(value interface{}) {
	node := list.Head
	for node != nil {
		if node.Next.Value == value {
			if node.Next.Next != nil {
				node.Next = node.Next.Next
			} else {
				node.Next = nil
			}
			return
		}
		node = node.Next
	}
}

// DeleteAfterValue, O(n)
func (list *List) DeleteAfterValue(value interface{}) {
	node := list.Head
	for node != nil {
		if node.Value == value {
			node.Next = node.Next.Next
			return
		}
		node = node.Next
	}
}

// DeleteBeforeValue, O(n)
func (list *List) DeleteBeforeValue(value interface{}) {
	node := list.Head
	for node != nil {
		if node.Next.Next.Value == value {
			node.Next = node.Next.Next
			return
		}
		node = node.Next
	}
}

// DeleteIndex, O(n)
func (list *List) DeleteIndex(index int) {
	if index < 0 {
		return
	}
	node := list.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
}

// DeleteAfterIndex, O(n)
func (list *List) DeleteAfterIndex(index int) {
	// list.DeleteIndex(index + 1)

	node := list.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
}

// DeleteBeforeIndex, O(n)
func (list *List) DeleteBeforeIndex(index int) {
	// list.DeleteIndex(index - 1)

	node := list.Head
	for i := 0; i < index-1; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
}

// Search, O(n)
func (list *List) Search(value interface{}) bool {
	node := list.Head
	for node != nil {
		if node.Value == value {
			return true
		}
		node = node.Next
	}
	return false
}

// Print, O(n)
func (list *List) Print() {
	node := list.Head
	for node != nil {
		fmt.Println(node.Value)
		node = node.Next
	}
}

// ToString, O(n)
func (list *List) ToString() string {
	node := list.Head
	result := ""
	for node != nil {
		result += fmt.Sprintf("%v", node.Value)
		node = node.Next
	}
	return result
}
