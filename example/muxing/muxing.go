package main

import (
	"fmt"
	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avformat"
	"github.com/giorgisio/goav/avutil"
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


}

type OutputStream struct {
	st *avformat.Stream
	enc *avcodec.Context

	nextPts int
	samplesCount int

	frame *avutil.Frame
	tmpFrame *avutil.Frame

	t, tincr, tincr2 float64
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
		c.SetPixFmt(avcodec.AV_PIX_FMT_YUV420P)
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
}

func AllocAudioFrame(sampleFmt avutil.SampleFormat, channelLayout int, sampleRage int, nbSamples int) *avutil.Frame {
	frame := avutil.AvFrameAlloc()
	if frame == nil {
		panic("alloc frame error")
	}

	frame.SetFormat(sampleFmt)
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
