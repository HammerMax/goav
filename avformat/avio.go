package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
/*
typedef int (*av_io_read_write)(void *opaque, uint8_t *buf, int buf_size);
typedef int64_t (*av_io_seek)(void *opaque, int64_t offset, int whence);

AVIOContext *avio_alloc_context(
                  unsigned char *buffer,
                  int buffer_size,
                  int write_flag,
                  void *opaque,
                  int (*read_packet)(void *opaque, uint8_t *buf, int buf_size),
                  int (*write_packet)(void *opaque, uint8_t *buf, int buf_size),
                  int64_t (*seek)(void *opaque, int64_t offset, int whence));
 */
import "C"
import (
	"reflect"
	"unsafe"
)

func avioReadWrite(opaque unsafe.Pointer) {

}

func convertToAVIOReadWrite(f func(interface{}, []byte) int) C.av_io_read_write {
	return nil
}

func convertToAVIOSeek(f func(opaque interface{}, offset, whence int) int) C.av_io_seek {
	return nil
}

func AVIOAllocContext(buffer []byte, writeFlag int32, opaque interface{}, readPacket, writePacket func(interface{}, []byte) int, seek func(opaque interface{}, offset, whence int) int) *AvIOContext {
	bufferHeader := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	C.avio_alloc_context((*C.uchar)(unsafe.Pointer(bufferHeader.Data)), C.int(len(buffer)), C.int(writeFlag), unsafe.Pointer(reflect.ValueOf(opaque).Pointer()), convertToAVIOReadWrite(readPacket), convertToAVIOReadWrite(writePacket), convertToAVIOSeek(seek))
	return nil
}