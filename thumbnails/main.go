package thumbnails

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"

	"github.com/minodisk/go-fix-orientation/processor"
	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
	"github.com/nfnt/resize"
)

var (
	// ResizeMethod is the parameter that describes an interpolation kernel
	ResizeMethod = resize.Lanczos3
	// JpegQuality default jpeg quality
	JpegQuality = 100
)

// Make thumbnail from io.Reader
func Make(img io.Reader, width, height int, crop bool) (*bytes.Reader, error) {
	thumb, err := fixOrientation(img)
	if err != nil {
		return nil, err
	}
	if crop {
		thumb, err = smartCrop(thumb, width, height)
		if err != nil {
			return nil, err
		}
	}
	resizedImage := resize.Thumbnail(uint(width), uint(height), thumb, ResizeMethod)
	return newReader(resizedImage)
}

func newReader(img image.Image) (*bytes.Reader, error) {
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, img, &jpeg.Options{Quality: JpegQuality})
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(buf.Bytes())
	return reader, nil
}

func smartCrop(img image.Image, width, height int) (image.Image, error) {
	analyzer := smartcrop.NewAnalyzer(nfnt.NewDefaultResizer())
	bestCrop, err := analyzer.FindBestCrop(img, width, height)
	if err != nil {
		return nil, err
	}
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}
	return img.(subImager).SubImage(bestCrop), nil
}

func fixOrientation(img io.Reader) (image.Image, error) {
	return processor.Process(img)
}
