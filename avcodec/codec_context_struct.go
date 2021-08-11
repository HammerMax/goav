// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avformat provides some generic global options, which can be set on all the muxers and demuxers.
//In addition each muxer or demuxer may support so-called private options, which are specific for that component.
//Supported formats (muxers and demuxers) provided by the libavformat library
package avcodec

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
import "C"
import (
	"github.com/giorgisio/goav/avutil"
	"reflect"
	"unsafe"
)

func (c *Context) Type() avutil.MediaType {
	return avutil.MediaType(c.codec_type)
}

func (c *Context) SetBitRate(br int64) {
	c.bit_rate = C.int64_t(br)
}

func (c *Context) GetCodecId() CodecId {
	return CodecId(c.codec_id)
}

func (c *Context) SetCodecId(codecId CodecId) {
	c.codec_id = C.enum_AVCodecID(codecId)
}

func (c *Context) GetCodecType() avutil.MediaType {
	return avutil.MediaType(c.codec_type)
}

func (c *Context) SetCodecType(ctype avutil.MediaType) {
	c.codec_type = C.enum_AVMediaType(ctype)
}

func (c *Context) GetTimeBase() avutil.Rational {
	return avutil.NewRational(int(c.time_base.num), int(c.time_base.den))
}

func (c *Context) SetTimeBase(timeBase avutil.Rational) {
	c.time_base.num = C.int(timeBase.Num())
	c.time_base.den = C.int(timeBase.Den())
}

func (c *Context) GetWidth() int {
	return int(c.width)
}

func (c *Context) SetWidth(w int) {
	c.width = C.int(w)
}

func (c *Context) GetHeight() int {
	return int(c.height)
}

func (c *Context) SetHeight(h int) {
	c.height = C.int(h)
}

func (c *Context) GetPixelFormat() PixelFormat {
	return PixelFormat(C.int(c.pix_fmt))
}

func (c *Context) SetPixelFormat(fmt PixelFormat) {
	c.pix_fmt = C.enum_AVPixelFormat(C.int(fmt))
}

func (c *Context) GetFlags() int {
	return int(c.flags)
}

func (c *Context) SetFlags(flags int) {
	c.flags = C.int(flags)
}

func (c *Context) GetMeRange() int {
	return int(c.me_range)
}

func (c *Context) SetMeRange(r int) {
	c.me_range = C.int(r)
}

func (c *Context) GetMaxQDiff() int {
	return int(c.max_qdiff)
}

func (c *Context) SetMaxQDiff(v int) {
	c.max_qdiff = C.int(v)
}

func (c *Context) GetQMin() int {
	return int(c.qmin)
}

func (c *Context) SetQMin(v int) {
	c.qmin = C.int(v)
}

func (c *Context) GetQMax() int {
	return int(c.qmax)
}

func (c *Context) SetQMax(v int) {
	c.qmax = C.int(v)
}

func (c *Context) GetQCompress() float32 {
	return float32(c.qcompress)
}

func (c *Context) SetQCompress(v float32) {
	c.qcompress = C.float(v)
}

func (c *Context) GetExtraData() []byte {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(c.extradata)),
		Len:  int(c.extradata_size),
		Cap:  int(c.extradata_size),
	}

	return *((*[]byte)(unsafe.Pointer(&header)))
}

func (c *Context) SetExtraData(data []byte) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&data))

	c.extradata = (*C.uint8_t)(unsafe.Pointer(header.Data))
	c.extradata_size = C.int(header.Len)
}

func (c *Context) Release() {
	C.avcodec_close((*C.struct_AVCodecContext)(unsafe.Pointer(c)))
	C.av_freep(unsafe.Pointer(c))
}
