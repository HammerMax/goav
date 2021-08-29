package avfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/buffersink.h>
*/
import "C"
import (
	"github.com/HammerMax/goav/avutil"
	"unsafe"
)

func (c *Context) AvBuffersinkGetFrame(frame *avutil.Frame) error {
	return avutil.ErrorFromCode(int(C.av_buffersink_get_frame((*C.struct_AVFilterContext)(c), (*C.struct_AVFrame)(unsafe.Pointer(frame)))))
}
