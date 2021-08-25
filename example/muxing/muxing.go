package main

import (
	"fmt"
	"github.com/HammerMax/goav/avcodec"
	"github.com/HammerMax/goav/avformat"
	"github.com/HammerMax/goav/avutil"
	"github.com/HammerMax/goav/swresample"
	"github.com/HammerMax/goav/swscale"
	"math"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s output_file\n" +
		"API example program to output a media file with libavformat.\n" +
		"This program generates a synthetic audio and video stream, encodes and\n" +
		"muxes them into a file named output_file.\n" +
		"The output format is automatically guessed according to the file extension.\n" +
		"Raw images can also be output by using '%%d' in the filename.\n" +
		"\n", os.Args[0]);
		return
	}

	filename := os.Args[1]

	var formatCtx *avformat.Context
	err := avformat.AvformatAllocOutputContext2(&formatCtx, nil, "", filename)
	if err != nil {
		panic(err)
	}

	ofmt := formatCtx.Oformat()

	videoSt := OutputStream{}
	audioSt := OutputStream{}
	videoCodec := &avcodec.Codec{}
	audioCodec := &avcodec.Codec{}
	if ofmt.VideoCodec() != avcodec.AV_CODEC_ID_NONE {
		AddStream(&videoSt, formatCtx, &videoCodec, ofmt.VideoCodec())
	}

	if ofmt.AudioCodec() != avcodec.AV_CODEC_ID_NONE {
		AddStream(&audioSt, formatCtx, &audioCodec, ofmt.AudioCodec())
	}

	OpenVideo(formatCtx, videoCodec, &videoSt, nil)
	OpenAudio(formatCtx, audioCodec, &audioSt, nil)

	formatCtx.AvDumpFormat(0, filename, 1)

	err = avformat.AvIOOpen(formatCtx.Pb2(), filename, avformat.AVIO_FLAG_WRITE)
	if err != nil {
		panic(err)
	}

	err = formatCtx.AvformatWriteHeader(nil)
	if err != nil {
		panic(err)
	}

	encodeVideo := 0
	encodeAudio := 0

	for encodeVideo == 0 || encodeAudio == 0 {
		if encodeVideo == 0 && (encodeAudio != 0 || avutil.AvCompareTs(videoSt.nextPts, videoSt.enc.GetTimeBase(), audioSt.nextPts, audioSt.enc.GetTimeBase()) <= 0) {
			encodeVideo = writeVideoFrame(formatCtx, &videoSt)
		} else {
			encodeAudio = writeAudioFrame(formatCtx, &audioSt)
		}
	}

	formatCtx.AvWriteTrailer()
}

type OutputStream struct {
	st *avformat.Stream
	enc *avcodec.Context

	nextPts int
	samplesCount int

	frame *avutil.Frame
	tmpFrame *avutil.Frame

	t, tincr, tincr2 float64

	swrCtx *swresample.Context
	swsCtx *swscale.Context
}

func AddStream(ost *OutputStream, fmtCtx *avformat.Context, codec **avcodec.Codec, codecId avcodec.CodecId) {
	*codec = avcodec.AvcodecFindEncoder(codecId)
	if *codec == nil {
		panic("find encoder error")
	}

	ost.st = fmtCtx.AvformatNewStream(nil)
	if ost.st == nil {
		panic("new stream error")
	}

	ost.st.SetID(int(fmtCtx.NbStreams())-1)
	c := (*codec).AvcodecAllocContext3()
	if c == nil {
		panic("alloc codec context error")
	}
	
	ost.enc = c

	switch (*codec).Type() {
	case avutil.AVMEDIA_TYPE_AUDIO:
		c.SetSampleFmt(avutil.AV_SAMPLE_FMT_FLTP)
		c.SetBitRate(64000)
		c.SetSampleRate(44100)
		c.SetChannels(avutil.AvGetChannelLayoutNbChannels(c.ChannelLayout()))
		c.SetChannelLayout(avutil.AV_CH_LAYOUT_STEREO)
		c.SetChannels(avutil.AvGetChannelLayoutNbChannels(c.ChannelLayout()))
		ost.st.SetTimeBase(avutil.NewRational(1, c.SampleRate()))

	case avutil.AVMEDIA_TYPE_VIDEO:
		c.SetCodecId(codecId)
		c.SetBitRate(400000)
		c.SetWidth(352)
		c.SetHeight(288)
		ost.st.SetTimeBase(avutil.NewRational(1, 25))
		c.SetTimeBase(ost.st.TimeBase())
		c.SetGopSize(12)
		c.SetPixFmt(avutil.AV_PIX_FMT_YUV420P)
	}
}

