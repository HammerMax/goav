package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/parseutils.h>
//#include <stdlib.h>
import "C"
import "unsafe"

func AvParseVideoSize(width, height *int, str string) error {
	s := C.CString(str)
	defer C.free(unsafe.Pointer(s))

	return ErrorFromCode(int(C.av_parse_video_size((*C.int)(unsafe.Pointer(width)), (*C.int)(unsafe.Pointer(height)), s)))
}