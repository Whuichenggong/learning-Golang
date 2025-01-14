package main

type MyReader struct {
	data string
}

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less returns whether the element with index i should sort
	// before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int) string
}

// io.ReadWriter
type ReadWriter interface {
	Reader
	Writer
}
