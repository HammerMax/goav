package main

import (
	"fmt"
	"github.com/HammerMax/goav/avcodec"
	"github.com/HammerMax/goav/avformat"
	"github.com/HammerMax/goav/avutil"
	"os"
	"strconv"
)

type clipItem struct {
	filename string
	start int // 以毫秒为单位
	duration int // 以毫秒为单位
}

func parseClipItem(args []string) []clipItem {
	var clipItemSlice []clipItem
	i := 0
	for len(args) - 1 > 3 * i {
		filename := args[3 * i + 1]
		start, err := strconv.Atoi(args[3 * i + 2])
		if err != nil {
			panic(err)
		}
		duration, err := strconv.Atoi(args[3 * i + 3])
		if err != nil {
			panic(err)
		}

		clipItemSlice = append(clipItemSlice, clipItem{
			filename: filename,
			start:    start,
			duration: duration,
		})
		i++
	}
	return clipItemSlice
}

const (
	outFilename = "out.mp4"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <input> <start> <duration> ...\n", os.Args[0])
		return
	}

	clipItemSlice := parseClipItem(os.Args)

	outFormatContext := avformat.AvformatAllocContext()
	err := avformat.AvformatAllocOutputContext2(&outFormatContext, nil, "", outFilename)
	if err != nil {
		panic(err)
	}

	err = avformat.AvIOOpen(outFormatContext.Pb2(), outFilename, avformat.AVIO_FLAG_WRITE)
	if err != nil {
		panic(err)
	}

	// 使用第一个文件作为模版复制流
	inFormatContext := avformat.AvformatAllocContext()
	err = avformat.AvformatOpenInput(&inFormatContext, clipItemSlice[0].filename, nil, nil)
	if err != nil {
		panic(err)
	}

	err = inFormatContext.AvformatFindStreamInfo(nil)
	if err != nil {
		panic(err)
	}

	for _, stream := range inFormatContext.Streams() {
		outStream := outFormatContext.AvformatNewStream(nil)
		if outStream == nil {
			panic("stream is nil")
		}

		err = avcodec.AvCodecParametersCopy(outStream.CodecParameters(), stream.CodecParameters())
		if err != nil {
			panic(err)
		}
	}

	err = outFormatContext.AvformatWriteHeader(nil)
	if err != nil {
		panic(err)
	}

	lastPts := 0
	lastDts := 0
	for _, item := range clipItemSlice {
		inFormatContext = avformat.AvformatAllocContext()
		err = avformat.AvformatOpenInput(&inFormatContext, item.filename, nil, nil)
		if err != nil {
			panic(err)
		}

		err = inFormatContext.AvSeekFrame(-1, int64(item.start * avutil.AV_TIME_BASE), avformat.AvseekFlagAny)
		if err != nil {
			panic(err)
		}

		startPts := lastPts
		startDts := lastDts
		dtsStartFrom := map[int]int{}
		ptsStartFrom := map[int]int{}
		packet := avcodec.AvPacketAlloc()
		for {
			err = inFormatContext.AvReadFrame(packet)
			if err == avutil.ErrEOF {
				break
			} else if err != nil {
				panic(err)
			}

			inStream := inFormatContext.Streams()[packet.StreamIndex()]
			outStream := outFormatContext.Streams()[packet.StreamIndex()]

			if inStream.CodecParameters().CodecType() != avutil.AVMEDIA_TYPE_VIDEO {
				packet.AvPacketUnref()
				continue
			}

			if dtsStartFrom[packet.StreamIndex()] == 0 {
				dtsStartFrom[packet.StreamIndex()] = int(packet.Dts())
			}
			if ptsStartFrom[packet.StreamIndex()] == 0 {
				ptsStartFrom[packet.StreamIndex()] = int(packet.Pts())
			}

			lastDts = startDts + int(packet.Dts()) - dtsStartFrom[packet.StreamIndex()] + 1
			lastPts = startPts + int(packet.Pts()) - ptsStartFrom[packet.StreamIndex()] + 1
			packet.SetDts(int64(avutil.AvRescaleQRnd(startDts + int(packet.Dts()) - dtsStartFrom[packet.StreamIndex()], inStream.TimeBase(), outStream.TimeBase(), avutil.AV_ROUND_INF|avutil.AV_ROUND_PASS_MINMAX)))
			packet.SetPts(int64(avutil.AvRescaleQRnd(startPts + int(packet.Pts()) - ptsStartFrom[packet.StreamIndex()], inStream.TimeBase(), outStream.TimeBase(), avutil.AV_ROUND_INF|avutil.AV_ROUND_PASS_MINMAX)))

			if packet.Pts() < 0 {
				packet.SetPts(0)
			}
			if packet.Dts() < 0 {
				packet.SetDts(0)
			}

			packet.SetDuration(int64(avutil.AvRescaleQ(int(packet.Duration()), inStream.TimeBase(), outStream.TimeBase())))
			packet.SetPos(-1)

			if packet.Pts() < packet.Dts() {
				continue
			}

			err = outFormatContext.AvInterleavedWriteFrame(packet)
			if err != nil {
				panic(err)
			}

			packet.AvPacketUnref()
		}
	}

	err = outFormatContext.AvWriteTrailer()
	if err != nil {
		panic(err)
	}
}
