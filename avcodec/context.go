// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavcodec libavutil
//#include <libavcodec/avcodec.h>
//#include <libavutil/opt.h>
import "C"
import (
	"github.com/giorgisio/goav/avutil"
	"unsafe"
)

func (c *Context) AvCodecGetPktTimebase() avutil.Rational {
	return avutilRational(C.av_codec_get_pkt_timebase((*C.struct_AVCodecContext)(c)))
}

func (c *Context) AvCodecSetPktTimebase(r avutil.Rational) {
	C.av_codec_set_pkt_timebase((*C.struct_AVCodecContext)(c), toRational(r))
}

func (c *Context) AvCodecGetCodecDescriptor() *Descriptor {
	return (*Descriptor)(C.av_codec_get_codec_descriptor((*C.struct_AVCodecContext)(c)))
}

func (c *Context) AvCodecSetCodecDescriptor(d *Descriptor) {
	C.av_codec_set_codec_descriptor((*C.struct_AVCodecContext)(c), (*C.struct_AVCodecDescriptor)(d))
}

func (c *Context) AvCodecGetLowres() int {
	return int(C.av_codec_get_lowres((*C.struct_AVCodecContext)(c)))
}

func (c *Context) AvCodecSetLowres(i int) {
	C.av_codec_set_lowres((*C.struct_AVCodecContext)(c), C.int(i))
}

func (c *Context) AvCodecGetSeekPreroll() int {
	return int(C.av_codec_get_seek_preroll((*C.struct_AVCodecContext)(c)))
}

func (c *Context) AvCodecSetSeekPreroll(i int) {
	C.av_codec_set_seek_preroll((*C.struct_AVCodecContext)(c), C.int(i))
}

func (c *Context) AvCodecGetChromaIntraMatrix() *uint16 {
	return (*uint16)(C.av_codec_get_chroma_intra_matrix((*C.struct_AVCodecContext)(c)))
}

func (c *Context) AvCodecSetChromaIntraMatrix(t *uint16) {
	C.av_codec_set_chroma_intra_matrix((*C.struct_AVCodecContext)(c), (*C.uint16_t)(t))
}

//Free the codec context and everything associated with it and write NULL to the provided pointer.
func (c *Context) AvcodecFreeContext() {
	C.avcodec_free_context((**C.struct_AVCodecContext)(unsafe.Pointer(&c)))
}

//Set the fields of the given Context to default values corresponding to the given codec (defaults may be codec-dependent).
func (c *Context) AvcodecGetContextDefaults3(codec *Codec) int {
	return int(C.avcodec_get_context_defaults3((*C.struct_AVCodecContext)(c), (*C.struct_AVCodec)(codec)))
}

//Copy the settings of the source Context into the destination Context.
func (c *Context) AvcodecCopyContext(ctxt2 *Context) int {
	return int(C.avcodec_copy_context((*C.struct_AVCodecContext)(c), (*C.struct_AVCodecContext)(ctxt2)))
}

