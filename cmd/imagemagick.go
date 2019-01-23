package main

// #cgo CFLAGS: -I ImageMagick-6
// #include <wand/magic_wand.h>
import "C"

func main() {

	C.MagickWandGenesis()

	/* Create a wand */
	mw := C.NewMagickWand()

	/* Read the input image */
	C.MagickReadImage(mw, "logo:")
	/* write it */
	C.MagickWriteImage(mw, "logo.jpg")

	/* Tidy up */
	if mw != nil {
		mw = DestroyMagickWand(mw)
	}

	C.MagickWandTerminus()
}
