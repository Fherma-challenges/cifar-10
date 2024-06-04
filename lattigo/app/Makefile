# Flags default values
sk = temps/sk.bin
cc = temps/cc.bin
key_public = temps/pub.bin
key_eval = temps/mult.bin
input = temps/in.bin
output = temps/out.bin

test-all: 
	go run setup.go --sk=$(sk) --cc=$(cc) --key_public=$(key_public) --key_eval=$(key_eval) --input=$(input)
	go run main.go --cc=$(cc) --key_public=$(key_public) --key_eval=$(key_eval) --input=$(input) --output=$(output)
	go run verify.go --sk=$(sk) --cc=$(cc) --input=$(input) --output=$(output)
	go run clean.go
	go clean

setup:
	go run setup.go --sk=$(sk) --cc=$(cc) --key_public=$(key_public) --key_eval=$(key_eval) --input=$(input)

solution:
	go run main.go --cc=$(cc) --key_public=$(key_public) --key_eval=$(key_eval) --input=$(input) --output=$(output)
	go run verify.go --sk=$(sk) --cc=$(cc) --input=$(input) --output=$(output)

clean:
	go run clean.go
	go clean
