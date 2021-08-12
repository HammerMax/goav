// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"reflect"
	"unsafe"
)

func (p *Packet) Buf() *AvBufferRef {
	return (*AvBufferRef)(p.buf)
}
func (p *Packet) Duration() int {
	return int(p.duration)
}
func (p *Packet) Flags() int {
	return int(p.flags)
}
func (p *Packet) SetFlags(flags int) {
	p.flags = C.int(flags)
}
func (p *Packet) SideDataElems() int {
	return int(p.side_data_elems)
}

func (p *Packet) Size() int {
	return int(p.size)
}

func (p *Packet) SizePoint() *int {
	return (*int)(unsafe.Pointer(&p.size))
}

func (p *Packet) StreamIndex() int {
	return int(p.stream_index)
}
func (p *Packet) SetStreamIndex(idx int) {
	p.stream_index = C.int(idx)
}
func (p *Packet) ConvergenceDuration() int64 {
	return int64(p.convergence_duration)
}
func (p *Packet) Dts() int64 {
	return int64(p.dts)
}
func (p *Packet) SetDts(dts int64) {
	p.dts = C.int64_t(dts)
}
func (p *Packet) Pos() int64 {
	return int64(p.pos)
}
func (p *Packet) Pts() int64 {
	return int64(p.pts)
}

func (p *Packet) SetPts(pts int64) {
	p.dts = C.int64_t(pts)
}

func (p *Packet) Data() *uint8 {
	return (*uint8)(p.data)
}

func (p *Packet) DataPoint() **uint8 {
	return (**uint8)(unsafe.Pointer(&p.data))
}

// DataSlice 使用 data 和 size 属性，返回[]byte
func (p *Packet) DataSlice() []byte {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p.data)),
		Len:  p.Size(),
		Cap:  p.Size(),
	}
	return *(*[]byte)(unsafe.Pointer(&header))
}