//Initialize the Context to use the given Codec
func (c *Context) AvcodecOpen2(codec *Codec, d **Dictionary) int {
	return int(C.avcodec_open2((*C.struct_AVCodecContext)(c), (*C.struct_AVCodec)(codec), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//Close a given Context and free all the data associated with it (but not the Context itself).
func (c *Context) AvcodecClose() int {
	return int(C.avcodec_close((*C.struct_AVCodecContext)(c)))
}

//The default callback for Context.get_buffer2().
func (c *Context) AvcodecDefaultGetBuffer2(f *Frame, l int) int {
	return int(C.avcodec_default_get_buffer2((*C.struct_AVCodecContext)(c), (*C.struct_AVFrame)(f), C.int(l)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you do not use any horizontal padding.
func (c *Context) AvcodecAlignDimensions(w, h *int) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(c), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you also ensure that all line sizes are a multiple of the respective linesize_align[i].
func (c *Context) AvcodecAlignDimensions2(w, h *int, l int) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(c), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(&l)))
}

//Decode the audio frame of size avpkt->size from avpkt->data into frame.
func (c *Context) AvcodecDecodeAudio4(f *Frame, g *int, a *Packet) int {
	return int(C.avcodec_decode_audio4((*C.struct_AVCodecContext)(c), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Decode the video frame of size avpkt->size from avpkt->data into picture.
func (c *Context) AvcodecDecodeVideo2(p *Frame, g *int, a *Packet) int {
	return int(C.avcodec_decode_video2((*C.struct_AVCodecContext)(c), (*C.struct_AVFrame)(p), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Decode a subtitle message.
func (c *Context) AvcodecDecodeSubtitle2(s *AvSubtitle, g *int, a *Packet) int {
	return int(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(c), (*C.struct_AVSubtitle)(s), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Encode a frame of audio.
func (c *Context) AvcodecEncodeAudio2(p *Packet, f *Frame, gp *int) int {
	return int(C.avcodec_encode_audio2((*C.struct_AVCodecContext)(c), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

//Encode a frame of video
func (c *Context) AvcodecEncodeVideo2(p *Packet, f *Frame, gp *int) int {
	return int(C.avcodec_encode_video2((*C.struct_AVCodecContext)(c), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

func (c *Context) AvcodecEncodeSubtitle(b *uint8, bs int, s *AvSubtitle) int {
	return int(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(c), (*C.uint8_t)(b), C.int(bs), (*C.struct_AVSubtitle)(s)))
}

func (c *Context) AvcodecDefaultGetFormat(f *PixelFormat) PixelFormat {
	return (PixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(c), (*C.enum_AVPixelFormat)(f)))
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
func (c *Context) AvParserParse2(ctxtp *ParserContext, p **uint8, ps *int, b *uint8, bs int, pt, dt, po int64) int {
	return int(C.av_parser_parse2((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(c), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), (C.int64_t)(pt), (C.int64_t)(dt), (C.int64_t)(po)))
}

func (c *Context) AvParserChange(ctxtp *ParserContext, pb **uint8, pbs *int, b *uint8, bs, k int) int {
	return int(C.av_parser_change((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(c), (**C.uint8_t)(unsafe.Pointer(pb)), (*C.int)(unsafe.Pointer(pbs)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
}

func AvParserInit(c int) *ParserContext {
	return (*ParserContext)(C.av_parser_init(C.int(c)))
}

func AvParserClose(ctxtp *ParserContext) {
	C.av_parser_close((*C.struct_AVCodecParserContext)(ctxtp))
}

func (p *Parser) AvParserNext() *Parser {
	return (*Parser)(C.av_parser_next((*C.struct_AVCodecParser)(p)))
}

func (p *Parser) AvRegisterCodecParser() {
	C.av_register_codec_parser((*C.struct_AVCodecParser)(p))
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

func (c *Context) SetPixFmt(pixFmt PixelFormat) {
	c.pix_fmt = C.enum_AVPixelFormat(pixFmt)
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

func (c *Context) AvcodecSendPacket(packet *Packet) int {
	return (int)(C.avcodec_send_packet((*C.struct_AVCodecContext)(c), (*C.struct_AVPacket)(packet)))
}

func (c *Context) AvcodecReceivePacket(pkt *Packet) int32 {
	return (int32)(C.avcodec_receive_packet((*C.struct_AVCodecContext)(c), (*C.struct_AVPacket)(pkt)))
}

func (c *Context) AvcodecReceiveFrame(frame *avutil.Frame) int32 {
	return (int32)(C.avcodec_receive_frame((*C.struct_AVCodecContext)(c), (*C.struct_AVFrame)(unsafe.Pointer(frame))))
}

func (c *Context) AvcodecSendFrame(frame *avutil.Frame) int32 {
	return (int32)(C.avcodec_send_frame((*C.struct_AVCodecContext)(c), (*C.struct_AVFrame)(unsafe.Pointer(frame))))
}
