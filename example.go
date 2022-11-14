/*
 * @Name: One Way Linked List Go
 * @Author: Max Base
 * @Date: 2022-11-14
 * @Repository: https://github.com/basemax/OneWayLinkedListGo
 */

package main

import "fmt"

// Main
func main() {
	// Create a list
	list := List{}
	list.Create(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)
	list.Insert(7)

	// Print the list
	list.Print()

	// Search in the list
	fmt.Println(list.Search(3))

	// Delete from the list
	list.Delete()
	list.Delete()
	list.Delete()
}
