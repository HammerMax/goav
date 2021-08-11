// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"

	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avutil"
)

func (avs *Stream) CodecParameters() *avcodec.AvCodecParameters {
	return (*avcodec.AvCodecParameters)(unsafe.Pointer(avs.codecpar))
}

func (avs *Stream) Codec() *avcodec.Context {
	return (*avcodec.Context)(unsafe.Pointer(avs.codec))
}

func (avs *Stream) Metadata() *avutil.Dictionary {
	return (*avutil.Dictionary)(unsafe.Pointer(avs.metadata))
}

func (avs *Stream) IndexEntries() *AvIndexEntry {
	return (*AvIndexEntry)(unsafe.Pointer(avs.index_entries))
}

func (avs *Stream) AttachedPic() avcodec.Packet {
	return *fromCPacket(&avs.attached_pic)
}

func (avs *Stream) SideData() *AvPacketSideData {
	return (*AvPacketSideData)(unsafe.Pointer(avs.side_data))
}

//func (avs *Stream) AvgFrameRate() avutil.Rational {
//	return avutil.NewRational(avs.avg_frame_rate)
//}
//
//func (avs *Stream) RFrameRate() avutil.Rational {
//	return avutil.NewRational(avs.r_frame_rate)
//}
//
//func (avs *Stream) SampleAspectRatio() avutil.Rational {
//	return avutil.NewRational(avs.sample_aspect_ratio)
//}
//
//func (avs *Stream) TimeBase() avutil.Rational {
//	return avutil.NewRational(avs.time_base)
//}

// func (avs *Stream) RecommendedEncoderConfiguration() string {
// 	return C.GoString(avs.recommended_encoder_configuration)
// }

func (avs *Stream) Discard() AvDiscard {
	return AvDiscard(avs.discard)
}

func (avs *Stream) NeedParsing() AvStreamParseType {
	return AvStreamParseType(avs.need_parsing)
}

func (avs *Stream) CodecInfoNbFrames() int {
	return int(avs.codec_info_nb_frames)
}

func (avs *Stream) Disposition() int {
	return int(avs.disposition)
}

func (avs *Stream) EventFlags() int {
	return int(avs.event_flags)
}

func (avs *Stream) Id() int {
	return int(avs.id)
}

func (avs *Stream) Index() int {
	return int(avs.index)
}

func (avs *Stream) LastIpDuration() int {
	return int(avs.last_IP_duration)
}

func (avs *Stream) NbIndexEntries() int {
	return int(avs.nb_index_entries)
}

func (avs *Stream) NbSideData() int {
	return int(avs.nb_side_data)
}

func (avs *Stream) ProbePackets() int {
	return int(avs.probe_packets)
}

func (avs *Stream) StreamIdentifier() int {
	return int(avs.stream_identifier)
}

func (avs *Stream) CurDts() int64 {
	return int64(avs.cur_dts)
}

func (avs *Stream) Duration() int64 {
	return int64(avs.duration)
}

func (avs *Stream) FirstDts() int64 {
	return int64(avs.first_dts)
}

func (avs *Stream) LastIpPts() int64 {
	return int64(avs.last_IP_pts)
}

func (avs *Stream) NbFrames() int64 {
	return int64(avs.nb_frames)
}

func (avs *Stream) StartTime() int64 {
	return int64(avs.start_time)
}

func (avs *Stream) Parser() *CodecParserContext {
	return (*CodecParserContext)(unsafe.Pointer(avs.parser))
}

func (avs *Stream) IndexEntriesAllocatedSize() uint {
	return uint(avs.index_entries_allocated_size)
}

func (avs *Stream) Free() {
	C.av_freep(unsafe.Pointer(avs))
}
