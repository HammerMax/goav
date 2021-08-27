// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"time"
	"unsafe"

	"github.com/HammerMax/goav/avcodec"
	"github.com/HammerMax/goav/avutil"
)

const (
	AvseekFlagBackward = 1 ///< seek backward
	AvseekFlagByte     = 2 ///< seeking based on position in bytes
	AvseekFlagAny      = 4 ///< seek to any frame, even non-keyframes
	AvseekFlagFrame    = 8 ///< seeking based on frame number
)

//This function will cause global side data to be injected in the next packet of each stream as well as after any subsequent seek.
func (c *Context) AvFormatInjectGlobalSideData() {
	C.av_format_inject_global_side_data((*C.struct_AVFormatContext)(c))
}

//Returns the method used to set ctx->duration.
func (c *Context) AvFmtCtxGetDurationEstimationMethod() AvDurationEstimationMethod {
	return (AvDurationEstimationMethod)(C.av_fmt_ctx_get_duration_estimation_method((*C.struct_AVFormatContext)(c)))
}

//Free an Context and all its streams.
func (c *Context) AvformatFreeContext() {
	C.avformat_free_context((*C.struct_AVFormatContext)(c))
}

//Add a new stream to a media file.
func (c *Context) AvformatNewStream(codec *avcodec.Codec) *Stream {
	return (*Stream)(C.avformat_new_stream((*C.struct_AVFormatContext)(c), (*C.struct_AVCodec)(unsafe.Pointer(codec))))
}

func (c *Context) AvNewProgram(id int) *AvProgram {
	return (*AvProgram)(C.av_new_program((*C.struct_AVFormatContext)(c), C.int(id)))
}

//Read packets of a media file to get stream information.
func (c *Context) AvformatFindStreamInfo(d **avutil.Dictionary) error {
	return avutil.ErrorFromCode(int(C.avformat_find_stream_info((*C.struct_AVFormatContext)(c), (**C.struct_AVDictionary)(unsafe.Pointer(d)))))
}

//Find the programs which belong to a given stream.
func (c *Context) AvFindProgramFromStream(l *AvProgram, su int) *AvProgram {
	return (*AvProgram)(C.av_find_program_from_stream((*C.struct_AVFormatContext)(c), (*C.struct_AVProgram)(l), C.int(su)))
}

//Find the "best" stream in the file.
func AvFindBestStream(ic *Context, t avutil.MediaType, ws, rs int, c **avcodec.Codec, f int) int {
	return int(C.av_find_best_stream((*C.struct_AVFormatContext)(ic), (C.enum_AVMediaType)(t), C.int(ws), C.int(rs), (**C.struct_AVCodec)(unsafe.Pointer(c)), C.int(f)))
}

//Return the next frame of a stream.
func (c *Context) AvReadFrame(pkt *avcodec.Packet) error {
	return avutil.ErrorFromCode(int(C.av_read_frame((*C.struct_AVFormatContext)(unsafe.Pointer(c)), (*C.struct_AVPacket)(unsafe.Pointer(pkt)))))
}

//Seek to the keyframe at timestamp.
func (c *Context) AvSeekFrame(st int, t int64, f int) error {
	return avutil.ErrorFromCode(int(C.av_seek_frame((*C.struct_AVFormatContext)(c), C.int(st), C.int64_t(t), C.int(f))))
}

// AvSeekFrameTime seeks to a specified time location.
// |timebase| is codec specific and can be obtained by calling AvCodecGetPktTimebase2
func (c *Context) AvSeekFrameTime(st int, at time.Duration, timebase avutil.Rational) int {
	t2 := C.double(C.double(at.Seconds())*C.double(timebase.Den())) / (C.double(timebase.Num()))
	return int(C.av_seek_frame((*C.struct_AVFormatContext)(c), C.int(st), C.int64_t(t2), AvseekFlagBackward))
}

//Seek to timestamp ts.
func (c *Context) AvformatSeekFile(si int, mit, ts, mat int64, f int) int {
	return int(C.avformat_seek_file((*C.struct_AVFormatContext)(c), C.int(si), C.int64_t(mit), C.int64_t(ts), C.int64_t(mat), C.int(f)))
}

//Start playing a network-based stream (e.g.
func (c *Context) AvReadPlay() int {
	return int(C.av_read_play((*C.struct_AVFormatContext)(c)))
}

//Pause a network-based stream (e.g.
func (c *Context) AvReadPause() int {
	return int(C.av_read_pause((*C.struct_AVFormatContext)(c)))
}

//Close an opened input Context.
func (c *Context) AvformatCloseInput() {
	C.avformat_close_input((**C.struct_AVFormatContext)(unsafe.Pointer(&c)))
}

//Allocate the stream private data and write the stream header to an output media file.
func (c *Context) AvformatWriteHeader(o **avutil.Dictionary) error {
	return avutil.ErrorFromCode(int(C.avformat_write_header((*C.struct_AVFormatContext)(c), (**C.struct_AVDictionary)(unsafe.Pointer(o)))))
}

//Write a packet to an output media file.
func (c *Context) AvWriteFrame(pkt *avcodec.Packet) int {
	return int(C.av_write_frame((*C.struct_AVFormatContext)(c), toCPacket(pkt)))
}

