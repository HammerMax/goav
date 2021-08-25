package main

import (
	"fmt"
	"github.com/HammerMax/goav/avformat"
	"os"
)

func main() {
	filename := os.Args[1]

	streams, err := getStreamInfo(filename)
	if err != nil {
		panic(err)
	}

	for _, s := range streams {
		fmt.Println(s)
	}
}

func getStreamInfo(filename string) ([]*avformat.Stream, error) {
	var formatContext *avformat.Context
	err := avformat.AvformatOpenInput(&formatContext, filename, nil, nil)
	if err != nil {
		return nil, err
	}

	err = formatContext.AvformatFindStreamInfo(nil)
	if err != nil {
		return nil, err
	}
	return formatContext.Streams(), nil
}
