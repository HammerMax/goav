// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avcodec contains the codecs (decoders and encoders) provided by the libavcodec library
//Provides some generic global options, which can be set on all the encoders and decoders.
package avcodec

//#cgo pkg-config: libavcodec libavutil
//#include <libavutil/avutil.h>
//#include <libswscale/swscale.h>
import "C"

const (

	SWS_FAST_BILINEAR        = C.SWS_FAST_BILINEAR
	SWS_BILINEAR             = C.SWS_BILINEAR
	SWS_BICUBIC              = C.SWS_BICUBIC
	SWS_X                    = C.SWS_X
	SWS_POINT                = C.SWS_POINT
	SWS_AREA                 = C.SWS_AREA
	SWS_BICUBLIN             = C.SWS_BICUBLIN
	SWS_GAUSS                = C.SWS_GAUSS
	SWS_SINC                 = C.SWS_SINC
	SWS_LANCZOS              = C.SWS_LANCZOS
	SWS_SPLINE               = C.SWS_SPLINE
	SWS_SRC_V_CHR_DROP_MASK  = C.SWS_SRC_V_CHR_DROP_MASK
	SWS_SRC_V_CHR_DROP_SHIFT = C.SWS_SRC_V_CHR_DROP_SHIFT
	SWS_PARAM_DEFAULT        = C.SWS_PARAM_DEFAULT
	SWS_PRINT_INFO           = C.SWS_PRINT_INFO
	SWS_FULL_CHR_H_INT       = C.SWS_FULL_CHR_H_INT
	SWS_FULL_CHR_H_INP       = C.SWS_FULL_CHR_H_INP
	SWS_DIRECT_BGR           = C.SWS_DIRECT_BGR
	SWS_ACCURATE_RND         = C.SWS_ACCURATE_RND
	SWS_BITEXACT             = C.SWS_BITEXACT
	SWS_ERROR_DIFFUSION      = C.SWS_ERROR_DIFFUSION
	SWS_MAX_REDUCE_CUTOFF    = C.SWS_MAX_REDUCE_CUTOFF
	SWS_CS_ITU709            = C.SWS_CS_ITU709
	SWS_CS_FCC               = C.SWS_CS_FCC
	SWS_CS_ITU601            = C.SWS_CS_ITU601
	SWS_CS_ITU624            = C.SWS_CS_ITU624
	SWS_CS_SMPTE170M         = C.SWS_CS_SMPTE170M
	SWS_CS_SMPTE240M         = C.SWS_CS_SMPTE240M
	SWS_CS_DEFAULT           = C.SWS_CS_DEFAULT
	SWS_CS_BT2020            = C.SWS_CS_BT2020
)