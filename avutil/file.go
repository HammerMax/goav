package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/file.h>
import "C"
import (
	"reflect"
	"unsafe"
)

func AvFileMap(fileName string, buf *[]byte, size *int, logOffset int, logCtx interface{}) int32 {
	bufHeader := (*reflect.SliceHeader)(unsafe.Pointer(buf))
	bufPtr := unsafe.Pointer(bufHeader.Data)
	result := int32(C.av_file_map(C.CString(fileName), (**C.uint8_t)(unsafe.Pointer(&bufPtr)), (*C.size_t)(unsafe.Pointer(size)), C.int(logOffset), unsafe.Pointer(&logCtx)))
	bufHeader.Len = *size
	bufHeader.Cap = *size
	return result
}
