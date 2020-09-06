package main

type Node struct {
	Character string
	Frequency int
	Left      *Node
	Right     *Node
}

type nodeArr []*Node

func (n nodeArr) Len() int {
	return len(n)
}
func (n nodeArr) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n nodeArr) Less(i, j int) bool {
	return n[j].Frequency < n[i].Frequency
}
