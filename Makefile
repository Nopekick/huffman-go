
all: encode decode

encode:
	go run *.go -e test/test1.txt test/output

decode:
	go run *.go -d test/output test/decoded.txt