//Write a packet to an output media file ensuring correct interleaving.
func (c *Context) AvInterleavedWriteFrame(pkt *avcodec.Packet) error {
	return avutil.ErrorFromCode(int(C.av_interleaved_write_frame((*C.struct_AVFormatContext)(c), toCPacket(pkt))))
}

//Write a uncoded frame to an output media file.
func (c *Context) AvWriteUncodedFrame(si int, f *Frame) int {
	return int(C.av_write_uncoded_frame((*C.struct_AVFormatContext)(c), C.int(si), (*C.struct_AVFrame)(f)))
}

//Write a uncoded frame to an output media file.
func (c *Context) AvInterleavedWriteUncodedFrame(si int, f *Frame) int {
	return int(C.av_interleaved_write_uncoded_frame((*C.struct_AVFormatContext)(c), C.int(si), (*C.struct_AVFrame)(f)))
}

//Test whether a muxer supports uncoded frame.
func (c *Context) AvWriteUncodedFrameQuery(si int) int {
	return int(C.av_write_uncoded_frame_query((*C.struct_AVFormatContext)(c), C.int(si)))
}

//Write the stream trailer to an output media file and free the file private data.
func (c *Context) AvWriteTrailer() error {
	return avutil.ErrorFromCode(int(C.av_write_trailer((*C.struct_AVFormatContext)(c))))
}

//Get timing information for the data currently output.
func (c *Context) AvGetOutputTimestamp(st int, dts, wall *int) int {
	return int(C.av_get_output_timestamp((*C.struct_AVFormatContext)(c), C.int(st), (*C.int64_t)(unsafe.Pointer(&dts)), (*C.int64_t)(unsafe.Pointer(&wall))))
}

func (c *Context) AvFindDefaultStreamIndex() int {
	return int(C.av_find_default_stream_index((*C.struct_AVFormatContext)(c)))
}

//Print detailed information about the input or output format, such as duration, bitrate, streams, container, programs, metadata, side data, codec and time base.
func (c *Context) AvDumpFormat(i int, url string, io int) {
	Curl := C.CString(url)
	defer C.free(unsafe.Pointer(Curl))

	C.av_dump_format((*C.struct_AVFormatContext)(unsafe.Pointer(c)), C.int(i), Curl, C.int(io))
}

////Guess the sample aspect ratio of a frame, based on both the stream and the frame aspect ratio.
//func (s *Context) AvGuessSampleAspectRatio(st *Stream, fr *Frame) avutil.Rational {
//	return avutil.NewRational(C.av_guess_sample_aspect_ratio((*C.struct_AVFormatContext)(s), (*C.struct_AVStream)(st),
//		(*C.struct_AVFrame)(fr)))
//}

////Guess the frame rate, based on both the container and codec information.
//func (s *Context) AvGuessFrameRate(st *Stream, fr *Frame) avutil.Rational {
//	return avutil.NewRational(C.av_guess_frame_rate((*C.struct_AVFormatContext)(s), (*C.struct_AVStream)(st),
//		(*C.struct_AVFrame)(fr)))
//}

//Check if the stream st contained in s is matched by the stream specifier spec.
func (c *Context) AvformatMatchStreamSpecifier(st *Stream, spec string) int {
	Cspec := C.CString(spec)
	defer C.free(unsafe.Pointer(Cspec))

	return int(C.avformat_match_stream_specifier((*C.struct_AVFormatContext)(c), (*C.struct_AVStream)(st), Cspec))
}

func (c *Context) AvformatQueueAttachedPictures() int {
	return int(C.avformat_queue_attached_pictures((*C.struct_AVFormatContext)(c)))
}

//func (s *Context) AvformatNewStream2(c *avcodec.Codec) *Stream {
//	stream := (*Stream)(C.avformat_new_stream((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c)))
//	stream.codec.pix_fmt = int32(avcodec.AV_PIX_FMT_YUV)
//	stream.codec.width = 640
//	stream.codec.height = 480
//	stream.time_base.num = 1
//	stream.time_base.num = 25
//	return stream
//}

// //av_format_control_message av_format_get_control_message_cb (const Context *s)
// func (s *Context) AvFormatControlMessage() C.av_format_get_control_message_cb {
// 	return C.av_format_get_control_message_cb((*C.struct_AVFormatContext)(s))
// }

// //void av_format_set_control_message_cb (Context *s, av_format_control_message callback)
// func (s *Context) AvFormatSetControlMessageCb(c AvFormatControlMessage) C.av_format_get_control_message_cb {
// 	C.av_format_set_control_message_cb((*C.struct_AVFormatContext)(s), (C.struct_AVFormatControlMessage)(c))
// }

// //AvCodec * av_format_get_data_codec (const Context *s)
// func (s *Context)AvFormatGetDataCodec() *AvCodec {
// 	return (*AvCodec)(C.av_format_get_data_codec((*C.struct_AVFormatContext)(s)))
// }

// //void av_format_set_data_codec (Context *s, AvCodec *c)
// func (s *Context)AvFormatSetDataCodec( c *AvCodec) {
// 	C.av_format_set_data_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
// }
