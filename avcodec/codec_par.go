package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/codec_par.h>
import "C"
import "unsafe"

func AvCodecParametersCopy(dst, src *AvCodecParameters) int {
	return int(C.avcodec_parameters_copy((*C.struct_AVCodecParameters)(unsafe.Pointer(dst)), (*C.struct_AVCodecParameters)(unsafe.Pointer(src))))
}

func (cp *AvCodecParameters) SetCodecTag(tag uint32) {
	cp.codec_tag = C.uint32_t(tag)
}