// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package avutil is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
//#include <libavutil/samplefmt.h>
import "C"
import (
	"fmt"
	"unsafe"
)

const (
	AV_NOPTS_VALUE = C.AV_NOPTS_VALUE
)

type (
	Options       C.struct_AVOptions
	AvTree        C.struct_AVTree
	Rational      C.struct_AVRational
	MediaType     C.enum_AVMediaType
	AvPictureType C.enum_AVPictureType
	PixelFormat   C.enum_AVPixelFormat
	AVRounding    C.enum_AVRounding
	SampleFormat  C.enum_AVSampleFormat
	File          C.FILE
)

const (
	AV_ROUND_ZERO     = AVRounding(0)
	AV_ROUND_INF      = AVRounding(1)
	AV_ROUND_DOWN     = AVRounding(2)
	AV_ROUND_UP       = AVRounding(3)
	AV_ROUND_NEAR_INF = AVRounding(5)
	AV_ROUND_PASS_MINMAX = AVRounding(8192)
)

const (
	AV_TIME_BASE = C.AV_TIME_BASE
)

func (r Rational) Num() int {
	return int(r.num)
}

func (r Rational) Den() int {
	return int(r.den)
}

func (r Rational) String() string {
	return fmt.Sprintln("%d/%d", int(r.num), int(r.den))
}

func (r *Rational) Assign(o Rational) {
	r.Set(o.Num(), o.Den())
}

func (r *Rational) Set(num, den int) {
	r.num, r.den = C.int(num), C.int(den)
}

func NewRational(num, den int) Rational {
	return Rational{
		num: C.int(num),
		den: C.int(den),
	}
}


//Return the LIBAvUTIL_VERSION_INT constant.
func AvutilVersion() uint {
	return uint(C.avutil_version())
}

//Return the libavutil build-time configuration.
func AvutilConfiguration() string {
	return C.GoString(C.avutil_configuration())
}

//Return the libavutil license.
func AvutilLicense() string {
	return C.GoString(C.avutil_license())
}

//Return a string describing the media_type enum, NULL if media_type is unknown.
func AvGetMediaTypeString(mt MediaType) string {
	return C.GoString(C.av_get_media_type_string((C.enum_AVMediaType)(mt)))
}

//Return a single letter to describe the given picture type pict_type.
func AvGetPictureTypeChar(pt AvPictureType) string {
	return string(C.av_get_picture_type_char((C.enum_AVPictureType)(pt)))
}

//Return x default pointer in case p is NULL.
func AvXIfNull(p, x int) {
	C.av_x_if_null(unsafe.Pointer(&p), unsafe.Pointer(&x))
}

//Compute the length of an integer list.
func AvIntListLengthForSize(e uint, l int, t uint64) uint {
	return uint(C.av_int_list_length_for_size(C.uint(e), unsafe.Pointer(&l), (C.uint64_t)(t)))
}

//Open a file using a UTF-8 filename.
func AvFopenUtf8(p, m string) *File {
	f := C.av_fopen_utf8(C.CString(p), C.CString(m))
	return (*File)(f)
}

//Return the fractional representation of the internal time base.
func AvGetTimeBaseQ() Rational {
	return (Rational)(C.av_get_time_base_q())
}
