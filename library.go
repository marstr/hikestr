package hikestr

import "github.com/marstr/collection"

type Library struct {
	Head LibraryNode `json:"head"`
}

type LibraryNode struct {
	Hikes  collection.Enumerable
	Parent *LibraryNode
}
