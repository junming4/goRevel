package main

import (
	"github.com/giorgisio/goav/avformat"
	"log"
)

func main()  {
	var (
		ctxtFormat    *avformat.Context
	)
	filename := "/Users/laraveljun/go/src/myapp/examples/sample.mp4"

	// Register all formats and codecs
	avformat.AvRegisterAll()

	// Open video file
	if avformat.AvformatOpenInput(&ctxtFormat, filename, nil, nil) != 0 {
		log.Println("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if ctxtFormat.AvformatFindStreamInfo(nil) < 0 {
		log.Println("Error: Couldn't find stream information.")
		return
	}

}
