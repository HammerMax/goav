package main

import (
	"fmt"
	"github.com/HammerMax/goav/avutil"
	"github.com/HammerMax/goav/swscale"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s output_file output_size\n" +
		"API example program to show how to scale an image with libswscale.\n" +
		"This program generates a series of pictures, rescales them to the given " +
		"output_size and saves them to an output file named output_file\n." +
		"\n", os.Args[0]);
		return
	}

	dstFilename := os.Args[1]
	dstSize := os.Args[2]

	dstW := 0
	dstH := 0
	err := avutil.AvParseVideoSize(&dstW, &dstH, dstSize)
	if err != nil {
		panic(err)
	}

	dstFile, err := os.Create(dstFilename)
	if err != nil {
		panic(err)
	}

	srcW := 320
	srcH := 240
	srcPixFmt := avutil.AV_PIX_FMT_YUV420P
	dstPixFmt := avutil.AV_PIX_FMT_RGB24
	swsCtx := swscale.SwsGetContext(srcW, srcH, srcPixFmt, dstW, dstH, dstPixFmt, swscale.SWS_BILINEAR, nil, nil, nil)
	if swsCtx == nil {
		panic("swsCtx is nil")
	}

	var srcLineSize = make([]int32, 4)
	var srcData = make([][]byte, 4)
	ret := avutil.AvImageAlloc(srcData, srcLineSize, srcW, srcH, srcPixFmt, 16)
	if ret < 0 {
		panic("alloc image error")
	}

	var dstLineSize = make([]int32, 4)
	var dstData = make([][]byte, 4)
	ret = avutil.AvImageAlloc(dstData, dstLineSize, dstW, dstH, dstPixFmt, 1)
	if ret < 0 {
		panic("alloc image error")
	}

	dstBufSize := ret
	fmt.Println(dstFile, dstBufSize)

	for i:=0; i<100; i++ {
		fillYuvImage(srcData, srcLineSize, srcW, srcH, i)

		swsCtx.SwsScale(srcData, srcLineSize, 0, srcH, dstData, dstLineSize)

		buf := make([]byte, dstBufSize)
		copy(buf, dstData[0])
		_, err = dstFile.Write(buf)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("Scaling succeeded. Play the output file with the command:\nffplay -f rawvideo -pix_fmt %s -video_size %dx%d %s\n",
		avutil.AvGetPixFmtName(dstPixFmt), dstW, dstH, dstFilename)
}

func fillYuvImage(data [][]byte, lineSize []int32, width, height, frameIndex int) {
	for y:=0; y<height; y++ {
		for x:=0; x<width; x++ {
			data[0][y*int(lineSize[0])+x] = byte(x+y+frameIndex*3)
		}
	}

	for y:=0; y<height/2; y++ {
		for x:=0; x<width/2; x++ {
			data[1][y*int(lineSize[1])+x] = byte(120+y+frameIndex*2)
			data[2][y*int(lineSize[2])+x] = byte(64+x+frameIndex*5)
		}
	}
}
