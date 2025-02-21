package pdf2img

import (
	"github.com/gen2brain/go-fitz"
	"image/jpeg"
	"os"
	"path/filepath"
	"fmt"
)

func Pdf2img(pdfPath string) {
	doc, err := fitz.New(pdfPath)
	if err != nil {
		panic(err)
	}
	defer doc.Close()
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			panic(err)
		}
		err = os.MkdirAll("img/123", 0755)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join("img/123/", fmt.Sprintf("image-%05d.jpg", n)))
		if err != nil {
			panic(err)
		}

		err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
		if err != nil {
			panic(err)
		}

		f.Close()

	}
	
}