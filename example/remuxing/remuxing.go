package main

import (
	"fmt"
	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avformat"
	"github.com/giorgisio/goav/avutil"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("usage: %s input output\n API example program to remux a media file with libavformat and libavcodec.\n The output format is guessed according to the file extension.\n", os.Args[0])
		return
	}

	var err error
	inFilename := os.Args[1]
	outFilename := os.Args[2]

	var ifmtCtx *avformat.Context
	if avformat.AvformatOpenInput(&ifmtCtx, inFilename, nil, nil) < 0 {
		panic("open input file error")
	}

	if ifmtCtx.AvformatFindStreamInfo(nil) < 0 {
		panic("find stream error")
	}

	ifmtCtx.AvDumpFormat(0, inFilename, 0)

	var ofmtCtx *avformat.Context
	avformat.AvformatAllocOutputContext2(&ofmtCtx, nil, "", outFilename)
	if ofmtCtx == nil {
		panic("out file format context error")
	}

	streamIndex := 0
	streamMapping := map[int]int{}


	for i:=0; i<int(ifmtCtx.NbStreams()); i++ {
		var outStream *avformat.Stream
		var inStream = ifmtCtx.Streams()[i]
		inCodecpar := inStream.CodecParameters()

		if inCodecpar.AvCodecGetType() != avutil.AVMEDIA_TYPE_AUDIO &&
			inCodecpar.AvCodecGetType() != avutil.AVMEDIA_TYPE_VIDEO &&
			inCodecpar.AvCodecGetType() != avutil.AVMEDIA_TYPE_SUBTITLE {
			streamMapping[i] = -1
			continue
		}

		streamMapping[i] = streamIndex
		streamIndex++

		outStream = ofmtCtx.AvformatNewStream(nil)
		if outStream == nil {
			panic("out new stream error")
		}

		if avcodec.AvCodecParametersCopy(outStream.CodecParameters(), inCodecpar) < 0 {
			panic("copy error")
		}

		outStream.CodecParameters().SetCodecTag(0)
	}

	ofmtCtx.AvDumpFormat(0, outFilename, 1)

	ofmt := ofmtCtx.Oformat()
	if ofmt.Flags() & avformat.AVFMT_NOFILE == 0 {
		err := avformat.AvIOOpen(ofmtCtx.Pb2(), outFilename, avformat.AVIO_FLAG_WRITE)
		if err != nil {
			panic(err)
		}
	}

	if err = ofmtCtx.AvformatWriteHeader(nil); err != nil {
		panic(err)
	}

	var packet = &avcodec.Packet{}
	for {
		var inStream, outStream *avformat.Stream

		ret := ifmtCtx.AvReadFrame(packet)
		if ret < 0 {
			break
		}

		inStream = ifmtCtx.Streams()[packet.StreamIndex()]
		if packet.StreamIndex() >= len(streamMapping) || streamMapping[packet.StreamIndex()] < 0 {
			packet.AvPacketUnref()
			continue
		}

		packet.SetStreamIndex(streamMapping[packet.StreamIndex()])
		outStream = ofmtCtx.Streams()[packet.StreamIndex()]

		// copy packet
		packet.SetPts(int64(avutil.AvRescaleQRnd(int(packet.Pts()), inStream.TimeBase(), outStream.TimeBase(), avutil.AV_ROUND_INF|avutil.AV_ROUND_PASS_MINMAX)))
		packet.SetDts(int64(avutil.AvRescaleQRnd(int(packet.Dts()), inStream.TimeBase(), outStream.TimeBase(), avutil.AV_ROUND_INF|avutil.AV_ROUND_PASS_MINMAX)))
		packet.SetDuration(int64(avutil.AvRescaleQ(int(packet.Duration()), inStream.TimeBase(), outStream.TimeBase())))
		packet.SetPos(-1)

		ofmtCtx.AvInterleavedWriteFrame(packet)

		packet.AvPacketUnref()
	}

	ofmtCtx.AvWriteTrailer()

	//ofmtCtx.Pb().Close()
	ifmtCtx.AvformatCloseInput()
	ofmtCtx.AvformatFreeContext()
}
