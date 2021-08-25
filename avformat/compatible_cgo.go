package avformat

//#cgo pkg-config: libavutil libavcodec
//#include <libavutil/avutil.h>
//#include <libavcodec/avcodec.h>
import "C"
import (
	"github.com/HammerMax/goav/avcodec"
	"github.com/HammerMax/goav/avutil"
	"unsafe"
)

func avutilRational(rational C.struct_AVRational) avutil.Rational {
	return avutil.NewRational(int(rational.num), int(rational.den))
}

func toRational(rational avutil.Rational) C.struct_AVRational {
	return C.struct_AVRational{num: C.int(rational.Num()), den: C.int(rational.Den())}
}

func CToCodec(c *C.struct_AVCodec) *avcodec.Codec {
	return (*avcodec.Codec)(unsafe.Pointer(c))
}

func CodecToC(c *avcodec.Codec) *C.struct_AVCodec {
	return (*C.struct_AVCodec)(unsafe.Pointer(c))
}
