package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/mathematics.h>
import "C"

func AvRescaleQRnd(a int, bq, cq Rational, rnd AVRounding) int {
	return int(C.av_rescale_q_rnd(C.int64_t(a), (C.struct_AVRational)(bq), (C.struct_AVRational)(cq), (C.enum_AVRounding)(rnd)))
}

func AvRescaleQ(a int, bq, cq Rational) int {
	return int(C.av_rescale_q(C.int64_t(a), (C.struct_AVRational)(bq), (C.struct_AVRational)(cq)))
}
