package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/opt.h>
import "C"
import (
	"reflect"
	"unsafe"
)

func AvOptSetInt(i interface{}, name string, val int, searchFlags int) error {
	return ErrorFromCode(int(C.av_opt_set_int(unsafe.Pointer(reflect.ValueOf(i).Pointer()), C.CString(name), C.int64_t(val), C.int(searchFlags))))
}

func AvOptSetSampleFmt(i interface{}, name string, sfmt SampleFormat, searchFlags int) error {
	return ErrorFromCode(int(C.av_opt_set_sample_fmt(unsafe.Pointer(reflect.ValueOf(i).Pointer()), C.CString(name), int32(sfmt), C.int(searchFlags))))
}