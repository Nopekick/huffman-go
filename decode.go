package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"sort"
)

type Decoder struct {
	InputFile  string
	OutputFile string
	Head       *Node
	Content    string
	Padding    int
}

func (d *Decoder) recoverTree() {
	file, err := os.Open(d.InputFile)
	if err != nil {
		fmt.Println("Could not open input file: " + d.InputFile)
		os.Exit(1)
	}

	//read number of nodes
	var nodeNumber uint8
	err = binary.Read(file, binary.LittleEndian, &nodeNumber)

	//read char:frequency pairs and retrieve nodes
	var list nodeArr
	for i := uint8(0); i < nodeNumber; i++ {
		var char byte
		var frequency uint32
		binary.Read(file, binary.LittleEndian, &char)
		binary.Read(file, binary.LittleEndian, &frequency)
		//fmt.Println(char, frequency)
		list = append(list, &Node{Character: string(char), Frequency: int(frequency)})
	}
	sort.Sort(nodeArr(list))

	//generate Huffman tree
	var least1, least2 *Node
	for len(list) > 1 {
		least1 = list[len(list)-1]
		list = list[:len(list)-1]
		least2 = list[len(list)-1]
		list = list[:len(list)-1]
		parent := &Node{Frequency: least1.Frequency + least2.Frequency}
		parent.Left = least2
		parent.Right = least1
		list = append(list, parent)

		sort.Sort(nodeArr(list))
	}
	root := list[len(list)-1]
	d.Head = root

	//get padding length
	var difference uint8
	binary.Read(file, binary.LittleEndian, &difference)

}

func (d *Decoder) decode() {

}
