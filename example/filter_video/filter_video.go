package main

import (
	"fmt"
	"github.com/HammerMax/goav/avcodec"
	"github.com/HammerMax/goav/avfilter"
	"github.com/HammerMax/goav/avformat"
	"github.com/HammerMax/goav/avutil"
	"os"
)

type Filter struct {
	inputFormatContext *avformat.Context
	inputVideoStream *avformat.Stream
	inputDecodeContext *avcodec.Context
	outputFormatContext *avformat.Context
	outputEncodeContext *avcodec.Context
	outputVideoStream *avformat.Stream

	filterInputs *avfilter.Inout
	filterOutputs *avfilter.Inout
	bufferSrc *avfilter.Filter
	bufferSrcContext *avfilter.Context
	bufferSink *avfilter.Filter
	bufferSinkContext *avfilter.Context

	filterGraph *avfilter.Graph
}

func newFilter(inputFilename, outputFilename string) *Filter {
	f := &Filter{}
	f.openInput(inputFilename)
	f.openOutput(outputFilename)
	f.initFilters("scale=100:100")
	return f
}

func (f *Filter) initFilters(filterDescription string) {
	f.bufferSrc = avfilter.AvfilterGetByName("buffer")
	f.bufferSink = avfilter.AvfilterGetByName("buffersink")
	f.filterInputs = avfilter.AvfilterInoutAlloc()
	f.filterOutputs = avfilter.AvfilterInoutAlloc()
	timeBase := f.inputVideoStream.TimeBase()

	f.filterGraph = avfilter.AvfilterGraphAlloc()

	if f.filterInputs == nil || f.filterOutputs == nil || f.filterGraph == nil {
		panic("alloc error")
	}

	args := fmt.Sprintf("video_size=%dx%d:pix_fmt=%d:time_base=%d/%d:pixel_aspect=%d/%d", f.inputDecodeContext.Width(), f.inputDecodeContext.Height(), f.inputDecodeContext.PixFmt(),
		timeBase.Num(), timeBase.Den(), f.inputDecodeContext.SampleAspectRatio().Num(), f.inputDecodeContext.SampleAspectRatio().Den())

	err := avfilter.AvfilterGraphCreateFilter(&f.bufferSrcContext, f.bufferSrc, "in", args, 0, f.filterGraph)
	if err != nil {
		panic(err)
	}

	err = avfilter.AvfilterGraphCreateFilter(&f.bufferSinkContext, f.bufferSink, "out", "", 0, f.filterGraph)
	if err != nil {
		panic(err)
	}

	f.filterOutputs.SetName("in")
	f.filterOutputs.SetFilterCtx(f.bufferSrcContext)
	f.filterOutputs.SetPadIdx(0)

	f.filterInputs.SetName("out")
	f.filterInputs.SetFilterCtx(f.bufferSinkContext)
	f.filterInputs.SetPadIdx(0)

	err = f.filterGraph.AvfilterGraphParsePtr(filterDescription, &f.filterInputs, &f.filterOutputs, nil)
	if err != nil {
		panic(err)
	}

	err = f.filterGraph.AvfilterGraphConfig(nil)
	if err != nil {
		panic(err)
	}
}

func (f *Filter) openInput(filename string) {
	f.inputFormatContext = avformat.AvformatAllocContext()
	err := avformat.AvformatOpenInput(&f.inputFormatContext, filename, nil, nil)
	if err != nil{
		panic(err)
	}

	err = f.inputFormatContext.AvformatFindStreamInfo(nil)
	if err != nil{
		panic(err)
	}

	for _, v := range f.inputFormatContext.Streams() {
		if v.CodecParameters().CodecType() != avutil.AVMEDIA_TYPE_VIDEO {
			continue
		}

		// 寻找视频流
		f.inputVideoStream = v
	}

	decoder := avcodec.AvcodecFindDecoder(f.inputVideoStream.CodecParameters().CodecId())
	if decoder == nil {
		panic("find codec error")
	}

	encoder := avcodec.AvcodecFindEncoder(avcodec.AV_CODEC_ID_H264)
	if encoder == nil {
		panic("find encoder error")
	}

	f.inputDecodeContext = avcodec.AvcodecAllocContext3(nil)
	if f.inputDecodeContext == nil {
		panic("alloc context 3 error")
	}

	f.outputEncodeContext = avcodec.AvcodecAllocContext3(encoder)
	if f.outputEncodeContext == nil {
		panic("alloc context 3 error")
	}

	err = f.inputDecodeContext.AvcodecParametersToContext(f.inputVideoStream.CodecParameters())
	if f.inputDecodeContext == nil {
		panic(err)
	}

	//err = f.outputEncodeContext.AvcodecParametersToContext(f.inputVideoStream.CodecParameters())
	//if err != nil {
	//	panic(err)
	//}

	f.outputEncodeContext.SetTimebase(1, 25)
	f.outputEncodeContext.SetPixFmt(avutil.AV_PIX_FMT_YUV420P)
	f.outputEncodeContext.SetWidth(100)
	f.outputEncodeContext.SetHeight(100)
	f.outputEncodeContext.SetMaxBFrames(1)

	err = f.inputDecodeContext.AvcodecOpen2(decoder, nil)
	if f.inputDecodeContext == nil {
		panic(err)
	}

	err = f.outputEncodeContext.AvcodecOpen2(encoder, nil)
	if err != nil {
		panic(err)
	}
}

