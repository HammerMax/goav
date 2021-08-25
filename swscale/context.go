// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package swscale

//#cgo pkg-config: libswscale libavutil
//#include <libswscale/swscale.h>
import "C"
import (
	"github.com/HammerMax/goav/avfilter"
	"github.com/HammerMax/goav/avutil"
	"reflect"
	"unsafe"
)

//Allocate an empty Context.
func SwsAllocContext() *Context {
	return (*Context)(C.sws_alloc_context())
}

//Initialize the swscaler context sws_context.
func SwsInitContext(ctxt *Context, sf, df *Filter) int {
	return int(C.sws_init_context((*C.struct_SwsContext)(ctxt), (*C.struct_SwsFilter)(sf), (*C.struct_SwsFilter)(df)))
}

//Free the swscaler context swsContext.
func SwsFreecontext(ctxt *Context) {
	C.sws_freeContext((*C.struct_SwsContext)(ctxt))
}

//Allocate and return an Context.
func SwsGetContext(sw, sh int, sf avutil.PixelFormat, dw, dh int, df avutil.PixelFormat, f int, sfl, dfl *avfilter.Filter, param []float64) *Context {
	h := (*reflect.SliceHeader)(unsafe.Pointer(&param))
	return (*Context)(C.sws_getContext(C.int(sw), C.int(sh), (C.enum_AVPixelFormat)(sf), C.int(dw), C.int(dh), (C.enum_AVPixelFormat)(df), C.int(f), (*C.struct_SwsFilter)(unsafe.Pointer(sfl)), (*C.struct_SwsFilter)(unsafe.Pointer(dfl)), (*C.double)(unsafe.Pointer(h.Data))))
}

//Check if context can be reused, otherwise reallocate a new one.
func SwsGetcachedcontext(ctxt *Context, sw, sh int, sf PixelFormat, dw, dh int, df PixelFormat, f int, sfl, dfl *Filter, p *float64) *Context {
	return (*Context)(C.sws_getCachedContext((*C.struct_SwsContext)(ctxt), C.int(sw), C.int(sh), (C.enum_AVPixelFormat)(sf), C.int(dw), C.int(dh), (C.enum_AVPixelFormat)(df), C.int(f), (*C.struct_SwsFilter)(sfl), (*C.struct_SwsFilter)(dfl), (*C.double)(p)))
}
