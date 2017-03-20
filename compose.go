package main

import (
    "math"
    "gopkg.in/gographics/imagick.v2/imagick"
)

const (
    FONT = "/Library/Fonts/Impact.ttf"
    FONT_ADJUSTED_VALUE = 2.5
    FONT_COLOR = "white"
)

func compose(imagePath string) {
    imagick.Initialize()
    defer imagick.Terminate()

    coalescedImages := fetchCoalescedImages(imagePath)

    decorate(coalescedImages)

    coalescedImages.OptimizeImageLayers()
    coalescedImages.WriteImages(imagePath, true)

    coalescedImages.Destroy()
}

func fetchCoalescedImages(imagePath string) *imagick.MagickWand {
    magicWand := imagick.NewMagickWand()
    defer magicWand.Destroy()

    magicWand.ReadImage(imagePath)
    magicWand.SetFirstIterator()

    return magicWand.CoalesceImages()
}

func decorate(magicWand *imagick.MagickWand) {
    textImage := fetchTextImage()

    fontSize := calcFontSize(magicWand.GetImageWidth(), magicWand.GetImageHeight())
    textImage.SetFontSize(fontSize)
    textImage.SetTextKerning(fontSize / 10)

    magicWand.AnnotateImage(textImage, 0, 0, 0, "LGTM")
    for magicWand.NextImage() {
      magicWand.AnnotateImage(textImage, 0, 0, 0, "LGTM")
    }

    textImage.Destroy()
}

func fetchTextImage() *imagick.DrawingWand {
    pixelWand := imagick.NewPixelWand()
    defer pixelWand.Destroy()
    pixelWand.SetColor(FONT_COLOR)

    drawingWand := imagick.NewDrawingWand()
    drawingWand.SetFont(FONT)
    drawingWand.SetFillColor(pixelWand)
    drawingWand.SetGravity(imagick.GRAVITY_CENTER)

    return drawingWand
}

func calcFontSize(width, height uint) float64 {
    return math.Min(float64(width), float64(height)) / FONT_ADJUSTED_VALUE
}
