package main

import (
	"fmt"
	"os"
)

func main() {
	fileArgs := os.Args[1:]

	if len(fileArgs) != 3 {
		printUsage()
	}

	//fmt.Println(fileArgs)
	if fileArgs[0] == "-e" {
		encoder := Encoder{
			InputFile:  fileArgs[1],
			OutputFile: fileArgs[2],
			Frequency:  map[string]int{},
			Bitmap:     map[string]string{},
			List:       []*Node{},
		}
		encoder.generateTree()
		encoder.encode()
		printFileSize(fileArgs[1], fileArgs[2])
	} else if fileArgs[0] == "-d" {
		decoder := Decoder{
			InputFile:  fileArgs[1],
			OutputFile: fileArgs[2],
			Head:       nil,
			Content:    "",
		}
		decoder.recoverTree()
		decoder.decode()
	} else {
		printUsage()
	}

}

func printUsage() {
	fmt.Println("Usage: -e/-d input-file output-file")
	fmt.Println("-e: encode the input-file, put encoded result into output-file")
	fmt.Println("-d: decode the input-file, put decoded result into output-file")
	os.Exit(1)
}

func printFileSize(in string, out string) {
	input, _ := os.Stat(in)
	output, _ := os.Stat(out)
	//should probably add error checking here

	fmt.Println("Original file size (bytes): ", input.Size())
	fmt.Println("Compressed file size (bytes): ", output.Size())
}
