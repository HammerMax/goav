package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/imgutils.h>
import "C"
import (
	"reflect"
	"unsafe"
)

func AvImageAlloc(pointers [][]byte, lineSizes []int32, w, h int, pixFmt PixelFormat, align int) int {
	var ps = make([]unsafe.Pointer, 4)
	for k := range pointers {
		pointerHeader := (*reflect.SliceHeader)(unsafe.Pointer(&pointers[k]))
		ps[k] = unsafe.Pointer(pointerHeader.Data)
	}
	ret := int(C.av_image_alloc((**C.uint8_t)(unsafe.Pointer(&ps[0])), (*C.int)(unsafe.Pointer(&lineSizes[0])), C.int(w), C.int(h), (C.enum_AVPixelFormat)(pixFmt), C.int(align)))
	for k, v := range ps {
		pointerHeader := (*reflect.SliceHeader)(unsafe.Pointer(&pointers[k]))
		pointerHeader.Len = 1000000
		pointerHeader.Cap = 1000000
		pointerHeader.Data = uintptr(v)
	}
	return ret
}
