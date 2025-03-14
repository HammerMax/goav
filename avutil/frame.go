// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avutil

/*
	#cgo pkg-config: libavutil
	#include <libavutil/frame.h>
	#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"image"
	"log"
	"reflect"
	"unsafe"
)

type (
	AvBuffer            C.struct_AVBuffer
	AvBufferRef         C.struct_AVBufferRef
	AvBufferPool        C.struct_AVBufferPool
	Frame               C.struct_AVFrame
	AvFrameSideData     C.struct_AVFrameSideData
	AvFrameSideDataType C.enum_AVFrameSideDataType
)

func AvprivFrameGetMetadatap(f *Frame) *Dictionary {
	return (*Dictionary)(unsafe.Pointer(f.metadata))
}

//Allocate an Frame and set its fields to default values.
func AvFrameAlloc() *Frame {
	return (*Frame)(unsafe.Pointer(C.av_frame_alloc()))
}

//Free the frame and any dynamically allocated objects in it, e.g.
func AvFrameFree(f *Frame) {
	C.av_frame_free((**C.struct_AVFrame)(unsafe.Pointer(&f)))
}

func (f *Frame) SetFormat(format int32) {
	f.format = C.int(format)
}

func (f *Frame) SetChannelLayout(cl int) {
	f.channel_layout = C.uint64_t(cl)
}

func (f *Frame) SetSampleRate(sr int) {
	f.sample_rate = C.int(sr)
}

func (f *Frame) SetNbSamples(nbSamples int) {
	f.nb_samples = C.int(nbSamples)
}

func (f *Frame) NbSamples() int {
	return (int)(f.nb_samples)
}

func (f *Frame) Width() int32 {
	return (int32)(f.width)
}

func (f *Frame) SetWidth(width int32) {
	f.width = C.int(width)
}

func (f *Frame) SetHeight(height int32) {
	f.height = C.int(height)
}

func (f *Frame) Height() int32 {
	return (int32)(f.height)
}

func (f *Frame) Pts() int64 {
	return int64(f.pts)
}

func (f *Frame) SetPts(pktPts int) {
	f.pts = C.int64_t(pktPts)
}

func (f *Frame) AvFrameGetBuffer(align int32) error {
	return ErrorFromCode(int(C.av_frame_get_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(align))))
}

func (f *Frame) AvFrameUnref() {
	C.av_frame_unref((*C.struct_AVFrame)(f))
}

func (f *Frame) Data() [][]byte {
	var data = make([][]byte, int(C.AV_NUM_DATA_POINTERS))
	for i:=int32(0);i<int32(C.AV_NUM_DATA_POINTERS);i++ {
		header := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(f.data[i])),
			Len:  100000000,
			Cap:  100000000,
		}
		data[i] = *(*[]byte)(unsafe.Pointer(&header))
	}
	return data
}

func (f *Frame) SetDataSimple(x, y int, data uint8) {
	if x >= int(C.AV_NUM_DATA_POINTERS) {
		panic("array out of bound")
	}

	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(f.data[x])),
		Len:  100000000,
		Cap:  100000000, // TODO frame data size 暂时无法获取，业务自己保证不会越界
	}

	a := *(*[]uint8)(unsafe.Pointer(&header))
	a[y] = data
}

func (f *Frame) LineSize() []int32 {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&f.linesize)),
		Len:  int(C.AV_NUM_DATA_POINTERS),
		Cap:  int(C.AV_NUM_DATA_POINTERS),
	}
	return *(*[]int32)(unsafe.Pointer(&header))
}

func (f *Frame) AvFrameFree() {
	C.av_frame_free((**C.struct_AVFrame)(unsafe.Pointer(&f)))
}

//Setup a new reference to the data described by an given frame.
func AvFrameRef(d, s *Frame) int {
	return int(C.av_frame_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s))))
}

//Create a new frame that references the same data as src.
func AvFrameClone(f *Frame) *Frame {
	return (*Frame)(C.av_frame_clone((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

//Unreference all the buffers referenced by frame and reset the frame fields.
func AvFrameUnref(f *Frame) {
	cf := (*C.struct_AVFrame)(unsafe.Pointer(f))
	C.av_frame_unref(cf)
}

//Move everythnig contained in src to dst and reset src.
func AvFrameMoveRef(d, s *Frame) {
	C.av_frame_move_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s)))
}

func (f *Frame) AvFrameMakeWritable() error {
	return ErrorFromCode(int(C.av_frame_make_writable((*C.struct_AVFrame)(unsafe.Pointer(f)))))
}

//Copy only "metadata" fields from src to dst.
func AvFrameCopyProps(d, s *Frame) int {
	return int(C.av_frame_copy_props((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s))))
}

//Get the buffer reference a given data plane is stored in.
func AvFrameGetPlaneBuffer(f *Frame, p int) *AvBufferRef {
	return (*AvBufferRef)(C.av_frame_get_plane_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(p)))
}

//Add a new side data to a frame.
func AvFrameNewSideData(f *Frame, d AvFrameSideDataType, s int) *AvFrameSideData {
	return (*AvFrameSideData)(C.av_frame_new_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(d), C.int(s)))
}

func AvFrameGetSideData(f *Frame, t AvFrameSideDataType) *AvFrameSideData {
	return (*AvFrameSideData)(C.av_frame_get_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(t)))
}

func Data(f *Frame) (data [8]*uint8) {
	for i := range data {
		data[i] = (*uint8)(f.data[i])
	}
	return
}

func Linesize(f *Frame) (linesize [8]int32) {
	for i := range linesize {
		linesize[i] = int32(f.linesize[i])
	}
	return
}

//GetPicture creates a YCbCr image from the frame
func GetPicture(f *Frame) (img *image.YCbCr, err error) {
	// For 4:4:4, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/1.
	// For 4:2:2, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/2.
	// For 4:2:0, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/4.
	// For 4:4:0, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/2.
	// For 4:1:1, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/4.
	// For 4:1:0, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/8.

	w := int(f.linesize[0])
	h := int(f.height)
	r := image.Rectangle{image.Point{0, 0}, image.Point{w, h}}
	// TODO: Use the sub sample ratio from the input image 'f.format'
	img = image.NewYCbCr(r, image.YCbCrSubsampleRatio420)
	// convert the frame data data to a Go byte array
	img.Y = C.GoBytes(unsafe.Pointer(f.data[0]), C.int(w*h))

	wCb := int(f.linesize[1])
	if unsafe.Pointer(f.data[1]) != nil {
		img.Cb = C.GoBytes(unsafe.Pointer(f.data[1]), C.int(wCb*h/2))
	}

	wCr := int(f.linesize[2])
	if unsafe.Pointer(f.data[2]) != nil {
		img.Cr = C.GoBytes(unsafe.Pointer(f.data[2]), C.int(wCr*h/2))
	}
	return
}

// SetPicture sets the image pointer of |f| to the image pointers of |img|
func SetPicture(f *Frame, img *image.YCbCr) {
	d := Data(f)
	// l := Linesize(f)
	// FIXME: Save the original pointers somewhere, this is a memory leak
	d[0] = (*uint8)(unsafe.Pointer(&img.Y[0]))
	// d[1] = (*uint8)(unsafe.Pointer(&img.Cb[0]))
}

func GetPictureRGB(f *Frame) (img *image.RGBA, err error) {
	w := int(f.linesize[0])
	h := int(f.height)
	r := image.Rectangle{image.Point{0, 0}, image.Point{w, h}}
	// TODO: Use the sub sample ratio from the input image 'f.format'
	img = image.NewRGBA(r)
	// convert the frame data data to a Go byte array
	img.Pix = C.GoBytes(unsafe.Pointer(f.data[0]), C.int(w*h))
	img.Stride = w
	log.Println("w", w, "h", h)
	return
}

func AvSetFrame(f *Frame, w int, h int, pixFmt int) (err error) {
	f.width = C.int(w)
	f.height = C.int(h)
	f.format = C.int(pixFmt)
	if ret := C.av_frame_get_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), 32 /*alignment*/); ret < 0 {
		err = fmt.Errorf("Error allocating avframe buffer. Err: %v", ret)
		return
	}
	return
}

func AvFrameGetInfo(f *Frame) (width int, height int, linesize [8]int32, data [8]*uint8) {
	width = int(f.linesize[0])
	height = int(f.height)
	for i := range linesize {
		linesize[i] = int32(f.linesize[i])
	}
	for i := range data {
		data[i] = (*uint8)(f.data[i])
	}
	// log.Println("Linesize is ", f.linesize, "Data is", data)
	return
}

func GetBestEffortTimestamp(f *Frame) int64 {
	return int64(f.best_effort_timestamp)
}

// //static int get_video_buffer (Frame *frame, int align)
// func GetVideoBuffer(f *Frame, a int) int {
// 	return int(C.get_video_buffer(f, C.int(a)))
// }

// //static int get_audio_buffer (Frame *frame, int align)
// func GetAudioBuffer(f *Frame, a int) int {
// 	return C.get_audio_buffer(f, C.int(a))
// }

// //static void get_frame_defaults (Frame *frame)
// func GetFrameDefaults(f *Frame) {
// 	C.get_frame_defaults(*C.struct_AVFrame(f))
// }
