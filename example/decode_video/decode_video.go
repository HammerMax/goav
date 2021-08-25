package main

import (
	"fmt"
	"github.com/HammerMax/goav/avcodec"
	"github.com/HammerMax/goav/avutil"
	"os"
)

func pgmSave(data []byte, wrap, x, y int32, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprintf(f, "P5\n%d %d\n%d\n", x, y, 255)
	if err != nil {
		panic(err)
	}

	for i:=int32(0); i<y; i++ {
		_, err := f.Write(data[i*wrap:i*wrap+x])
		if err != nil {
			panic(err)
		}
	}
	f.Close()
}

func decode(ctx *avcodec.Context, frame *avutil.Frame, packet *avcodec.Packet, filename string) {
	if ctx.AvcodecSendPacket(packet) != nil {
		panic("send packet error")
	}

	for {
		ret := ctx.AvcodecReceiveFrame(frame)
		fmt.Println("receive frame ret ", ret)
		if ret != nil {
			return
		}

		fmt.Println("saving frame", ctx.FrameNumber())

		pgmSave(frame.Data()[0], frame.LineSize()[0], frame.Width(), frame.Height(), fmt.Sprintf("%s-%d", filename, ctx.FrameNumber()))
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage:%s <input file> <output file>", os.Args[0])
		return
	}

	fileName := os.Args[1]
	outfilename := "frame"

	packet := avcodec.AvPacketAlloc()
	if packet == nil {
		panic("alloc packet error")
	}

	codec := avcodec.AvcodecFindDecoder(avcodec.AV_CODEC_ID_H264)
	if codec == nil {
		panic("can not find codec")
	}

	parser := avcodec.AvParserInit(codec.Id())
	if parser == nil {
		panic("parser is nil")
	}

	codecCtx := codec.AvcodecAllocContext3()
	if codecCtx == nil {
		panic("codecCtx is nil")
	}

	if codecCtx.AvcodecOpen2(codec, nil) != nil {
		panic("can not open")
	}

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	frame := avutil.AvFrameAlloc()
	if frame == nil {
		panic("frame alloc error")
	}

	for {
		data := make([]byte, 4096)
		dataSize, err := f.Read(data)
		if err != nil {
			break
		}

		for dataSize > 0 {
			ret := parser.AvParserParse2(codecCtx, packet.DataPoint() , packet.SizePoint(), data, dataSize, avutil.AV_NOPTS_VALUE, avutil.AV_NOPTS_VALUE, 0)
			if ret < 0 {
				panic("parser parse error")
			}

			data = data[ret:]
			dataSize -= ret

			if packet.Size() > 0 {
				decode(codecCtx, frame, packet, outfilename)
			}
		}
	}

	decode(codecCtx, frame, nil, outfilename)

	f.Close()
	parser.AvParserClose()
	codecCtx.AvcodecFreeContext()
	frame.AvFrameFree()
	packet.AvPacketFree()
}
