// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package swscale performs highly optimized image scaling and colorspace and pixel format conversion operations.
//Rescaling: is the process of changing the video size. Several rescaling options and algorithms are available.
//Pixel format conversion: is the process of converting the image format and colorspace of the image.
package swscale

//#cgo pkg-config: libswscale libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <string.h>
//#include <stdint.h>
//#include <libswscale/swscale.h>
import "C"
import (
	"reflect"
	"unsafe"
)

type (
	Context     C.struct_SwsContext
	Filter      C.struct_SwsFilter
	Vector      C.struct_SwsVector
	Class       C.struct_AVClass
	PixelFormat C.enum_AVPixelFormat
)

const (
 SWS_FAST_BILINEAR        =C.SWS_FAST_BILINEAR
 SWS_BILINEAR             =C.SWS_BILINEAR
 SWS_BICUBIC              =C.SWS_BICUBIC
 SWS_X                    =C.SWS_X
 SWS_POINT                =C.SWS_POINT
 SWS_AREA                 =C.SWS_AREA
 SWS_BICUBLIN             =C.SWS_BICUBLIN
 SWS_GAUSS                =C.SWS_GAUSS
 SWS_SINC                 =C.SWS_SINC
 SWS_LANCZOS              =C.SWS_LANCZOS
 SWS_SPLINE               =C.SWS_SPLINE
 SWS_SRC_V_CHR_DROP_MASK  =C.SWS_SRC_V_CHR_DROP_MASK
 SWS_SRC_V_CHR_DROP_SHIFT =C.SWS_SRC_V_CHR_DROP_SHIFT
 SWS_PARAM_DEFAULT        =C.SWS_PARAM_DEFAULT
 SWS_PRINT_INFO           =C.SWS_PRINT_INFO
 SWS_FULL_CHR_H_INT       =C.SWS_FULL_CHR_H_INT
 SWS_FULL_CHR_H_INP       =C.SWS_FULL_CHR_H_INP
 SWS_DIRECT_BGR           =C.SWS_DIRECT_BGR
 SWS_ACCURATE_RND         =C.SWS_ACCURATE_RND
 SWS_BITEXACT             =C.SWS_BITEXACT
 SWS_ERROR_DIFFUSION      =C.SWS_ERROR_DIFFUSION
 SWS_MAX_REDUCE_CUTOFF    =C.SWS_MAX_REDUCE_CUTOFF
 SWS_CS_ITU709            =C.SWS_CS_ITU709
 SWS_CS_FCC               =C.SWS_CS_FCC
 SWS_CS_ITU601            =C.SWS_CS_ITU601
 SWS_CS_ITU624            =C.SWS_CS_ITU624
 SWS_CS_SMPTE170M         =C.SWS_CS_SMPTE170M
 SWS_CS_SMPTE240M         =C.SWS_CS_SMPTE240M
 SWS_CS_DEFAULT           =C.SWS_CS_DEFAULT
 SWS_CS_BT2020            =C.SWS_CS_BT2020
)

//Return the LIBSWSCALE_VERSION_INT constant.
func SwscaleVersion() uint {
	return uint(C.swscale_version())
}

//Return the libswscale build-time configuration.
func SwscaleConfiguration() string {
	return C.GoString(C.swscale_configuration())
}

//Return the libswscale license.
func SwscaleLicense() string {
	return C.GoString(C.swscale_license())
}

//Return a pointer to yuv<->rgb coefficients for the given colorspace suitable for sws_setColorspaceDetails().
func SwsGetcoefficients(c int) *int {
	return (*int)(unsafe.Pointer(C.sws_getCoefficients(C.int(c))))
}

//Return a positive value if pix_fmt is a supported input format, 0 otherwise.
func SwsIssupportedinput(p PixelFormat) int {
	return int(C.sws_isSupportedInput((C.enum_AVPixelFormat)(p)))
}

//Return a positive value if pix_fmt is a supported output format, 0 otherwise.
func SwsIssupportedoutput(p PixelFormat) int {
	return int(C.sws_isSupportedOutput((C.enum_AVPixelFormat)(p)))
}

func SwsIssupportedendiannessconversion(p PixelFormat) int {
	return int(C.sws_isSupportedEndiannessConversion((C.enum_AVPixelFormat)(p)))
}

////Scale the image slice in srcSlice and put the resulting scaled slice in the image in dst.
func SwsScale(ctxt *Context, src *uint8, str int, y, h int, d *uint8, ds int) int {
	cctxt := (*C.struct_SwsContext)(unsafe.Pointer(ctxt))
	csrc := (*C.uint8_t)(unsafe.Pointer(src))
	cstr := (*C.int)(unsafe.Pointer(&str))
	cd := (*C.uint8_t)(unsafe.Pointer(d))
	cds := (*C.int)(unsafe.Pointer(&ds))
	return int(C.sws_scale(cctxt, &csrc, cstr, C.int(y), C.int(h), &cd, cds))
}

