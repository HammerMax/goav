// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"fmt"
	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avutil"
)

func (s *Stream) AvStreamGetParser() *CodecParserContext {
	return (*CodecParserContext)(C.av_stream_get_parser((*C.struct_AVStream)(s)))
}

//int64_t av_stream_get_end_pts (const Stream *st)
//Returns the pts of the last muxed packet + its duration.
func (s *Stream) AvStreamGetEndPts() int64 {
	return int64(C.av_stream_get_end_pts((*C.struct_AVStream)(s)))
}

func (s *Stream) SetID(id int) {
	s.id = C.int(id)
}

func (s *Stream) String() string {
	codecPar := s.CodecParameters()
	mediaType := codecPar.CodecType()
	codecID := codecPar.CodecId()
	str := fmt.Sprintf("Stream index:%d id:%d media_type:%s codec_id:%s ",
		s.Index(), s.Id(), avutil.AvGetMediaTypeString(mediaType), avcodec.AvcodecGetName(codecID))
	switch mediaType {
	case avutil.AVMEDIA_TYPE_VIDEO:
		str += fmt.Sprintf("format:%s ", avutil.AvGetPixFmtName(avutil.PixelFormat(codecPar.Format())))
		str += fmt.Sprintf("%dx%d ", codecPar.Width(), codecPar.Height())
	case avutil.AVMEDIA_TYPE_AUDIO:
		str += fmt.Sprintf("format:%s", avutil.AvGetSampleFmtName(avutil.SampleFormat(codecPar.Format())))
	}
	return str
}