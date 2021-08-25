// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavcodec libavutil
//#include <libavcodec/avcodec.h>
//#include <libavutil/opt.h>
import "C"
import (
	"github.com/HammerMax/goav/avutil"
	"reflect"
	"unsafe"
)


//Free the codec context and everything associated with it and write NULL to the provided pointer.
func (c *Context) AvcodecFreeContext() {
	C.avcodec_free_context((**C.struct_AVCodecContext)(unsafe.Pointer(&c)))
}

//Set the fields of the given Context to default values corresponding to the given codec (defaults may be codec-dependent).
func (c *Context) AvcodecGetContextDefaults3(codec *Codec) int {
	return int(C.avcodec_get_context_defaults3((*C.struct_AVCodecContext)(c), (*C.struct_AVCodec)(codec)))
}

//Initialize the Context to use the given Codec
func (c *Context) AvcodecOpen2(codec *Codec, d **Dictionary) error {
	return avutil.ErrorFromCode(int(C.avcodec_open2((*C.struct_AVCodecContext)(c), (*C.struct_AVCodec)(codec), (**C.struct_AVDictionary)(unsafe.Pointer(d)))))
}

//Close a given Context and free all the data associated with it (but not the Context itself).
func (c *Context) AvcodecClose() int {
	return int(C.avcodec_close((*C.struct_AVCodecContext)(c)))
}

//The default callback for Context.get_buffer2().
func (c *Context) AvcodecDefaultGetBuffer2(f *avutil.Frame, l int) int {
	return int(C.avcodec_default_get_buffer2((*C.struct_AVCodecContext)(c), (*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(l)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you do not use any horizontal padding.
func (c *Context) AvcodecAlignDimensions(w, h *int) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(c), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you also ensure that all line sizes are a multiple of the respective linesize_align[i].
func (c *Context) AvcodecAlignDimensions2(w, h *int, l int) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(c), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(&l)))
}

//Decode a subtitle message.
func (c *Context) AvcodecDecodeSubtitle2(s *AvSubtitle, g *int, a *Packet) int {
	return int(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(c), (*C.struct_AVSubtitle)(s), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

func (c *Context) AvcodecEncodeSubtitle(b *uint8, bs int, s *AvSubtitle) int {
	return int(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(c), (*C.uint8_t)(b), C.int(bs), (*C.struct_AVSubtitle)(s)))
}

func (c *Context) AvcodecDefaultGetFormat(f *PixelFormat) PixelFormat {
	return (PixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(c), (*C.enum_AVPixelFormat)(f)))
}

func (c *Context) AvcodecParametersToContext(parameter *AvCodecParameters) error {
	return avutil.ErrorFromCode(int(C.avcodec_parameters_to_context((*C.struct_AVCodecContext)(c), (*C.struct_AVCodecParameters)(parameter))))
}

//Reset the internal decoder state / flush internal buffers.
func (c *Context) AvcodecFlushBuffers() {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(c))
}

//Return audio frame duration.
func (c *Context) AvGetAudioFrameDuration(f int) int {
	return int(C.av_get_audio_frame_duration((*C.struct_AVCodecContext)(c), C.int(f)))
}

func (c *Context) AvcodecIsOpen() int {
	return int(C.avcodec_is_open((*C.struct_AVCodecContext)(c)))
}

//Parse a packet.
func (p *ParserContext) AvParserParse2(c *Context, packetData **uint8, packetSize *int, data []byte, dataSize int, pt, dt, po int64) int {
	sliceHeader := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	sliceHeaderPoint := sliceHeader.Data
	return int(C.av_parser_parse2((*C.struct_AVCodecParserContext)(p), (*C.struct_AVCodecContext)(unsafe.Pointer(c)), (**C.uint8_t)(unsafe.Pointer(packetData)), (*C.int)(unsafe.Pointer(packetSize)), (*C.uint8_t)(unsafe.Pointer(sliceHeaderPoint)), C.int(dataSize), (C.int64_t)(pt), (C.int64_t)(dt), (C.int64_t)(po)))
}

func AvParserInit(c CodecId) *ParserContext {
	return (*ParserContext)(C.av_parser_init(C.int(c)))
}

func (c *ParserContext) AvParserClose() {
	C.av_parser_close((*C.struct_AVCodecParserContext)(c))
}

func (c *Context) SetTimebase(num int, den int) {
	c.time_base.num = C.int(num)
	c.time_base.den = C.int(den)
}

func (c *Context) SetFramerate(num int, den int) {
	c.framerate.num = C.int(num)
	c.framerate.den = C.int(den)
}

func (c *Context) SetGopSize(size int) {
	c.gop_size = C.int(size)
}

func (c *Context) SetMaxBFrames(frames int) {
	c.max_b_frames = C.int(frames)
}

func (c *Context) SetPixFmt(pixFmt avutil.PixelFormat) {
	c.pix_fmt = C.enum_AVPixelFormat(pixFmt)
}

func (c *Context) SetSampleFmt(sampleFmt avutil.SampleFormat) {
	c.sample_fmt = C.enum_AVSampleFormat(sampleFmt)
}

func (c *Context) SetEncodeParams2(width int, height int, pxlFmt PixelFormat, hasBframes bool, gopSize int) {
	c.width = C.int(width)
	c.height = C.int(height)
	// ctxt.bit_rate = 1000000
	c.gop_size = C.int(gopSize)
	// ctxt.max_b_frames = 2
	if hasBframes {
		c.has_b_frames = 1
	} else {
		c.has_b_frames = 0
	}
	// ctxt.extradata = nil
	// ctxt.extradata_size = 0
	// ctxt.channels = 0
	c.pix_fmt = int32(pxlFmt)
	// C.av_opt_set(ctxt.priv_data, "preset", "ultrafast", 0)
}

func (c *Context) AvOptSet(name, val string, searchFlags int) int {
	Cname := C.CString(name)
	defer C.free(unsafe.Pointer(Cname))

	Cval := C.CString(val)
	defer C.free(unsafe.Pointer(Cval))

	return int(C.av_opt_set(c.priv_data, Cname, Cval, C.int(searchFlags)))
}

func (c *Context) SetEncodeParams(width int, height int, pxlFmt PixelFormat) {
	c.SetEncodeParams2(width, height, pxlFmt, false /*no b frames*/, 10)
}

func (c *Context) AvcodecSendPacket(packet *Packet) error {
	return avutil.ErrorFromCode((int)(C.avcodec_send_packet((*C.struct_AVCodecContext)(c), (*C.struct_AVPacket)(packet))))
}

func (c *Context) AvcodecReceivePacket(pkt *Packet) error {
	return avutil.ErrorFromCode((int)(C.avcodec_receive_packet((*C.struct_AVCodecContext)(c), (*C.struct_AVPacket)(pkt))))
}

func (c *Context) AvcodecReceiveFrame(frame *avutil.Frame) error {
	return avutil.ErrorFromCode((int)(C.avcodec_receive_frame((*C.struct_AVCodecContext)(c), (*C.struct_AVFrame)(unsafe.Pointer(frame)))))
}

func (c *Context) AvcodecSendFrame(frame *avutil.Frame) error {
	return avutil.ErrorFromCode((int)(C.avcodec_send_frame((*C.struct_AVCodecContext)(c), (*C.struct_AVFrame)(unsafe.Pointer(frame)))))
}
