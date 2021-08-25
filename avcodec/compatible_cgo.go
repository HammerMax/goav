package avcodec

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"
import "github.com/HammerMax/goav/avutil"

func avutilRational(rational C.struct_AVRational) avutil.Rational {
	return avutil.NewRational(int(rational.num), int(rational.den))
}

func toRational(rational avutil.Rational) C.struct_AVRational {
	return C.struct_AVRational{num: C.int(rational.Num()), den: C.int(rational.Den())}
}
