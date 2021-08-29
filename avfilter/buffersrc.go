package avfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/buffersrc.h>
*/
import "C"
import (
	"github.com/HammerMax/goav/avutil"
	"unsafe"
)

const (
	AV_BUFFERSRC_FLAG_NO_CHECK_FORMAT = C.AV_BUFFERSRC_FLAG_NO_CHECK_FORMAT
	AV_BUFFERSRC_FLAG_PUSH = C.AV_BUFFERSRC_FLAG_PUSH
	AV_BUFFERSRC_FLAG_KEEP_REF = C.AV_BUFFERSRC_FLAG_KEEP_REF
)

func (c *Context) AvBuffersrcAddFrameFlags(frame *avutil.Frame, flags int) error {
	return avutil.ErrorFromCode(int(C.av_buffersrc_add_frame_flags((*C.struct_AVFilterContext)(c), (*C.struct_AVFrame)(unsafe.Pointer(frame)), C.int(flags))))
}
