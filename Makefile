
all: main

main:
	go build -o main

# encode:
# 	go run *.go -e input/test1.txt test/output1

# decode:
# 	go run *.go -d test/output1 test/decode1

clean:
	rm -f test/output1 test/output2 test/decode1 test/decode2

tests: test-1 test-2

test-1:
	go run *.go -e test/input1 test/output1	
	go run *.go -d test/output1 test/decode1

test-2:
	go run *.go -e test/input2 test/output2	
	go run *.go -d test/output2 test/decode2