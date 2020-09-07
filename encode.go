package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
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

	tmp := make(nodeArr, len(e.List))
	copy(tmp, e.List)

	var least1, least2 *Node
	for len(tmp) > 1 {
		least1 = tmp[len(tmp)-1]
		tmp = tmp[:len(tmp)-1]
		least2 = tmp[len(tmp)-1]
		tmp = tmp[:len(tmp)-1]
		parent := &Node{Frequency: least1.Frequency + least2.Frequency}
		parent.Left = least2
		parent.Right = least1
		tmp = append(tmp, parent)

		sort.Sort(nodeArr(tmp))
	}
	root := tmp[len(tmp)-1]
	e.Head = root

	e.recHelper(e.Head, "")

	//Print (char: bitstring) pairs
	// for k, v := range e.Bitmap {
	// 	fmt.Println(k, v)
	// }

	//Print (char: frequency) pairs
	// for k, v := range e.Frequency {
	// 	fmt.Println(k, v)
	// }
}

func (e *Encoder) encode() {
	data, err := ioutil.ReadFile(e.InputFile)

	if err != nil {
		fmt.Println("Could not open input file: " + e.InputFile)
		os.Exit(1)
	}

	encodedFile := ""
	reader := bufio.NewReader(strings.NewReader(string(data)))
	for {
		if c, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			encodedFile += e.Bitmap[string(c)]
		}
	}
	//fmt.Println(len(encodedFile), encodedFile)

	//first 8 bits (1 byte) of file will be number of encoded data pairs
	numNodes := uint8(len(e.List))
	outf, _ := os.Create(e.OutputFile)
	binary.Write(outf, binary.LittleEndian, numNodes)

	//encoded char (character):int (frequency) pairs
	for _, node := range e.List {
		char := []byte(node.Character)[0]
		freq := uint32(node.Frequency)
		binary.Write(outf, binary.LittleEndian, char)
		binary.Write(outf, binary.LittleEndian, freq)
		//fmt.Println(string(char), freq)
	}

	//next 8 bits of file will be length of padding of last byte written to file
	var difference uint8 = uint8(8 - (len(encodedFile) % 8))
	binary.Write(outf, binary.LittleEndian, difference)

	var binRep uint8
	var firstByte string
	for len(encodedFile) > 0 {
		if len(encodedFile) >= 8 {
			firstByte = encodedFile[0:8]
			res, _ := strconv.ParseUint(firstByte, 2, 8)
			binRep = uint8(res)
			encodedFile = encodedFile[8:]
		} else {
			firstByte = encodedFile[0:len(encodedFile)]
			res, _ := strconv.ParseUint(firstByte, 2, 8)
			binRep = uint8(res)
			encodedFile = ""
		}
		//fmt.Println(firstByte, binRep)
		binary.Write(outf, binary.LittleEndian, binRep)
	}
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
