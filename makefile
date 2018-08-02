OUTPUT_BIN="build/gchip"

all:
	go build -o ${OUTPUT_BIN} ./main.go
	chmod +x ${OUTPUT_BIN}
run:
	go run ./main.go