func OpenAudio(fmtCtx *avformat.Context, codec *avcodec.Codec, ost *OutputStream, arg *avutil.Dictionary) {
	c := ost.enc

	err := ost.enc.AvcodecOpen2(codec, nil)
	if err != nil {
		panic(err)
	}

	ost.t = 0
	ost.tincr =2 * math.Pi * 110 / float64(c.SampleRate())
	ost.tincr2 =  2 * math.Pi * 110 / float64(c.SampleRate()) / float64(c.SampleRate())

	nbSamples := c.FrameSize()

	ost.frame = AllocAudioFrame(c.SampleFmt(), c.ChannelLayout(), c.SampleRate(), nbSamples)
	ost.tmpFrame = AllocAudioFrame(avutil.AV_SAMPLE_FMT_S16, c.ChannelLayout(), c.SampleRate(), nbSamples)

	err = ost.st.CodecParameters().AvCodecParametersFromContext(c)
	if err != nil {
		panic(err)
	}

	ost.swrCtx = swresample.SwrAlloc()
	if ost.swrCtx == nil {
		panic("swr alloc error")
	}

	err = avutil.AvOptSetInt(ost.swrCtx, "in_channel_count", c.Channels(), 0)
	if err != nil {
		panic(err)
	}
	_ = avutil.AvOptSetInt(ost.swrCtx, "in_sample_rate", c.SampleRate(), 0)
	_ = avutil.AvOptSetSampleFmt(ost.swrCtx, "in_sample_fmt", avutil.AV_SAMPLE_FMT_S16, 0)
	_ = avutil.AvOptSetInt(ost.swrCtx, "out_channel_count", c.Channels(), 0)
	_ = avutil.AvOptSetInt(ost.swrCtx, "out_sample_rate", c.SampleRate(), 0)
	_ = avutil.AvOptSetSampleFmt(ost.swrCtx, "out_sample_fmt", c.SampleFmt(), 0)

	err = ost.swrCtx.SwrInit()
	if err != nil {
		panic(err)
	}
}

func AllocAudioFrame(sampleFmt avutil.SampleFormat, channelLayout int, sampleRage int, nbSamples int) *avutil.Frame {
	frame := avutil.AvFrameAlloc()
	if frame == nil {
		panic("alloc frame error")
	}

	frame.SetFormat(int32(sampleFmt))
	frame.SetChannelLayout(channelLayout)
	frame.SetSampleRate(sampleRage)
	frame.SetNbSamples(nbSamples)

	if nbSamples != 0 {
		err := frame.AvFrameGetBuffer(0)
		if err != nil {
			panic(err)
		}
	}
	return frame
}

func OpenVideo(fmtCtx *avformat.Context, codec *avcodec.Codec, ost *OutputStream, arg *avutil.Dictionary) {
	c := ost.enc

	err := c.AvcodecOpen2(codec, nil)
	if err != nil {
		panic(err)
	}

	ost.frame = allocPicture(c.PixFmt(), c.Width(), c.Height())
	if ost.frame == nil {
		panic(err)
	}

	err = ost.st.CodecParameters().AvCodecParametersFromContext(c)
	if err != nil {
		panic(err)
	}
}

func allocPicture(pixFmt avutil.PixelFormat, width, height int32) *avutil.Frame {
	picture := avutil.AvFrameAlloc()
	if picture == nil {
		panic("alloc frame error")
	}

	picture.SetFormat(int32(pixFmt))
	picture.SetWidth(width)
	picture.SetHeight(height)

	err := picture.AvFrameGetBuffer(0)
	if err != nil {
		panic(err)
	}

	return picture
}

func writeVideoFrame(formatCtx *avformat.Context, ost *OutputStream) int {
	return writeFrame(formatCtx, ost.enc, ost.st, getVideoFrame(ost))
}

func getVideoFrame(ost *OutputStream) *avutil.Frame {
	c := ost.enc

	if avutil.AvCompareTs(ost.nextPts, c.GetTimeBase(), 10, avutil.NewRational(1, 1)) > 0 {
		return nil
	}

	err := ost.frame.AvFrameMakeWritable()
	if err != nil {
		panic(err)
	}

	fillYuvImage(ost.frame, ost.nextPts, int(c.Width()), int(c.Height()))

	ost.frame.SetPts(ost.nextPts)
	ost.nextPts++
	return ost.frame
}

func fillYuvImage(frame *avutil.Frame, frameIndex, width, height int) {
	i := frameIndex

	for y:=0; y<height; y++ {
		for x:=0; x<width; x++ {
			frame.SetDataSimple(0, y*int(frame.LineSize()[0]) + x, uint8(x + y + i * 3))
		}
	}
	for y:=0; y<height/2; y++ {
		for x:=0; x<width/2; x++ {
			frame.SetDataSimple(1, y * int(frame.LineSize()[1]) + x, uint8(128 + y + i * 2))
			frame.SetDataSimple(2, y * int(frame.LineSize()[2]) + x, uint8(64 + x + i * 5))
		}
	}
}

func writeFrame(formatCtx *avformat.Context, c *avcodec.Context, st *avformat.Stream, frame *avutil.Frame) int {
	err := c.AvcodecSendFrame(frame)
	if err != nil {
		panic(err)
	}

	for err == nil {
		pkt := avcodec.AvPacketAlloc()

		err = c.AvcodecReceivePacket(pkt)
		if err == avutil.ErrEAGAIN || err == avutil.ErrEOF {
			break
		} else if err != nil {
			panic(err)
		}

		pkt.AvPacketRescaleTs(c.GetTimeBase(), st.TimeBase())
		pkt.SetStreamIndex(st.Index())

		err = formatCtx.AvInterleavedWriteFrame(pkt)
		if err != nil {
			panic(err)
		}
	}

	if err == avutil.ErrEOF {
		return 1
	}
	return 0
}

func writeAudioFrame(formatCtx *avformat.Context, ost *OutputStream) int {
	return -1
	c := ost.enc
	frame := getAudioFrame(ost)
	err := frame.AvFrameMakeWritable()
	if err != nil {
		panic(err)
	}
	return writeFrame(formatCtx, c, ost.st, frame)
}

func getAudioFrame(ost *OutputStream) *avutil.Frame {
	frame :=ost.tmpFrame
	frame.SetPts(ost.nextPts)
	ost.nextPts += frame.NbSamples()
	return frame
}
