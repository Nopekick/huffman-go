# Huffman Encoder/Decoder (in Golang)

## Descriptions of Huffman Encoding

Wikipedia: https://en.wikipedia.org/wiki/Huffman_coding

Nachenberg's slides: https://drive.google.com/file/d/1z2lC4vjmO98x4LbYCfHSmc2N5tnSz6YT/view?usp=sharing

See this implementation in C++: https://github.com/Nopekick/huffman 

## Usage
You have two options. Either use go run as below
```
go run *.go -e/-f -inputfile -outputfile
```
or build the binary first, and execute it
```
go build -o main
./main -e/-f -inputfile -outputfile
```


## Tests 
To automatically run the program on the given files
```
make tests
```
```
make test-1
```
```
make test-2
```

This program currently assumes that the input file for decoding
was encoded by the same program. Using a random file as the input file
for decoding will cause undefined behavior.

Additionally, for small input files, the resulting compressed output 
file may be larger. The compression is more effective on 
larger text files. 