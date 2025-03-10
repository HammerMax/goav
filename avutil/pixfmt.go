package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/pixfmt.h>
import "C"

const (
	AV_PIX_FMT_BAYER_GRBG16BE=PixelFormat(C.AV_PIX_FMT_BAYER_GRBG16BE)
	AV_PIX_FMT_NONE          =PixelFormat(C.AV_PIX_FMT_NONE)
	AV_PIX_FMT_YUV420P       =PixelFormat(C.AV_PIX_FMT_YUV420P)
	AV_PIX_FMT_YUYV422       =PixelFormat(C.AV_PIX_FMT_YUYV422)
	AV_PIX_FMT_RGB24         =PixelFormat(C.AV_PIX_FMT_RGB24)
	AV_PIX_FMT_BGR24         =PixelFormat(C.AV_PIX_FMT_BGR24)
	AV_PIX_FMT_YUV422P       =PixelFormat(C.AV_PIX_FMT_YUV422P)
	AV_PIX_FMT_YUV444P       =PixelFormat(C.AV_PIX_FMT_YUV444P)
	AV_PIX_FMT_YUV410P       =PixelFormat(C.AV_PIX_FMT_YUV410P)
	AV_PIX_FMT_YUV411P       =PixelFormat(C.AV_PIX_FMT_YUV411P)
	AV_PIX_FMT_GRAY8         =PixelFormat(C.AV_PIX_FMT_GRAY8)
	AV_PIX_FMT_MONOWHITE     =PixelFormat(C.AV_PIX_FMT_MONOWHITE)
	AV_PIX_FMT_MONOBLACK     =PixelFormat(C.AV_PIX_FMT_MONOBLACK)
	AV_PIX_FMT_PAL8          =PixelFormat(C.AV_PIX_FMT_PAL8)
	AV_PIX_FMT_YUVJ420P      =PixelFormat(C.AV_PIX_FMT_YUVJ420P)
	AV_PIX_FMT_YUVJ422P      =PixelFormat(C.AV_PIX_FMT_YUVJ422P)
	AV_PIX_FMT_YUVJ444P      =PixelFormat(C.AV_PIX_FMT_YUVJ444P)
	AV_PIX_FMT_UYVY422       =PixelFormat(C.AV_PIX_FMT_UYVY422)
	AV_PIX_FMT_UYYVYY411     =PixelFormat(C.AV_PIX_FMT_UYYVYY411)
	AV_PIX_FMT_BGR8          =PixelFormat(C.AV_PIX_FMT_BGR8)
	AV_PIX_FMT_BGR4          =PixelFormat(C.AV_PIX_FMT_BGR4)
	AV_PIX_FMT_BGR4_BYTE     =PixelFormat(C.AV_PIX_FMT_BGR4_BYTE)
	AV_PIX_FMT_RGB8          =PixelFormat(C.AV_PIX_FMT_RGB8)
	AV_PIX_FMT_RGB4          =PixelFormat(C.AV_PIX_FMT_RGB4)
	AV_PIX_FMT_RGB4_BYTE     =PixelFormat(C.AV_PIX_FMT_RGB4_BYTE)
	AV_PIX_FMT_NV12          =PixelFormat(C.AV_PIX_FMT_NV12)
	AV_PIX_FMT_NV21          =PixelFormat(C.AV_PIX_FMT_NV21)
	AV_PIX_FMT_ARGB          =PixelFormat(C.AV_PIX_FMT_ARGB)
	AV_PIX_FMT_RGBA          =PixelFormat(C.AV_PIX_FMT_RGBA)
	AV_PIX_FMT_ABGR          =PixelFormat(C.AV_PIX_FMT_ABGR)
	AV_PIX_FMT_BGRA          =PixelFormat(C.AV_PIX_FMT_BGRA)
	AV_PIX_FMT_GRAY16BE      =PixelFormat(C.AV_PIX_FMT_GRAY16BE)
	AV_PIX_FMT_GRAY16LE      =PixelFormat(C.AV_PIX_FMT_GRAY16LE)
	AV_PIX_FMT_YUV440P       =PixelFormat(C.AV_PIX_FMT_YUV440P)
	AV_PIX_FMT_YUVJ440P      =PixelFormat(C.AV_PIX_FMT_YUVJ440P)
	AV_PIX_FMT_YUVA420P      =PixelFormat(C.AV_PIX_FMT_YUVA420P)
	AV_PIX_FMT_RGB48BE       =PixelFormat(C.AV_PIX_FMT_RGB48BE)
	AV_PIX_FMT_RGB48LE       =PixelFormat(C.AV_PIX_FMT_RGB48LE)
	AV_PIX_FMT_RGB565BE      =PixelFormat(C.AV_PIX_FMT_RGB565BE)
	AV_PIX_FMT_RGB565LE      =PixelFormat(C.AV_PIX_FMT_RGB565LE)
	AV_PIX_FMT_RGB555BE      =PixelFormat(C.AV_PIX_FMT_RGB555BE)
	AV_PIX_FMT_RGB555LE      =PixelFormat(C.AV_PIX_FMT_RGB555LE)
	AV_PIX_FMT_BGR565BE      =PixelFormat(C.AV_PIX_FMT_BGR565BE)
	AV_PIX_FMT_BGR565LE      =PixelFormat(C.AV_PIX_FMT_BGR565LE)
	AV_PIX_FMT_BGR555BE      =PixelFormat(C.AV_PIX_FMT_BGR555BE)
	AV_PIX_FMT_BGR555LE      =PixelFormat(C.AV_PIX_FMT_BGR555LE)
	AV_PIX_FMT_VAAPI_MOCO    =PixelFormat(C.AV_PIX_FMT_VAAPI_MOCO)
	AV_PIX_FMT_VAAPI_IDCT    =PixelFormat(C.AV_PIX_FMT_VAAPI_IDCT)
	AV_PIX_FMT_VAAPI_VLD     =PixelFormat(C.AV_PIX_FMT_VAAPI_VLD)
	AV_PIX_FMT_VAAPI         =PixelFormat(C.AV_PIX_FMT_VAAPI)
	AV_PIX_FMT_YUV420P16LE   =PixelFormat(C.AV_PIX_FMT_YUV420P16LE)
	AV_PIX_FMT_YUV420P16BE   =PixelFormat(C.AV_PIX_FMT_YUV420P16BE)
	AV_PIX_FMT_YUV422P16LE   =PixelFormat(C.AV_PIX_FMT_YUV422P16LE)
	AV_PIX_FMT_YUV422P16BE   =PixelFormat(C.AV_PIX_FMT_YUV422P16BE)
	AV_PIX_FMT_YUV444P16LE   =PixelFormat(C.AV_PIX_FMT_YUV444P16LE)
	AV_PIX_FMT_YUV444P16BE   =PixelFormat(C.AV_PIX_FMT_YUV444P16BE)
	AV_PIX_FMT_DXVA2_VLD     =PixelFormat(C.AV_PIX_FMT_DXVA2_VLD)
	AV_PIX_FMT_RGB444LE      =PixelFormat(C.AV_PIX_FMT_RGB444LE)
	AV_PIX_FMT_RGB444BE      =PixelFormat(C.AV_PIX_FMT_RGB444BE)
	AV_PIX_FMT_BGR444LE      =PixelFormat(C.AV_PIX_FMT_BGR444LE)
	AV_PIX_FMT_BGR444BE      =PixelFormat(C.AV_PIX_FMT_BGR444BE)
	AV_PIX_FMT_YA8           =PixelFormat(C.AV_PIX_FMT_YA8)
	AV_PIX_FMT_Y400A         =PixelFormat(C.AV_PIX_FMT_Y400A)
	AV_PIX_FMT_GRAY8A        =PixelFormat(C.AV_PIX_FMT_GRAY8A)
	AV_PIX_FMT_BGR48BE       =PixelFormat(C.AV_PIX_FMT_BGR48BE)
	AV_PIX_FMT_BGR48LE       =PixelFormat(C.AV_PIX_FMT_BGR48LE)
	AV_PIX_FMT_YUV420P9BE    =PixelFormat(C.AV_PIX_FMT_YUV420P9BE)
	AV_PIX_FMT_YUV420P9LE    =PixelFormat(C.AV_PIX_FMT_YUV420P9LE)
	AV_PIX_FMT_YUV420P10BE   =PixelFormat(C.AV_PIX_FMT_YUV420P10BE)
	AV_PIX_FMT_YUV420P10LE   =PixelFormat(C.AV_PIX_FMT_YUV420P10LE)
	AV_PIX_FMT_YUV422P10BE   =PixelFormat(C.AV_PIX_FMT_YUV422P10BE)
	AV_PIX_FMT_YUV422P10LE   =PixelFormat(C.AV_PIX_FMT_YUV422P10LE)
	AV_PIX_FMT_YUV444P9BE    =PixelFormat(C.AV_PIX_FMT_YUV444P9BE)
	AV_PIX_FMT_YUV444P9LE    =PixelFormat(C.AV_PIX_FMT_YUV444P9LE)
	AV_PIX_FMT_YUV444P10BE   =PixelFormat(C.AV_PIX_FMT_YUV444P10BE)
	AV_PIX_FMT_YUV444P10LE   =PixelFormat(C.AV_PIX_FMT_YUV444P10LE)
	AV_PIX_FMT_YUV422P9BE    =PixelFormat(C.AV_PIX_FMT_YUV422P9BE)
	AV_PIX_FMT_YUV422P9LE    =PixelFormat(C.AV_PIX_FMT_YUV422P9LE)
	AV_PIX_FMT_GBRP          =PixelFormat(C.AV_PIX_FMT_GBRP)
	AV_PIX_FMT_GBR24P        =PixelFormat(C.AV_PIX_FMT_GBR24P)
	AV_PIX_FMT_GBRP9BE       =PixelFormat(C.AV_PIX_FMT_GBRP9BE)
	AV_PIX_FMT_GBRP9LE       =PixelFormat(C.AV_PIX_FMT_GBRP9LE)
	AV_PIX_FMT_GBRP10BE      =PixelFormat(C.AV_PIX_FMT_GBRP10BE)
	AV_PIX_FMT_GBRP10LE      =PixelFormat(C.AV_PIX_FMT_GBRP10LE)
	AV_PIX_FMT_GBRP16BE      =PixelFormat(C.AV_PIX_FMT_GBRP16BE)
	AV_PIX_FMT_GBRP16LE      =PixelFormat(C.AV_PIX_FMT_GBRP16LE)
	AV_PIX_FMT_YUVA422P      =PixelFormat(C.AV_PIX_FMT_YUVA422P)
	AV_PIX_FMT_YUVA444P      =PixelFormat(C.AV_PIX_FMT_YUVA444P)
	AV_PIX_FMT_YUVA420P9BE   =PixelFormat(C.AV_PIX_FMT_YUVA420P9BE)
	AV_PIX_FMT_YUVA420P9LE   =PixelFormat(C.AV_PIX_FMT_YUVA420P9LE)
	AV_PIX_FMT_YUVA422P9BE   =PixelFormat(C.AV_PIX_FMT_YUVA422P9BE)
	AV_PIX_FMT_YUVA422P9LE   =PixelFormat(C.AV_PIX_FMT_YUVA422P9LE)
	AV_PIX_FMT_YUVA444P9BE   =PixelFormat(C.AV_PIX_FMT_YUVA444P9BE)
	AV_PIX_FMT_YUVA444P9LE   =PixelFormat(C.AV_PIX_FMT_YUVA444P9LE)
	AV_PIX_FMT_YUVA420P10BE  =PixelFormat(C.AV_PIX_FMT_YUVA420P10BE)
	AV_PIX_FMT_YUVA420P10LE  =PixelFormat(C.AV_PIX_FMT_YUVA420P10LE)
	AV_PIX_FMT_YUVA422P10BE  =PixelFormat(C.AV_PIX_FMT_YUVA422P10BE)
	AV_PIX_FMT_YUVA422P10LE  =PixelFormat(C.AV_PIX_FMT_YUVA422P10LE)
	AV_PIX_FMT_YUVA444P10BE  =PixelFormat(C.AV_PIX_FMT_YUVA444P10BE)
	AV_PIX_FMT_YUVA444P10LE  =PixelFormat(C.AV_PIX_FMT_YUVA444P10LE)
	AV_PIX_FMT_YUVA420P16BE  =PixelFormat(C.AV_PIX_FMT_YUVA420P16BE)
	AV_PIX_FMT_YUVA420P16LE  =PixelFormat(C.AV_PIX_FMT_YUVA420P16LE)
	AV_PIX_FMT_YUVA422P16BE  =PixelFormat(C.AV_PIX_FMT_YUVA422P16BE)
	AV_PIX_FMT_YUVA422P16LE  =PixelFormat(C.AV_PIX_FMT_YUVA422P16LE)
	AV_PIX_FMT_YUVA444P16BE  =PixelFormat(C.AV_PIX_FMT_YUVA444P16BE)
	AV_PIX_FMT_YUVA444P16LE  =PixelFormat(C.AV_PIX_FMT_YUVA444P16LE)
	AV_PIX_FMT_VDPAU         =PixelFormat(C.AV_PIX_FMT_VDPAU)
	AV_PIX_FMT_XYZ12LE       =PixelFormat(C.AV_PIX_FMT_XYZ12LE)
	AV_PIX_FMT_XYZ12BE       =PixelFormat(C.AV_PIX_FMT_XYZ12BE)
	AV_PIX_FMT_NV16          =PixelFormat(C.AV_PIX_FMT_NV16)
	AV_PIX_FMT_NV20LE        =PixelFormat(C.AV_PIX_FMT_NV20LE)
	AV_PIX_FMT_NV20BE        =PixelFormat(C.AV_PIX_FMT_NV20BE)
	AV_PIX_FMT_RGBA64BE      =PixelFormat(C.AV_PIX_FMT_RGBA64BE)
	AV_PIX_FMT_RGBA64LE      =PixelFormat(C.AV_PIX_FMT_RGBA64LE)
	AV_PIX_FMT_BGRA64BE      =PixelFormat(C.AV_PIX_FMT_BGRA64BE)
	AV_PIX_FMT_BGRA64LE      =PixelFormat(C.AV_PIX_FMT_BGRA64LE)
	AV_PIX_FMT_YVYU422       =PixelFormat(C.AV_PIX_FMT_YVYU422)
	AV_PIX_FMT_YA16BE        =PixelFormat(C.AV_PIX_FMT_YA16BE)
	AV_PIX_FMT_YA16LE        =PixelFormat(C.AV_PIX_FMT_YA16LE)
	AV_PIX_FMT_GBRAP         =PixelFormat(C.AV_PIX_FMT_GBRAP)
	AV_PIX_FMT_GBRAP16BE     =PixelFormat(C.AV_PIX_FMT_GBRAP16BE)
	AV_PIX_FMT_GBRAP16LE     =PixelFormat(C.AV_PIX_FMT_GBRAP16LE)
	AV_PIX_FMT_QSV           =PixelFormat(C.AV_PIX_FMT_QSV)
	AV_PIX_FMT_MMAL          =PixelFormat(C.AV_PIX_FMT_MMAL)
	AV_PIX_FMT_D3D11VA_VLD   =PixelFormat(C.AV_PIX_FMT_D3D11VA_VLD)
	AV_PIX_FMT_CUDA          =PixelFormat(C.AV_PIX_FMT_CUDA)
	AV_PIX_FMT_0RGB          =PixelFormat(C.AV_PIX_FMT_0RGB)
	AV_PIX_FMT_RGB0          =PixelFormat(C.AV_PIX_FMT_RGB0)
	AV_PIX_FMT_0BGR          =PixelFormat(C.AV_PIX_FMT_0BGR)
	AV_PIX_FMT_BGR0          =PixelFormat(C.AV_PIX_FMT_BGR0)
	AV_PIX_FMT_YUV420P12BE   =PixelFormat(C.AV_PIX_FMT_YUV420P12BE)
	AV_PIX_FMT_YUV420P12LE   =PixelFormat(C.AV_PIX_FMT_YUV420P12LE)
	AV_PIX_FMT_YUV420P14BE   =PixelFormat(C.AV_PIX_FMT_YUV420P14BE)
	AV_PIX_FMT_YUV420P14LE   =PixelFormat(C.AV_PIX_FMT_YUV420P14LE)
	AV_PIX_FMT_YUV422P12BE   =PixelFormat(C.AV_PIX_FMT_YUV422P12BE)
	AV_PIX_FMT_YUV422P12LE   =PixelFormat(C.AV_PIX_FMT_YUV422P12LE)
	AV_PIX_FMT_YUV422P14BE   =PixelFormat(C.AV_PIX_FMT_YUV422P14BE)
	AV_PIX_FMT_YUV422P14LE   =PixelFormat(C.AV_PIX_FMT_YUV422P14LE)
	AV_PIX_FMT_YUV444P12BE   =PixelFormat(C.AV_PIX_FMT_YUV444P12BE)
	AV_PIX_FMT_YUV444P12LE   =PixelFormat(C.AV_PIX_FMT_YUV444P12LE)
	AV_PIX_FMT_YUV444P14BE   =PixelFormat(C.AV_PIX_FMT_YUV444P14BE)
	AV_PIX_FMT_YUV444P14LE   =PixelFormat(C.AV_PIX_FMT_YUV444P14LE)
	AV_PIX_FMT_GBRP12BE      =PixelFormat(C.AV_PIX_FMT_GBRP12BE)
	AV_PIX_FMT_GBRP12LE      =PixelFormat(C.AV_PIX_FMT_GBRP12LE)
	AV_PIX_FMT_GBRP14BE      =PixelFormat(C.AV_PIX_FMT_GBRP14BE)
	AV_PIX_FMT_GBRP14LE      =PixelFormat(C.AV_PIX_FMT_GBRP14LE)
	AV_PIX_FMT_YUVJ411P      =PixelFormat(C.AV_PIX_FMT_YUVJ411P)
	AV_PIX_FMT_BAYER_BGGR8   =PixelFormat(C.AV_PIX_FMT_BAYER_BGGR8)
	AV_PIX_FMT_BAYER_RGGB8   =PixelFormat(C.AV_PIX_FMT_BAYER_RGGB8)
	AV_PIX_FMT_BAYER_GBRG8   =PixelFormat(C.AV_PIX_FMT_BAYER_GBRG8)
	AV_PIX_FMT_BAYER_GRBG8   =PixelFormat(C.AV_PIX_FMT_BAYER_GRBG8)
	AV_PIX_FMT_BAYER_BGGR16LE=PixelFormat(C.AV_PIX_FMT_BAYER_BGGR16LE)
	AV_PIX_FMT_BAYER_BGGR16BE=PixelFormat(C.AV_PIX_FMT_BAYER_BGGR16BE)
	AV_PIX_FMT_BAYER_RGGB16LE=PixelFormat(C.AV_PIX_FMT_BAYER_RGGB16LE)
	AV_PIX_FMT_BAYER_RGGB16BE=PixelFormat(C.AV_PIX_FMT_BAYER_RGGB16BE)
	AV_PIX_FMT_BAYER_GBRG16LE=PixelFormat(C.AV_PIX_FMT_BAYER_GBRG16LE)
	AV_PIX_FMT_BAYER_GBRG16BE=PixelFormat(C.AV_PIX_FMT_BAYER_GBRG16BE)
	AV_PIX_FMT_BAYER_GRBG16LE=PixelFormat(C.AV_PIX_FMT_BAYER_GRBG16LE)
	AV_PIX_FMT_XVMC          =PixelFormat(C.AV_PIX_FMT_XVMC)
	AV_PIX_FMT_YUV440P10LE   =PixelFormat(C.AV_PIX_FMT_YUV440P10LE)
	AV_PIX_FMT_YUV440P10BE   =PixelFormat(C.AV_PIX_FMT_YUV440P10BE)
	AV_PIX_FMT_YUV440P12LE   =PixelFormat(C.AV_PIX_FMT_YUV440P12LE)
	AV_PIX_FMT_YUV440P12BE   =PixelFormat(C.AV_PIX_FMT_YUV440P12BE)
	AV_PIX_FMT_AYUV64LE      =PixelFormat(C.AV_PIX_FMT_AYUV64LE)
	AV_PIX_FMT_AYUV64BE      =PixelFormat(C.AV_PIX_FMT_AYUV64BE)
	AV_PIX_FMT_VIDEOTOOLBOX  =PixelFormat(C.AV_PIX_FMT_VIDEOTOOLBOX)
	AV_PIX_FMT_P010LE        =PixelFormat(C.AV_PIX_FMT_P010LE)
	AV_PIX_FMT_P010BE        =PixelFormat(C.AV_PIX_FMT_P010BE)
	AV_PIX_FMT_GBRAP12BE     =PixelFormat(C.AV_PIX_FMT_GBRAP12BE)
	AV_PIX_FMT_GBRAP12LE     =PixelFormat(C.AV_PIX_FMT_GBRAP12LE)
	AV_PIX_FMT_GBRAP10BE     =PixelFormat(C.AV_PIX_FMT_GBRAP10BE)
	AV_PIX_FMT_GBRAP10LE     =PixelFormat(C.AV_PIX_FMT_GBRAP10LE)
	AV_PIX_FMT_MEDIACODEC    =PixelFormat(C.AV_PIX_FMT_MEDIACODEC)
	AV_PIX_FMT_GRAY12BE      =PixelFormat(C.AV_PIX_FMT_GRAY12BE)
	AV_PIX_FMT_GRAY12LE      =PixelFormat(C.AV_PIX_FMT_GRAY12LE)
	AV_PIX_FMT_GRAY10BE      =PixelFormat(C.AV_PIX_FMT_GRAY10BE)
	AV_PIX_FMT_GRAY10LE      =PixelFormat(C.AV_PIX_FMT_GRAY10LE)
	AV_PIX_FMT_P016LE        =PixelFormat(C.AV_PIX_FMT_P016LE)
	AV_PIX_FMT_P016BE        =PixelFormat(C.AV_PIX_FMT_P016BE)
	AV_PIX_FMT_D3D11         =PixelFormat(C.AV_PIX_FMT_D3D11)
	AV_PIX_FMT_GRAY9BE       =PixelFormat(C.AV_PIX_FMT_GRAY9BE)
	AV_PIX_FMT_GRAY9LE       =PixelFormat(C.AV_PIX_FMT_GRAY9LE)
	AV_PIX_FMT_GBRPF32BE     =PixelFormat(C.AV_PIX_FMT_GBRPF32BE)
	AV_PIX_FMT_GBRPF32LE     =PixelFormat(C.AV_PIX_FMT_GBRPF32LE)
	AV_PIX_FMT_GBRAPF32BE    =PixelFormat(C.AV_PIX_FMT_GBRAPF32BE)
	AV_PIX_FMT_GBRAPF32LE    =PixelFormat(C.AV_PIX_FMT_GBRAPF32LE)
	AV_PIX_FMT_OPENCL        =PixelFormat(C.AV_PIX_FMT_OPENCL)
	AV_PIX_FMT_GRAY14BE      =PixelFormat(C.AV_PIX_FMT_GRAY14BE)
	AV_PIX_FMT_GRAY14LE      =PixelFormat(C.AV_PIX_FMT_GRAY14LE)
	AV_PIX_FMT_GRAYF32BE     =PixelFormat(C.AV_PIX_FMT_GRAYF32BE)
	AV_PIX_FMT_GRAYF32LE     =PixelFormat(C.AV_PIX_FMT_GRAYF32LE)
	AV_PIX_FMT_YUVA422P12BE  =PixelFormat(C.AV_PIX_FMT_YUVA422P12BE)
	AV_PIX_FMT_YUVA422P12LE  =PixelFormat(C.AV_PIX_FMT_YUVA422P12LE)
	AV_PIX_FMT_YUVA444P12BE  =PixelFormat(C.AV_PIX_FMT_YUVA444P12BE)
	AV_PIX_FMT_YUVA444P12LE  =PixelFormat(C.AV_PIX_FMT_YUVA444P12LE)
	AV_PIX_FMT_NV24          =PixelFormat(C.AV_PIX_FMT_NV24)
	AV_PIX_FMT_NV42          =PixelFormat(C.AV_PIX_FMT_NV42)
	AV_PIX_FMT_VULKAN        =PixelFormat(C.AV_PIX_FMT_VULKAN)
	AV_PIX_FMT_Y210BE        =PixelFormat(C.AV_PIX_FMT_Y210BE)
	AV_PIX_FMT_Y210LE        =PixelFormat(C.AV_PIX_FMT_Y210LE)
	AV_PIX_FMT_X2RGB10LE     =PixelFormat(C.AV_PIX_FMT_X2RGB10LE)
	AV_PIX_FMT_X2RGB10BE     =PixelFormat(C.AV_PIX_FMT_X2RGB10BE)
	AV_PIX_FMT_NB            =PixelFormat(C.AV_PIX_FMT_NB)

)
