package main

import (
	"fmt"
	"github.com/HammerMax/goav/avcodec"
	"github.com/HammerMax/goav/avformat"
	"github.com/HammerMax/goav/avutil"
	"os"
	"strconv"
)

type decoder struct {
	filename string
	formatCtx *avformat.Context

	packet *avcodec.Packet
	frame *avutil.Frame

	files map[int]*avformat.Context // 每个流对应一个outputFormat
}

func newDecoder(filename string) *decoder {
	formatCtx := avformat.AvformatAllocContext()
	err := avformat.AvformatOpenInput(&formatCtx, filename, nil, nil)
	if err != nil {
		panic(err)
	}

	err = formatCtx.AvformatFindStreamInfo(nil)
	if err != nil {
		panic(err)
	}

	packet := avcodec.AvPacketAlloc()
	if packet == nil {
		panic("alloc packet error")
	}

	frame := avutil.AvFrameAlloc()
	if frame == nil {
		panic("alloc frame error")
	}

	d := &decoder{filename: filename, formatCtx: formatCtx, packet: packet, frame: frame}
	d.files = map[int]*avformat.Context{}

	return d
}

func (d *decoder) Decode() {
	var err error
	for _, stream := range d.formatCtx.Streams() {
		fmt.Println(stream)
		outFilename := d.filename + "_output_" + strconv.Itoa(stream.Index()) + ".mp4"

		octx := avformat.AvformatAllocContext()
		err = avformat.AvformatAllocOutputContext2(&octx, nil, "", outFilename)
		if err != nil {
			panic(err)
		}

		outStream := octx.AvformatNewStream(nil)
		if outStream == nil {
			panic("new stream error")
		}

		err = avcodec.AvCodecParametersCopy(outStream.CodecParameters(), stream.CodecParameters())
		if err != nil {
			panic(err)
		}

		err = avformat.AvIOOpen(octx.Pb2(), outFilename, avformat.AVIO_FLAG_WRITE)
		if err != nil {
			panic(err)
		}

		err = octx.AvformatWriteHeader(nil)
		if err != nil {
			panic(err)
		}

		d.files[stream.Index()] = octx
	}

	for {
		err = d.formatCtx.AvReadFrame(d.packet)
		if err == avutil.ErrEOF {
			break
		} else if err != nil {
			panic(err)
		}
		ofmt := d.files[d.packet.StreamIndex()]

		d.packet.SetStreamIndex(0)
		d.packet.SetPts(int64(avutil.AvRescaleQRnd(int(d.packet.Pts()), d.formatCtx.Streams()[d.packet.StreamIndex()].TimeBase(), ofmt.Streams()[0].TimeBase(), avutil.AV_ROUND_NEAR_INF|avutil.AV_ROUND_PASS_MINMAX)))
		d.packet.SetDts(int64(avutil.AvRescaleQRnd(int(d.packet.Dts()), d.formatCtx.Streams()[d.packet.StreamIndex()].TimeBase(), ofmt.Streams()[0].TimeBase(), avutil.AV_ROUND_NEAR_INF|avutil.AV_ROUND_PASS_MINMAX)))
		d.packet.SetDuration(int64(avutil.AvRescaleQ(int(d.packet.Duration()), d.formatCtx.Streams()[d.packet.StreamIndex()].TimeBase(), ofmt.Streams()[0].TimeBase())))
		d.packet.SetPos(-1)

		err = ofmt.AvInterleavedWriteFrame(d.packet)
		if err != nil {
			panic(err)
		}

		d.packet.AvPacketUnref()
	}

	for _, o := range d.files {
		err = o.AvWriteTrailer()
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage %s <input_filename>", os.Args[0])
		return
	}

	filename := os.Args[1]

	dec := newDecoder(filename)
	dec.Decode()
}
