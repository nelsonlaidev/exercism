package pov

import "slices"

type Tree struct {
	// Add the needed fields here
	value    string
	children []*Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	return &Tree{
		value:    value,
		children: children,
	}
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.value
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	return tr.children
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// POV problem-specific functions

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	path := findPath(tr, from)

	if path == nil {
		return nil
	}

	for i := 0; i < len(path)-1; i++ {
		p := path[i]
		c := path[i+1]

		for j, child := range p.children {
			if child == c {
				p.children = append(p.children[:j], p.children[j+1:]...)
				break
			}
		}
		c.children = append(c.children, p)
	}

	return path[len(path)-1]
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	fromPath := pathFromRoot(tr, from)
	toPath := pathFromRoot(tr, to)

	if fromPath == nil || toPath == nil {
		return nil
	}

	i := 0

	for i < len(fromPath) && i < len(toPath) && fromPath[i] == toPath[i] {
		i++
	}

	up := fromPath[i:]
	down := toPath[i-1:]
	result := make([]string, 0, len(up)+len(down))

	for _, u := range slices.Backward(up) {
		result = append(result, u)
	}

	result = append(result, down...)

	return result
}

func findPath(tr *Tree, from string) []*Tree {
	if tr.Value() == from {
		return []*Tree{tr}
	}

	for _, child := range tr.Children() {
		if path := findPath(child, from); path != nil {
			return append([]*Tree{tr}, path...)
		}
	}

	return nil
}

func pathFromRoot(tr *Tree, target string) []string {
	if tr.Value() == target {
		return []string{tr.Value()}
	}

	for _, child := range tr.Children() {
		if path := pathFromRoot(child, target); path != nil {
			return append([]string{tr.Value()}, path...)
		}
	}

	return nil
}
