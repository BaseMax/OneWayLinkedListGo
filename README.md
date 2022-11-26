# One-Way Linked-List Go

This is a simple implementation of a one-way linked-list in Go.

# Features

- [x] Create
- [x] Insert at head
- [x] Insert at tail
- [x] Insert at index
- [x] Delete at head
- [ ] Delete at tail
- [ ] Delete at index
- [x] Get length
- [ ] Get head
- [ ] Get tail
- [ ] Get index
- [ ] Get index of value
- [ ] Get value at index
- [ ] Get values
- [ ] Get values in reverse

## Structure

```go
// Create, O(1)
func (list *List) Create(value interface{})

// InsertBegin, O(1)
func (list *List) InsertBegin(value interface{})

// InsertEnd, O(n)
func (list *List) InsertEnd(value interface{})

// InsertAfter, O(n)
func (list *List) InsertAfter(value interface{}, after interface{})

// InsertBefore, O(n)
func (list *List) InsertBefore(value interface{}, before interface{})

// Delete, O(1)
func (list *List) Delete()

// Search, O(n)
func (list *List) Search(value interface{}) bool

// Print, O(n)
func (list *List) Print()

// ToString, O(n)
func (list *List) ToString()
```

© Copyright 2022, Max Base