func SwsScale2(ctxt *Context, srcData [8]*uint8, srcStride [8]int32, y, h int32, dstData [8]*uint8,
	dstStride [8]int32) int {
	cctxt := (*C.struct_SwsContext)(unsafe.Pointer(ctxt))
	csrc := (**C.uint8_t)(unsafe.Pointer(&srcData[0]))
	cstr := (*C.int)(unsafe.Pointer(&srcStride[0]))
	cd := (**C.uint8_t)(unsafe.Pointer(&dstData[0]))
	cds := (*C.int)(unsafe.Pointer(&dstStride))
	return int(C.sws_scale(cctxt, csrc, cstr, C.int(y), C.int(h), cd, cds))
}

func (c *Context) SwsScale(srcSlice [][]byte, srcStride []int32, srcSliceY, srcSliceH int, dst [][]byte, dstStride []int32) int {
	var srcSlicePointers []unsafe.Pointer
	for k := range srcSlice {
		srcSlicePointers = append(srcSlicePointers, unsafe.Pointer(&srcSlice[k][0]))
	}

	var dstSlicePointers []unsafe.Pointer
	for k := range dst {
		dstSlicePointers = append(dstSlicePointers, unsafe.Pointer(&dst[k][0]))
	}

	ret := int(C.sws_scale((*C.struct_SwsContext)(unsafe.Pointer(c)), (**C.uint8_t)(unsafe.Pointer(&srcSlicePointers[0])), (*C.int)(unsafe.Pointer(&srcStride[0])), C.int(srcSliceY), C.int(srcSliceH), (**C.uint8_t)(unsafe.Pointer(&dstSlicePointers[0])), (*C.int)(unsafe.Pointer(&dstStride[0]))))

	for k, v := range dstSlicePointers {
		header := (*reflect.SliceHeader)(unsafe.Pointer(&dst[k]))
		header.Data = uintptr(v)
		header.Len = 10000000
		header.Cap = 10000000
	}
	return ret
}

func SwsSetcolorspacedetails(ctxt *Context, it *int, sr int, t *int, dr, b, c, s int) int {
	cit := (*C.int)(unsafe.Pointer(it))
	ct := (*C.int)(unsafe.Pointer(t))
	return int(C.sws_setColorspaceDetails((*C.struct_SwsContext)(ctxt), cit, C.int(sr), ct, C.int(dr), C.int(b), C.int(c), C.int(s)))
}

func SwsGetcolorspacedetails(ctxt *Context, it, sr, t, dr, b, c, s *int) int {
	cit := (**C.int)(unsafe.Pointer(it))
	csr := (*C.int)(unsafe.Pointer(sr))
	ct := (**C.int)(unsafe.Pointer(t))
	cdr := (*C.int)(unsafe.Pointer(dr))
	cb := (*C.int)(unsafe.Pointer(b))
	cc := (*C.int)(unsafe.Pointer(c))
	cs := (*C.int)(unsafe.Pointer(s))
	return int(C.sws_getColorspaceDetails((*C.struct_SwsContext)(ctxt), cit, csr, ct, cdr, cb, cc, cs))
}

func SwsGetdefaultfilter(lb, cb, ls, cs, chs, cvs float32, v int) *Filter {
	return (*Filter)(unsafe.Pointer(C.sws_getDefaultFilter(C.float(lb), C.float(cb), C.float(ls), C.float(cs), C.float(chs), C.float(cvs), C.int(v))))
}

func SwsFreefilter(f *Filter) {
	C.sws_freeFilter((*C.struct_SwsFilter)(f))
}

//Convert an 8-bit paletted frame into a frame with a color depth of 32 bits.
func SwsConvertpalette8topacked32(s, d *uint8, px int, p *uint8) {
	C.sws_convertPalette8ToPacked32((*C.uint8_t)(s), (*C.uint8_t)(d), C.int(px), (*C.uint8_t)(p))
}

//Convert an 8-bit paletted frame into a frame with a color depth of 24 bits.
func SwsConvertpalette8topacked24(s, d *uint8, px int, p *uint8) {
	C.sws_convertPalette8ToPacked24((*C.uint8_t)(s), (*C.uint8_t)(d), C.int(px), (*C.uint8_t)(p))
}

//Get the Class for swsContext.
func SwsGetClass() *Class {
	return (*Class)(C.sws_get_class())
}
