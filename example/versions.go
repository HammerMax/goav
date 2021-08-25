package main

import (
	"log"

	"github.com/HammerMax/goav/avcodec"
	"github.com/HammerMax/goav/avdevice"
	"github.com/HammerMax/goav/avfilter"
	"github.com/HammerMax/goav/avformat"
	"github.com/HammerMax/goav/avutil"
	"github.com/HammerMax/goav/swresample"
	"github.com/HammerMax/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
