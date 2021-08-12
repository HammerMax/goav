package main

import (
	"fmt"
	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avutil"
	"os"
)

func encode(encodeCtx *avcodec.Context, frame *avutil.Frame, packet *avcodec.Packet, outfile *os.File) {
	fmt.Printf("Send frame %d\n", frame.Pts())
	ret := encodeCtx.AvcodecSendFrame(frame)
	if ret < 0 {
		fmt.Println("error sending a frame for encoding")
		return
	}

	for {
		ret = encodeCtx.AvcodecReceivePacket(packet)
		if ret < 0 {
			fmt.Printf("end. ret:%d", ret)
			return
		}

		fmt.Printf("Write packet %d (size=%d)\n", packet.Pts(), packet.Size())
		_, err := outfile.Write(packet.DataSlice())
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <output file> <codec name>", os.Args[0])
		return
	}

	outputFileName := os.Args[1]
	codecName := os.Args[2]

	// codec 只是一个大类
	codec := avcodec.AvcodecFindEncoderByName(codecName)
	if codec == nil {
		fmt.Println("codec not found")
		return
	}

	// codecCtx 是一个实例
	codecCtx := codec.AvcodecAllocContext3()
	if codecCtx == nil {
		fmt.Println("codec context alloc error")
		return
	}

	pkt := avcodec.AvPacketAlloc()
	if pkt == nil {
		fmt.Println("packet alloc error")
		return
	}

	codecCtx.SetBitRate(400000)
	codecCtx.SetWidth(352)
	codecCtx.SetHeight(288)
	codecCtx.SetTimebase(1, 25)
	codecCtx.SetFramerate(25, 1)
	codecCtx.SetPixFmt(avcodec.AV_PIX_FMT_YUV420P)

	if codec.Id() == avcodec.AV_CODEC_ID_H264 {
		codecCtx.AvOptSet("preset", "slow", 0)
	}

	if codecCtx.AvcodecOpen2(codec, nil) < 0 {
		fmt.Println("codec open error")
		return
	}

	f, err := os.OpenFile(outputFileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file error")
		return
	}

	frame := avutil.AvFrameAlloc()
	if frame == nil {
		fmt.Println("alloc frame error")
		return
	}

	frame.SetFormat(int32(codecCtx.PixFmt()))
	frame.SetWidth(codecCtx.Width())
	frame.SetHeight(codecCtx.Height())

	bufferNumber := frame.AvFrameGetBuffer(0)
	if bufferNumber < 0 {
		fmt.Println("could not allocate the video frame data")
		return
	}

	// encode 1 second of video
	for i:=0; i<25; i++ {
		if frame.AvFrameMakeWritable() < 0 {
			fmt.Println("not writable")
			return
		}

		// prepare a dummy image
		for y := int32(0); y < codecCtx.Height(); y++ {
			for x := int32(0); x < codecCtx.Width(); x++ {
				frame.SetDataSimple(0, int(y * frame.LineSize()[0] + x), uint8(x + y + int32(i * 3)))
			}
		}

		for y := int32(0); y < codecCtx.Height()/2; y++ {
			for x :=int32(0); x < codecCtx.Width()/2; x++ {
				frame.SetDataSimple(1, int(y * frame.LineSize()[1] + x), uint8(128 + y + int32(i * 2)))
				frame.SetDataSimple(2, int(y * frame.LineSize()[2] + x), uint8(64 + x + int32(i * 5)))
			}
		}

		frame.SetPts(i)

		encode(codecCtx, frame, pkt, f)
	}

	if codec.Id() == avcodec.AV_CODEC_ID_MPEG1VIDEO || codec.Id() == avcodec.AV_CODEC_ID_MPEG2VIDEO {
		endcode := []byte{0, 0, 1, 0xb7}
		f.Write(endcode)
	}
	f.Close()

	codecCtx.AvcodecFreeContext()
	frame.AvFrameFree()
	pkt.AvPacketFree()
}
