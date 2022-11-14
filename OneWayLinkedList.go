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

// Insert, O(1)
func (list *List) Insert(value interface{}) {
	node := &Node{Value: value}
	node.Next = list.Head
	list.Head = node
}

// Delete, O(1)
func (list *List) Delete() {
	list.Head = list.Head.Next
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
