package main

import (
	"encoding/binary"
	"fmt"
	"io"
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
	d.Padding = int(difference)

	//get file contents as binary string
	content := ""
	var numRep uint8
	for {
		if err = binary.Read(file, binary.LittleEndian, &numRep); err != nil {
			break
		} else {
			num := fmt.Sprintf("%08b", numRep)
			content += num
			//fmt.Println(num, numRep)
		}
	}
	d.Content = content
	//fmt.Println(len(content), content)
}

func (d *Decoder) decode() {
	var buildOutput, content string

	//substract padding from beginning of last byte of binary file content
	lastByte := d.Content[len(d.Content)-8:]
	if d.Padding != 8 {
		lastByte = lastByte[d.Padding:]
	}
	content = d.Content[0:len(d.Content)-8] + string(lastByte)

	pos := 0
	temp := d.Head
	for pos < len(content) {
		if content[pos] == '0' {
			if temp.Left != nil {
				temp = temp.Left
				if temp.Left == nil && temp.Right == nil {
					buildOutput += temp.Character
					temp = d.Head
				}
			}
		} else if content[pos] == '1' {
			if temp.Right != nil {
				temp = temp.Right
				if temp.Left == nil && temp.Right == nil {
					buildOutput += temp.Character
					temp = d.Head
				}
			}
		}
		pos++
	}
	fmt.Println(buildOutput)
	output, _ := os.Create(d.OutputFile)
	_, err := io.WriteString(output, buildOutput)
	if err != nil {
		fmt.Println(err)
	}
	output.Close()
}
