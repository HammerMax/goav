package main

import (
	"fmt"
	"github.com/HammerMax/goav/avcodec"
	"github.com/HammerMax/goav/avformat"
	"github.com/HammerMax/goav/avutil"
	"os"
	"strconv"
)

type mergeOutput struct {
	filenames []string
	outputFormatContext *avformat.Context
}

const (
	outputFilename = "output.mp4"
)

func newMergeOutput(filenames []string) *mergeOutput {
	outputFormatContext := avformat.AvformatAllocContext()
	err := avformat.AvformatAllocOutputContext2(&outputFormatContext, nil, "", outputFilename)
	if err != nil {
		panic(err)
	}

	err = avformat.AvIOOpen(outputFormatContext.Pb2(), outputFilename, avformat.AVIO_FLAG_WRITE)
	if err != nil {
		panic(err)
	}

	return &mergeOutput{filenames: filenames, outputFormatContext: outputFormatContext}
}

func (m *mergeOutput) Merge() {
	var err error

	var streamIndexMap = map[string]int{} // 用于存储inputStreamIndex和outputStreamIndex的对应关系
	for _, filename := range m.filenames {
		formatContext := avformat.AvformatAllocContext()
		err = avformat.AvformatOpenInput(&formatContext, filename, nil, nil)
		if err != nil {
			panic(err)
		}

		err = formatContext.AvformatFindStreamInfo(nil)
		if err != nil {
			panic(err)
		}

		for _, stream := range formatContext.Streams() {
			outStream := m.outputFormatContext.AvformatNewStream(nil)
			if outStream == nil {
				panic("new stream error")
			}

			err = avcodec.AvCodecParametersCopy(outStream.CodecParameters(), stream.CodecParameters())
			if err != nil {
				panic(err)
			}

			streamIndexMap[filename + strconv.Itoa(stream.Index())] = outStream.Index()
		}
	}

	err = m.outputFormatContext.AvformatWriteHeader(nil)
	if err != nil {
		panic(err)
	}

	for _, filename := range m.filenames {
		formatContext := avformat.AvformatAllocContext()
		err = avformat.AvformatOpenInput(&formatContext, filename, nil, nil)
		if err != nil {
			panic(err)
		}

		packet := avcodec.AvPacketAlloc()
		for {
			err = formatContext.AvReadFrame(packet)
			if err == avutil.ErrEOF {
				break
			} else if err != nil {
				panic(err)
			}

			inStream := formatContext.Streams()[packet.StreamIndex()]
			outStream := m.outputFormatContext.Streams()[streamIndexMap[filename + strconv.Itoa(packet.StreamIndex())]]

			packet.SetStreamIndex(outStream.Index())
			packet.SetPts(int64(avutil.AvRescaleQRnd(int(packet.Pts()), inStream.TimeBase(), outStream.TimeBase(), avutil.AV_ROUND_NEAR_INF|avutil.AV_ROUND_PASS_MINMAX)))
			packet.SetDts(int64(avutil.AvRescaleQRnd(int(packet.Dts()), inStream.TimeBase(), outStream.TimeBase(), avutil.AV_ROUND_NEAR_INF|avutil.AV_ROUND_PASS_MINMAX)))
			packet.SetDuration(int64(avutil.AvRescaleQ(int(packet.Duration()), inStream.TimeBase(), outStream.TimeBase())))
			packet.SetPos(-1)

			err = m.outputFormatContext.AvInterleavedWriteFrame(packet)
			if err != nil {
				panic(err)
			}

			packet.AvPacketUnref()
		}
	}

	err = m.outputFormatContext.AvWriteTrailer()
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <input> <input> ..." ,os.Args[0])
		return
	}

	merge := newMergeOutput(os.Args[1:])
	merge.Merge()
}
