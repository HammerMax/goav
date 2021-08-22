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

func AvCompareTs(tsA int, tbA Rational, tsB int, tbB Rational) int {
	return int(C.av_compare_ts(C.int64_t(tsA), C.struct_AVRational(tbA), C.int64_t(tsB), C.struct_AVRational(tbB)))
}