func (f *Filter) openOutput(outputFilename string) {
	f.outputFormatContext = avformat.AvformatAllocContext()
	err := avformat.AvformatAllocOutputContext2(&f.outputFormatContext, nil, "", outputFilename)
	if err != nil {
		panic(err)
	}

	err = avformat.AvIOOpen(f.outputFormatContext.Pb2(), outputFilename, avformat.AVIO_FLAG_WRITE)
	if err != nil {
		panic(err)
	}

	f.outputVideoStream = f.outputFormatContext.AvformatNewStream(nil)
	err = avcodec.AvCodecParametersCopy(f.outputVideoStream.CodecParameters(), f.inputVideoStream.CodecParameters())
	if err != nil {
		panic(err)
	}
}

func (f *Filter) filter() {
	var err error
	err = f.outputFormatContext.AvformatWriteHeader(nil)
	if err != nil {
		panic(err)
	}

	frame := avutil.AvFrameAlloc()
	filterFrame := avutil.AvFrameAlloc()

	packet := avcodec.AvPacketAlloc()
	for {
		err = f.inputFormatContext.AvReadFrame(packet)
		if err == avutil.ErrEOF {
			break
		} else if err != nil {
			panic(err)
		}

		if f.inputFormatContext.Streams()[packet.StreamIndex()].CodecParameters().CodecType() != avutil.AVMEDIA_TYPE_VIDEO {
			continue
		}

		err = f.inputDecodeContext.AvcodecSendPacket(packet)
		if err != nil {
			panic(err)
		}

		for {
			err = f.inputDecodeContext.AvcodecReceiveFrame(frame)
			if err == avutil.ErrEOF || err == avutil.ErrEAGAIN {
				break
			} else if err != nil {
				panic(err)
			}

			err = f.bufferSrcContext.AvBuffersrcAddFrameFlags(frame, avfilter.AV_BUFFERSRC_FLAG_KEEP_REF)
			if err != nil {
				panic(err)
			}

			encodePacket := avcodec.AvPacketAlloc()
			for {
				err = f.bufferSinkContext.AvBuffersinkGetFrame(filterFrame)
				if err == avutil.ErrEOF || err == avutil.ErrEAGAIN {
					break
				} else if err != nil {
					panic(err)
				}

				// 新的帧写入输出文件
				err = f.outputEncodeContext.AvcodecSendFrame(filterFrame)
				if err != nil {
					panic(err)
				}

				for {
					err = f.outputEncodeContext.AvcodecReceivePacket(encodePacket)
					if err == avutil.ErrEOF || err == avutil.ErrEAGAIN {
						break
					} else if err != nil {
						panic(err)
					}

					encodePacket.SetStreamIndex(f.outputVideoStream.Index())
					//encodePacket.AvPacketRescaleTs(f.outputEncodeContext.GetTimeBase(), f.outputVideoStream.TimeBase())
					//encodePacket.SetPts(int64(avutil.AvRescaleQRnd(int(encodePacket.Pts()), f.inputVideoStream.TimeBase(), f.outputVideoStream.TimeBase(), avutil.AV_ROUND_INF | avutil.AV_ROUND_PASS_MINMAX)))
					//encodePacket.SetDts(int64(avutil.AvRescaleQRnd(int(encodePacket.Dts()), f.inputVideoStream.TimeBase(), f.outputVideoStream.TimeBase(), avutil.AV_ROUND_INF | avutil.AV_ROUND_PASS_MINMAX)))
					//encodePacket.SetDuration(int64(avutil.AvRescaleQ(int(encodePacket.Duration()), f.inputVideoStream.TimeBase(), f.outputVideoStream.TimeBase())))

					fmt.Println(encodePacket.Pts(), encodePacket.Dts())

					err = f.outputFormatContext.AvInterleavedWriteFrame(encodePacket)
					if err != nil {
						panic(err)
					}
					encodePacket.AvPacketUnref()
				}

				filterFrame.AvFrameUnref()
			}

			frame.AvFrameUnref()
		}

		packet.AvPacketUnref()
	}
	err = f.outputFormatContext.AvWriteTrailer()
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage:%s <input_filename> <output_filename>\n", os.Args[0])
		return
	}

	f := newFilter(os.Args[1], os.Args[2])
	f.filter()
}
