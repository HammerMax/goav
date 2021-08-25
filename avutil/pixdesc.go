package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/pixdesc.h>
import "C"

func AvGetPixFmtName(pixFmt PixelFormat) string {
	return C.GoString(C.av_get_pix_fmt_name(C.enum_AVPixelFormat(pixFmt)))
}
