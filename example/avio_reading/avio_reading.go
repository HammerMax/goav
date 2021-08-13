package main

import (
	"fmt"
	"github.com/giorgisio/goav/avformat"
	"github.com/giorgisio/goav/avutil"
	"os"
)

type opaque struct {
	buf []byte
}

func ReadPacket(o interface{}, data []byte) int {
	op := o.(*opaque)

	n := copy(data, op.buf)
	op.buf = op.buf[n:]
	return n
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <input_file>", os.Args[0])
		return
	}

	inputFileName := os.Args[1]

	var buffer []byte
	var bufferSize int

	if avutil.AvFileMap(inputFileName, &buffer, &bufferSize, 0, nil) < 0 {
		panic("file map error")
	}

	formatCtx := avformat.AvformatAllocContext()
	if formatCtx == nil {
		panic("format alloc context error")
	}

	opaqueBuf := opaque{buf: buffer}

	buf := make([]byte, 4096)
	avformat.AVIOAllocContext(buf, 0, &opaqueBuf, ReadPacket, nil, nil)
}
