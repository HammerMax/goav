// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"github.com/HammerMax/goav/avcodec"
	"reflect"
	"unsafe"

	"github.com/HammerMax/goav/avutil"
)

func (c *Context) Chapters() **AvChapter {
	return (**AvChapter)(unsafe.Pointer(c.chapters))
}

func (c *Context) AudioCodec() *avcodec.Codec {
	return (*avcodec.Codec)(unsafe.Pointer(c.audio_codec))
}

func (c *Context) SubtitleCodec() *avcodec.Codec {
	return (*avcodec.Codec)(unsafe.Pointer(c.subtitle_codec))
}

func (c *Context) VideoCodec() *avcodec.Codec {
	return (*avcodec.Codec)(unsafe.Pointer(c.video_codec))
}

func (c *Context) Metadata() *avutil.Dictionary {
	return (*avutil.Dictionary)(unsafe.Pointer(c.metadata))
}

func (c *Context) Internal() *AvFormatInternal {
	return (*AvFormatInternal)(unsafe.Pointer(c.internal))
}

func (c *Context) Pb() *AvIOContext {
	return (*AvIOContext)(unsafe.Pointer(c.pb))
}

func (c *Context) InterruptCallback() AvIOInterruptCB {
	return AvIOInterruptCB(c.interrupt_callback)
}

func (c *Context) Programs() []*AvProgram {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(c.programs)),
		Len:  int(c.NbPrograms()),
		Cap:  int(c.NbPrograms()),
	}

	return *((*[]*AvProgram)(unsafe.Pointer(&header)))
}

func (c *Context) Streams() []*Stream {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(c.streams)),
		Len:  int(c.NbStreams()),
		Cap:  int(c.NbStreams()),
	}

	return *((*[]*Stream)(unsafe.Pointer(&header)))
}

func (c *Context) Filename() string {
	return C.GoString((*C.char)(unsafe.Pointer(&c.filename[0])))
}

// func (ctxt *Context) CodecWhitelist() string {
// 	return C.GoString(ctxt.codec_whitelist)
// }

// func (ctxt *Context) FormatWhitelist() string {
// 	return C.GoString(ctxt.format_whitelist)
// }

func (c *Context) AudioCodecId() avcodec.CodecId {
	return avcodec.CodecId(c.audio_codec_id)
}

func (c *Context) SubtitleCodecId() avcodec.CodecId {
	return avcodec.CodecId(c.subtitle_codec_id)
}

func (c *Context) VideoCodecId() avcodec.CodecId {
	return avcodec.CodecId(c.video_codec_id)
}

func (c *Context) DurationEstimationMethod() AvDurationEstimationMethod {
	return AvDurationEstimationMethod(c.duration_estimation_method)
}

func (c *Context) AudioPreload() int {
	return int(c.audio_preload)
}

func (c *Context) AvioFlags() int {
	return int(c.avio_flags)
}

func (c *Context) AvoidNegativeTs() int {
	return int(c.avoid_negative_ts)
}

func (c *Context) BitRate() int {
	return int(c.bit_rate)
}

func (c *Context) CtxFlags() int {
	return int(c.ctx_flags)
}

func (c *Context) Debug() int {
	return int(c.debug)
}

func (c *Context) ErrorRecognition() int {
	return int(c.error_recognition)
}

func (c *Context) EventFlags() int {
	return int(c.event_flags)
}

func (c *Context) Flags() int {
	return int(c.flags)
}

func (c *Context) FlushPackets() int {
	return int(c.flush_packets)
}

func (c *Context) FormatProbesize() int {
	return int(c.format_probesize)
}

func (c *Context) FpsProbeSize() int {
	return int(c.fps_probe_size)
}

func (c *Context) IoRepositioned() int {
	return int(c.io_repositioned)
}

func (c *Context) Keylen() int {
	return int(c.keylen)
}

func (c *Context) MaxChunkDuration() int {
	return int(c.max_chunk_duration)
}

func (c *Context) MaxChunkSize() int {
	return int(c.max_chunk_size)
}

func (c *Context) MaxDelay() int {
	return int(c.max_delay)
}

func (c *Context) MaxTsProbe() int {
	return int(c.max_ts_probe)
}

func (c *Context) MetadataHeaderPadding() int {
	return int(c.metadata_header_padding)
}

func (c *Context) ProbeScore() int {
	return int(c.probe_score)
}

func (c *Context) Seek2any() int {
	return int(c.seek2any)
}

func (c *Context) StrictStdCompliance() int {
	return int(c.strict_std_compliance)
}

func (c *Context) TsId() int {
	return int(c.ts_id)
}

func (c *Context) UseWallclockAsTimestamps() int {
	return int(c.use_wallclock_as_timestamps)
}

func (c *Context) Duration() int64 {
	return int64(c.duration)
}

func (c *Context) MaxAnalyzeDuration2() int64 {
	return int64(c.max_analyze_duration)
}

func (c *Context) MaxInterleaveDelta() int64 {
	return int64(c.max_interleave_delta)
}

func (c *Context) OutputTsOffset() int64 {
	return int64(c.output_ts_offset)
}

func (c *Context) Probesize2() int64 {
	return int64(c.probesize)
}

func (c *Context) SkipInitialBytes() int64 {
	return int64(c.skip_initial_bytes)
}

func (c *Context) StartTime() int64 {
	return int64(c.start_time)
}

func (c *Context) StartTimeRealtime() int64 {
	return int64(c.start_time_realtime)
}

func (c *Context) Iformat() *InputFormat {
	return (*InputFormat)(unsafe.Pointer(c.iformat))
}

func (c *Context) Oformat() *OutputFormat {
	return (*OutputFormat)(unsafe.Pointer(c.oformat))
}

// func (ctxt *Context) DumpSeparator() uint8 {
// 	return uint8(ctxt.dump_separator)
// }

func (c *Context) CorrectTsOverflow() int {
	return int(c.correct_ts_overflow)
}

func (c *Context) MaxIndexSize() uint {
	return uint(c.max_index_size)
}

func (c *Context) MaxPictureBuffer() uint {
	return uint(c.max_picture_buffer)
}

func (c *Context) NbChapters() uint {
	return uint(c.nb_chapters)
}

func (c *Context) NbPrograms() uint {
	return uint(c.nb_programs)
}

func (c *Context) NbStreams() uint {
	return uint(c.nb_streams)
}

func (c *Context) PacketSize() uint {
	return uint(c.packet_size)
}

func (c *Context) Probesize() uint {
	return uint(c.probesize)
}

func (c *Context) SetPb(pb *AvIOContext) {
	c.pb = (*C.struct_AVIOContext)(unsafe.Pointer(pb))
}

func (c *Context) Pb2() **AvIOContext {
	return (**AvIOContext)(unsafe.Pointer(&c.pb))
}
