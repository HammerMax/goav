package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/codec_par.h>
import "C"
import (
	"github.com/HammerMax/goav/avutil"
	"unsafe"
)

func AvCodecParametersCopy(dst, src *AvCodecParameters) error {
	return avutil.ErrorFromCode(int(C.avcodec_parameters_copy((*C.struct_AVCodecParameters)(unsafe.Pointer(dst)), (*C.struct_AVCodecParameters)(unsafe.Pointer(src)))))
}

func (cp *AvCodecParameters) SetCodecTag(tag uint32) {
	cp.codec_tag = C.uint32_t(tag)
}
