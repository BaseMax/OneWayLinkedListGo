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

	// Insert some values at begin
	list.InsertBegin(2)
	list.InsertBegin(3)
	list.InsertBegin(4)

	// Insert some values at end
	list.InsertEnd(5)
	list.InsertEnd(6)
	list.InsertEnd(7)

	// Insert some values after
	list.InsertAfter(8, 5)
	list.InsertAfter(9, 6)
	list.InsertAfter(10, 7)

	// Search in the list
	fmt.Println(list.Search(3))

	// Print the list
	list.Print()

	// Re-create a list
	list.Create(100)

	// Print the list
	list.Print()

	// Search in the list
	fmt.Println(list.Search(3))

	// Delete from the list
	list.Delete()
	list.Delete()
	list.Delete()

	// Print the list
	list.Print()
}
