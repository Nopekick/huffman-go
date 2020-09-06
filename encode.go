package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

type Encoder struct {
	InputFile  string
	OutputFile string
	Head       *Node
	Frequency  map[string]int
	Bitmap     map[string]string
	List       nodeArr
}

func (e *Encoder) generateTree() {
	data, err := ioutil.ReadFile(e.InputFile)

	if err != nil {
		fmt.Println("Could not open input file: " + e.InputFile)
		os.Exit(1)
	}

	//check if input file is empty
	fi, err := os.Stat(e.InputFile)
	if err != nil {
		fmt.Println("Could not open input file: " + e.InputFile)
		os.Exit(1)
	} else if fi.Size() == 0 {
		fmt.Println("Input file is empty. No need to compress anything")
		os.Exit(1)
	}

	//generate char frequency map
	reader := bufio.NewReader(strings.NewReader(string(data)))
	for {
		if c, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			if _, ok := e.Frequency[string(c)]; ok {
				e.Frequency[string(c)]++
			} else {
				e.Frequency[string(c)] = 1
			}
		}
	}

	for k, v := range e.Frequency {
		node := &Node{
			Character: k,
			Frequency: v,
			Left:      nil,
			Right:     nil,
		}
		e.List = append(e.List, node)
	}
	sort.Sort(nodeArr(e.List))
	for _, v := range e.List {
		fmt.Println(*v)
	}

	tmp := make(nodeArr, len(e.List))
	copy(tmp, e.List)

	var least1, least2 *Node
	for len(tmp) > 1 {
		least1 = tmp[len(tmp)-1]
		tmp = tmp[:len(tmp)-1]
		least2 = tmp[len(tmp)-1]
		tmp = tmp[:len(tmp)-1]
		parent := &Node{
			Frequency: least1.Frequency + least2.Frequency,
		}
		parent.Left = least2
		parent.Right = least1
		tmp = append(tmp, parent)

		sort.Sort(nodeArr(tmp))
	}
	root := tmp[len(tmp)-1]
	e.Head = root

	e.recHelper(e.Head, "")

	// for k, v := range e.Bitmap {
	// 	fmt.Println(k, v)
	// }
}

func (e *Encoder) recHelper(node *Node, built string) {
	if node.Left == nil && node.Right == nil {
		e.Bitmap[node.Character] = built
	}
	if node.Left != nil {
		e.recHelper(node.Left, built+"0")
	}
	if node.Right != nil {
		e.recHelper(node.Right, built+"1")
	}
}
