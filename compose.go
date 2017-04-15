package main

import (
    "math"
    "gopkg.in/gographics/imagick.v2/imagick"
    "github.com/myamamoto88/lgtm-capture/config"
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

    magicWand.AnnotateImage(textImage, 10, 0, 0, "LGTM")

    textImage.Destroy()
}

func fetchTextImage() *imagick.DrawingWand {
    fillWand := imagick.NewPixelWand()
    defer fillWand.Destroy()
    fillWand.SetColor(config.Instance().GetString("font.base_color"))

    strokeWand := imagick.NewPixelWand()
    defer strokeWand.Destroy()
    strokeWand.SetColor(config.Instance().GetString("font.border_color"))

    drawingWand := imagick.NewDrawingWand()
    drawingWand.SetFont(config.Instance().GetString("font.ttf"))
    drawingWand.SetFillColor(fillWand)
    drawingWand.SetStrokeColor(strokeWand)
    drawingWand.SetGravity(imagick.GRAVITY_SOUTH_WEST)

    return drawingWand
}

func calcFontSize(width, height uint) float64 {
    return math.Min(float64(width), float64(height)) / config.Instance().GetFloat64("font.adjusted_value")
}
