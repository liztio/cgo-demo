package main

// #cgo pkg-config: MagickWand
// #include <wand/MagickWand.h>
import "C"
import (
	"flag"
	"fmt"
	"os"
)

var (
	inputFile  = flag.String("in", ":logo", "The file to read in")
	outputFile = flag.String("out", "logo.jpg", "The file to write out")
	scale      = flag.Float64("scale", 0.5, "how much to scale the image by")
)

func main() {
	flag.Parse()

	if *scale <= 0 {
		fmt.Fprintf(os.Stderr, "Scale must be greater than zero\n")
		os.Exit(1)
	}

	C.MagickWandGenesis()

	// Create a wand
	mw := C.NewMagickWand()

	defer func() {
		// Tidy up
		if mw != nil {
			C.DestroyMagickWand(mw)
		}

		C.MagickWandTerminus()
	}()

	// Read the input image
	C.MagickReadImage(mw, C.CString(*inputFile))

	// Resize it
	height := float64(C.MagickGetImageHeight(mw))
	width := float64(C.MagickGetImageWidth(mw))

	newHeight := int(height * (*scale))
	newWidth := int(width * (*scale))
	if newHeight < 1 {
		newHeight = 1
	}
	if newWidth < 1 {
		newWidth = 1
	}
	C.MagickResizeImage(mw, C.ulong(newWidth), C.ulong(newHeight), C.LanczosFilter, 1)

	// Write the new image
	C.MagickWriteImage(mw, C.CString(*outputFile))

}
